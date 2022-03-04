package Layouts

import (
	"fmt"
)

type TestRow struct {
	Row
	ID       int    `excelLayout:"column:A,required,min:1"`
	Username string `excelLayout:"column:B,required,minLength:6"`
	Password string `excelLayout:"column:C,required,minLength:8"`
	Avatar   string `excelLayout:"column:D,url"`
	Fullname string `excelLayout:"column:E,required,maxLength:25"`
	Email    string `excelLayout:"column:F,required,email"`
	Age      int    `excelLayout:"column:G,required,min:18,max:50"`
	Key      string `excelLayout:"column:H,required,regex:p([a-z]+)ch"`
}

/**
 * Field Tag Parser Test struct
 */
type FiledTagsParserTests struct {
	input       string
	expected    fieldTags
	errExpected error
}

/**
 * Compare expected data with parameter
 */
func (pt *FiledTagsParserTests) compareTo(ft *fieldTags) (bool, []string) {

	errors := []string{}

	if pt.expected.Column != ft.Column {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Column, recieved: \"%v\"",
			pt.expected.Column, ft.Column,
		))
	}
	if pt.expected.CommaSeparatedValue != ft.CommaSeparatedValue {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field CommaSeparatedValue, recieved: \"%v\"",
			pt.expected.Column, ft.Column,
		))
	}
	if pt.expected.Email != ft.Email {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Email, recieved: \"%v\"",
			pt.expected.Email, ft.Email,
		))
	}
	if pt.expected.Required != ft.Required {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Required, recieved: \"%v\"",
			pt.expected.Required, ft.Required,
		))
	}
	if pt.expected.Regex != ft.Regex {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Regex, recieved: \"%v\"",
			pt.expected.Regex, ft.Regex,
		))
	}
	if pt.expected.Max != ft.Max {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Max, recieved: \"%v\"",
			pt.expected.Max, ft.Max,
		))
	}
	if pt.expected.Min != ft.Min {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Min, recieved: \"%v\"",
			pt.expected.Min, ft.Min,
		))
	}
	if pt.expected.Url != ft.Url {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Url, recieved: \"%v\"",
			pt.expected.Url, ft.Url,
		))
	}
	if pt.expected.Unique != ft.Unique {
		errors = append(errors, fmt.Sprintf(
			"Expected \"%v\" for field Unique, recieved: \"%v\"",
			pt.expected.Unique, ft.Unique,
		))
	}

	return len(errors) == 0, errors
}

/**
 * Row Parser Test struct
 */
type RowParserTests struct {
	input       []string
	expected    TestRow
	errExpected []error
}

/**
 * Compare expected data with parameter
 */
func (rt *RowParserTests) compareTo(r *TestRow) (bool, []string) {
	errors := []string{}

	if r.ID != rt.expected.ID {
		errors = append(
			errors,
			fmt.Sprintf("ID expected \"%v\", Recived: \"%v\"", rt.expected.ID, r.ID),
		)
	}
	if r.Username != rt.expected.Username {
		errors = append(
			errors,
			fmt.Sprintf("Username expected \"%v\", Recived: \"%v\"", rt.expected.Username, r.Username),
		)
	}
	if r.Password != rt.expected.Password {
		errors = append(
			errors,
			fmt.Sprintf("Password expected \"%v\", Recived: \"%v\"", rt.expected.Password, r.Password),
		)
	}
	if r.Avatar != rt.expected.Avatar {
		errors = append(
			errors,
			fmt.Sprintf("Avatar expected \"%v\", Recived: \"%v\"", rt.expected.Avatar, r.Avatar),
		)
	}
	if r.Fullname != rt.expected.Fullname {
		errors = append(
			errors,
			fmt.Sprintf("Fullname expected \"%v\", Recived: \"%v\"", rt.expected.Fullname, r.Fullname),
		)
	}
	if r.Age != rt.expected.Age {
		errors = append(
			errors,
			fmt.Sprintf("Age expected \"%v\", Recived: \"%v\"", rt.expected.Age, r.Age),
		)
	}

	return len(errors) == 0, errors
}

/**
 * Check if an error is in expected errors list
 */
func (rt *RowParserTests) IsErrorExpected(e error) bool {
	if rt.errExpected == nil {
		return false
	}

	for _, err := range rt.errExpected {
		if err == e {
			return true
		}
	}

	return false
}

type StructParserTests struct {
	input       TestRow
	errExpected []error
}

/**
 * Check if an error is in expected errors list
 */
func (rt *StructParserTests) IsErrorExpected(e error) bool {
	if rt.errExpected == nil {
		return false
	}

	for _, err := range rt.errExpected {
		if err == e {
			return true
		}
	}

	return false
}
