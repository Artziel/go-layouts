package Layouts

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var ErrTagNoFieldTag error = errors.New("no \"excelLayout\" tag found")
var ErrTagEmptyFieldTag error = errors.New("empty \"excelLayout\" tag found")
var ErrTagMissingColumnValue error = errors.New("expected value for \"column\" tag entry")
var ErrTagMissingRegexValue error = errors.New("expected value for \"regex\" tag entry")
var ErrTagMissingMaxValue error = errors.New("expected value for \"max\" tag entry")
var ErrTagMissingMinValue error = errors.New("expected value for \"min\" tag entry")
var ErrTagInvalidMaxMinValues error = errors.New("the \"max\" value should be greater than \"min\" value tag entry")
var ErrTagInvalidMaxMinLengthValues error = errors.New("the \"maxLength\" value should be greater than \"minLength\" value tag entry")
var ErrTagMissingMinLengthValue error = errors.New("expected value for \"minLength\" tag entry")
var ErrTagMissingMaxLengthValue error = errors.New("expected value for \"maxLength\" tag entry")
var ErrTagMinForbidden error = errors.New("the use of value \"min\" tag entry is not allow for strings")
var ErrTagMaxForbidden error = errors.New("the use of value \"max\" tag entry is not allow for strings")
var ErrTagMinLengthForbidden error = errors.New("the use of value \"minLength\" tag entry is not allow for numbers")
var ErrTagMaxLengthForbidden error = errors.New("the use of value \"maxLength\" tag entry is not allow for numbers")

var ErrRequiredValueRuleFail error = errors.New("value required rule fail")
var ErrMinValueRuleFail error = errors.New("min value rule fail")
var ErrMaxValueRuleFail error = errors.New("max value rule fail")
var ErrMinLengthValueRuleFail error = errors.New("min length rule fail")
var ErrMaxLengthValueRuleFail error = errors.New("max length rule fail")
var ErrUrlValueRuleFail error = errors.New("url value rule validation fail")
var ErrEmailValueRuleFail error = errors.New("email value rule validation fail")
var ErrRegexRuleFail error = errors.New("regex matching rule fail")
var ErrRegexInvalid error = errors.New("invalid regex value")
var ErrIntegerInvalid error = errors.New("invalid integer value")
var ErrDecimalInvalid error = errors.New("invalid integer value")
var ErrCommaSeparatedInvalid error = errors.New("invalid comma separated expected value")
var ErrNotUnique error = errors.New("value is not unique")

type fieldTags struct {
	Column              string
	CommaSeparatedValue bool
	Email               bool
	Required            bool
	Regex               string
	Max                 float64
	Min                 float64
	MaxLength           int64
	MinLength           int64
	Url                 bool
	Unique              bool
	hasMin              bool
	hasMax              bool
	hasMinLength        bool
	hasMaxLength        bool
}

