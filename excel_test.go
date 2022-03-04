package Layouts

import (
	"reflect"
	"testing"
)

func TestExcelRowParser(t *testing.T) {

	tests := []RowParserTests{
		{
			input: []string{
				"1", "xxx@yyy.com", "12345678", "https://www.asdasd.com", "Artziel Narvaiza",
				"xxx@yyy.com", "44",
			},
			expected: TestRow{
				ID:       1,
				Username: "xxx@yyy.com",
				Password: "12345678",
				Avatar:   "https://www.asdasd.com",
				Fullname: "Artziel Narvaiza",
				Email:    "xxx@yyy.com",
				Age:      44,
			},
			errExpected: nil,
		},
	}

	l := ExcelLayout{}

	elType := reflect.TypeOf(TestRow{})
	for i, test := range tests {
		elItem := reflect.New(elType).Interface()
		f := reflect.Indirect(reflect.ValueOf(elItem)).FieldByName("Index")
		f.SetInt(int64(i) + 1)
		result := elItem.(*TestRow)
		if ok, errors := l.ParseCells(elItem, test.input); ok {
			if success, errors := test.compareTo(result); !success {
				for _, e := range errors {
					t.Errorf("Test %d: %s\n", i, e)
				}
			}
		} else {
			if len(errors) > 0 && test.errExpected != nil {
				t.Errorf("Test %d: Error expected, recive none\n", i)
			} else {
				for _, e := range errors {
					if !test.IsErrorExpected(e.Error) {
						t.Errorf("Test %d: Error not expected, recived: %s\n", i, ErrToMessage(&e))
					}
				}
			}
		}
	}
}

func TestExcelStructParser(t *testing.T) {

	tests := []StructParserTests{
		{
			input: TestRow{
				ID:       1,
				Username: "xxx@yyy.com",
				Password: "123456",
				Avatar:   ".asdasd.com",
				Fullname: "Artziel Narvaiza",
				Email:    "xxx@yyy.com",
				Age:      12,
			},
			errExpected: []error{
				ErrMinLengthValueRuleFail,
				ErrUrlValueRuleFail,
				ErrMinValueRuleFail,
			},
		},
		{
			input: TestRow{
				ID:       1,
				Username: "xxx@yyy.com",
				Password: "",
				Avatar:   ".asdasd.com",
				Fullname: "Artziel Ángel Narvaiza González",
				Email:    "xxx@yyy.com",
				Age:      100,
			},
			errExpected: []error{
				ErrMinLengthValueRuleFail,
				ErrUrlValueRuleFail,
				ErrMaxValueRuleFail,
				ErrRequiredValueRuleFail,
				ErrMaxLengthValueRuleFail,
			},
		},
	}

	l := ExcelLayout{}

	for i, test := range tests {
		errs := l.ParseStruct(test.input)
		if errs == nil && test.errExpected != nil {
			t.Errorf("Test %d: Expected Error, None recibed\n", i)
		} else if errs != nil && test.errExpected == nil {
			for _, e := range errs {
				t.Errorf("Test %d: Unexpected error: %s\n", i, e.Error.Error())
			}
		} else {
			for _, e := range errs {
				if !test.IsErrorExpected(e.Error) {
					t.Errorf("Test %d: Unexpected error: %s\n", i, e.Error.Error())
				}
			}
		}
	}
}

func TestExcelFileRead(t *testing.T) {
	l := ExcelLayout{}

	fileName := "./sample/sample.xlsx"

	err := l.ReadFile(TestRow{}, fileName)
	rows := l.GetRows().([]TestRow)
	if err == ErrNoSheetFound || err == ErrValidationFail {
		t.Errorf("Test 0: Read file \"%s\" fail. Error: %s ", fileName, err.Error())
	} else if len(rows) != 1 {
		t.Errorf("Test 1: Expected one row, Recived: %d", len(rows))
	}
}
