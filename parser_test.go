package Layouts

import (
	"testing"
)

/**
 * Test Parse options expected errors
 */
func TestTagsParserOptionsErrors(t *testing.T) {

	tests := []FiledTagsParserTests{

		// General tests
		{``, fieldTags{}, ErrTagNoFieldTag},
		{`excelLayout:`, fieldTags{}, ErrTagEmptyFieldTag},
		{`excelLayout:"`, fieldTags{}, ErrTagEmptyFieldTag},
		{`excelLayout:""`, fieldTags{}, ErrTagEmptyFieldTag},

		// Column field tests
		{`excelLayout:"column"`, fieldTags{}, ErrTagMissingColumnValue},

		{`excelLayout:"column:"`, fieldTags{}, ErrTagMissingColumnValue},
		{`excelLayout:"column: "`, fieldTags{}, ErrTagMissingColumnValue},
		{`excelLayout:" column: "`, fieldTags{}, ErrTagMissingColumnValue},
		{`excelLayout:"cOlumN:A"`, fieldTags{}, nil},

		// Regex field tests
		{`excelLayout:"regex"`, fieldTags{}, ErrTagMissingRegexValue},
		{`excelLayout:"regex:"`, fieldTags{}, ErrTagMissingRegexValue},
		{`excelLayout:"regex: "`, fieldTags{}, ErrTagMissingRegexValue},
		{`excelLayout:" regex: "`, fieldTags{}, ErrTagMissingRegexValue},
		{`excelLayout:"rEgEx:ASDF"`, fieldTags{}, nil},

		// Min field tests
		{`excelLayout:"min"`, fieldTags{}, ErrTagMissingMinValue},
		{`excelLayout:"min:"`, fieldTags{}, ErrTagMissingMinValue},
		{`excelLayout:"min: "`, fieldTags{}, ErrTagMissingMinValue},
		{`excelLayout:" min: "`, fieldTags{}, ErrTagMissingMinValue},
		{`excelLayout:"mIn:1"`, fieldTags{}, nil},

		// Max field tests
		{`excelLayout:"max"`, fieldTags{}, ErrTagMissingMaxValue},
		{`excelLayout:"max:"`, fieldTags{}, ErrTagMissingMaxValue},
		{`excelLayout:"max: "`, fieldTags{}, ErrTagMissingMaxValue},
		{`excelLayout:" max: "`, fieldTags{}, ErrTagMissingMaxValue},
		{`excelLayout:"mAx:1"`, fieldTags{}, nil},

		// Max / Min field tests
		{`excelLayout:"max:1,min:2"`, fieldTags{}, ErrTagInvalidMaxMinValues},
		{`excelLayout:"min:2,max:1"`, fieldTags{}, ErrTagInvalidMaxMinValues},
		{`excelLayout:"min:1,max:2"`, fieldTags{}, nil},
		{`excelLayout:"min:2,max:2"`, fieldTags{}, nil},
	}

	for i, test := range tests {
		_, err := parseOptions(test.input)
		if err != nil {
			if test.errExpected != err && test.errExpected != nil {
				t.Errorf("Test %d Failed:\n Return [ %s ]\n Expected [ %s ]", i, err.Error(), test.errExpected.Error())
			} else if test.errExpected == nil {
				t.Errorf("Test %d Failed:\n Return [ %s ]\n Expected [ No Error ]", i, err.Error())
			}
		} else if err == nil && test.errExpected != nil {
			t.Errorf("Test %d Failed:\n Return [ No Error ]\n Expected [ %s ]", i, test.errExpected.Error())
		}
	}
}

func TestTagsParserOptionsValues(t *testing.T) {

	tests := []FiledTagsParserTests{
		{
			`excelLayout:"COLUMN:A,MIN:1,MAX:1,REQUIRED,COMMASEPARATEDVALUE,EMAIL,REGEX:ASDF,URL,UNIQUE"`,
			fieldTags{
				Column: "A", Min: 1, Max: 1, Required: true, CommaSeparatedValue: true, Email: true,
				Regex: "ASDF", Url: true, Unique: true,
			},
			nil,
		},
		{
			`excelLayout:"column:a,min:1,max:1,required,commaseparatedvalue,email,regex:ASDF,url,unique"`,
			fieldTags{
				Column: "A", Min: 1, Max: 1, Required: true, CommaSeparatedValue: true, Email: true,
				Regex: "ASDF", Url: true, Unique: true,
			},
			nil,
		},
		{
			`excelLayout:"column:A"`, fieldTags{Column: "A"}, nil,
		},
		{
			`excelLayout:"column:A"`, fieldTags{Column: "A"}, nil,
		},
		{
			`excelLayout:"column:A"`, fieldTags{Column: "A"}, nil,
		},
		{
			`excelLayout:"column:A"`, fieldTags{Column: "A"}, nil,
		},
	}

	for i, test := range tests {
		fo, err := parseOptions(test.input)

		if err != nil {
			if test.errExpected != err && test.errExpected != nil {
				t.Errorf("Test %d Failed:\n Return [ %s ]\n Expected [ %s ]", i, err.Error(), test.errExpected.Error())
			} else if test.errExpected == nil {
				t.Errorf("Test %d Failed:\n Return [ %s ]\n Expected [ No Error ]", i, err.Error())
			}
		} else if err == nil && test.errExpected != nil {
			t.Errorf("Test %d Failed:\n Return [ No Error ]\n Expected [ %s ]", i, test.errExpected.Error())
		} else {
			if ok, errs := test.compareTo(&fo); !ok {
				t.Errorf("Test %d Failed:", i)
				for _, e := range errs {
					t.Errorf("\t%s", e)
				}
			}
		}

	}
}