func parseOptions(tags string) (fieldTags, error) {
	ft := fieldTags{}
	tags = strings.TrimSpace(tags)

	if len(tags) == 0 {
		return ft, ErrTagNoFieldTag
	}

	name := "excelLayout:"
	res := strings.Index(tags, name)
	if res < 0 {
		return ft, ErrTagNoFieldTag
	}
	tags = tags[res+len(name):]
	if len(tags) == 0 {
		return ft, ErrTagEmptyFieldTag
	}

	res = strings.Index(tags, "\"")
	if res < 0 {
		return ft, ErrTagEmptyFieldTag
	}
	tags = tags[res+1:]

	res = strings.Index(tags, "\"")
	if res < 0 {
		return ft, ErrTagEmptyFieldTag
	}
	tags = strings.TrimSpace(tags[:res])
	if len(tags) < 1 {
		return ft, ErrTagEmptyFieldTag
	}
	options := strings.Split(tags, ",")

	if len(options) == 0 {
		return ft, ErrTagEmptyFieldTag
	}

	for _, o := range options {
		pair := strings.SplitN(o, ":", 2)
		key := strings.ToLower(strings.TrimSpace(pair[0]))
		val := ""
		if len(pair) > 1 {
			val = strings.TrimSpace(pair[1])
		}

		switch key {
		case "column":
			if val == "" {
				return ft, ErrTagMissingColumnValue
			}
			ft.Column = strings.TrimSpace(strings.ToUpper(pair[1]))
		case "commaseparatedvalue":
			ft.CommaSeparatedValue = true
		case "regex":
			if val == "" {
				return ft, ErrTagMissingRegexValue
			}
			ft.Regex = strings.TrimSpace(pair[1])
		case "email":
			ft.Email = true
		case "required":
			ft.Required = true
		case "max":
			if val == "" {
				return ft, ErrTagMissingMaxValue
			}
			ft.hasMax = true
			v, _ := strconv.ParseFloat(val, 32)
			ft.Max = v
		case "min":
			if val == "" {
				return ft, ErrTagMissingMinValue
			}
			ft.hasMin = true
			v, _ := strconv.ParseFloat(val, 32)
			ft.Min = v
		case "maxlength":
			if val == "" {
				return ft, ErrTagMissingMaxValue
			}
			ft.hasMaxLength = true
			v, _ := strconv.ParseInt(val, 0, 32)
			ft.MaxLength = v
		case "minlength":
			if val == "" {
				return ft, ErrTagMissingMaxValue
			}
			ft.hasMinLength = true
			v, _ := strconv.ParseInt(val, 0, 32)
			ft.MinLength = v
		case "url":
			ft.Url = true
		case "unique":
			ft.Unique = true
		}

	}

	if ft.Regex != "" {
		if _, err := regexp.Compile(ft.Regex); err != nil {
			return ft, ErrRegexInvalid
		}
	}

	if (ft.hasMax && ft.hasMin) && (ft.Max < ft.Min) {
		return ft, ErrTagInvalidMaxMinValues
	}

	if (ft.hasMaxLength && ft.hasMinLength) && (ft.MaxLength < ft.MinLength) {
		return ft, ErrTagInvalidMaxMinLengthValues
	}

	return ft, nil
}

func parseStringRules(v string, tags fieldTags) (string, error) {
	value := strings.TrimSpace(v)
	if tags.Required && len(value) == 0 {
		return "", ErrRequiredValueRuleFail
	}
	if tags.hasMin {
		return "", ErrTagMinForbidden
	}
	if tags.hasMax {
		return "", ErrTagMaxForbidden
	}
	if len(value) > 0 {

		if tags.hasMinLength && (int(tags.MinLength) > len(value)) {
			return "", ErrMinLengthValueRuleFail
		}
		if tags.hasMaxLength && (int(tags.MaxLength) < len(value)) {
			return "", ErrMaxLengthValueRuleFail
		}
		if tags.Url {
			if _, err := url.ParseRequestURI(value); err != nil {
				return "", ErrUrlValueRuleFail
			}
		}
		if tags.Email {
			if _, err := mail.ParseAddress(value); err != nil {
				return "", ErrEmailValueRuleFail
			}
		}
		if tags.Regex != "" {
			regex, err := regexp.Compile(tags.Regex)
			if err != nil {
				return "", ErrRegexInvalid
			}
			if match := regex.MatchString(value); !match {
				return "", ErrRegexRuleFail
			}
		}
	}
	return value, nil
}

func parseIntRules(v string, tags fieldTags) (int64, error) {
	value := strings.TrimSpace(v)
	if tags.Required && value == "" {
		return 0, ErrRequiredValueRuleFail
	}
	if tags.hasMinLength {
		return 0, ErrTagMinLengthForbidden
	}
	if tags.hasMaxLength {
		return 0, ErrTagMaxLengthForbidden
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, ErrIntegerInvalid
	}
	if tags.hasMin && int(tags.Min) > val {
		return 0, ErrMinValueRuleFail
	}
	if tags.hasMax && int(tags.Max) < val {
		return 0, ErrMaxValueRuleFail
	}
	return int64(val), nil
}

func parseFloat64Rules(v string, tags fieldTags) (float64, error) {
	value := strings.TrimSpace(v)
	if tags.Required && value == "" {
		return 0, ErrRequiredValueRuleFail
	}
	if tags.hasMinLength {
		return 0, ErrTagMinLengthForbidden
	}
	if tags.hasMaxLength {
		return 0, ErrTagMaxLengthForbidden
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, ErrDecimalInvalid
	}
	if tags.hasMin && tags.Min > val {
		return 0, ErrMinValueRuleFail
	}
	if tags.hasMax && tags.Max < val {
		return 0, ErrMaxValueRuleFail
	}

	return val, nil
}
