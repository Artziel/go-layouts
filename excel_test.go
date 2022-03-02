package Layouts

import (
	"reflect"
	"testing"
)

func TestExcelRowParser(t *testing.T) {

	tests := []RowParserTests{
		{
			input: []string{"1", "artziel@gmail.com", "12345678", "https://www.asdasd.com", "Artziel Narvaiza", "44"},
			expected: TestRow{
				ID:       1,
				Username: "artziel@gmail.com",
				Password: "12345678",
				Avatar:   "https://www.asdasd.com",
				Fullname: "Artziel Narvaiza",
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
		if ok, errors := l.Parse(elItem, test.input); ok {
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

func TestExcelFileRead(t *testing.T) {
	l := ExcelLayout{}

	fileName := "./sample/sample.xlsx"

	err := l.ReadFile(TestRow{}, fileName)
	if err == ErrNoSheetFound || err == ErrValidationFail {
		t.Errorf("Test 0: Read file \"%s\" fail. Error: %s ", fileName, err.Error())
	} else if l.CountRows() != 1 {
		t.Errorf("Test 1: Expected one row, Recived: %d", l.CountRows())
	}
}
