package Layouts

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/xuri/excelize/v2"
)

var ErrNoSheetFound error = errors.New("no sheet found on file")
var ErrValidationFail error = errors.New("file rows validation fail")

type ExcelLayout struct {
	Layout
}

func (l *ExcelLayout) ParseStruct(r interface{}) []Error {
	s := reflect.ValueOf(r)
	errors := []Error{}

	for i := 0; i < s.NumField(); i++ {
		tags, err := parseOptions(string(s.Type().Field(i).Tag))
		if err == nil {
			f := s.Field(i)
			value := fmt.Sprintf("%v", f)
			switch f.Kind() {
			case reflect.Slice:
				if tags.CommaSeparatedValue {
					values := strings.Split(value, ",")
					for _, v := range values {
						switch reflect.TypeOf(f.Interface()).Elem().Kind() {
						case reflect.String:
							if _, err := parseStringRules(v, tags); err != nil {
								for _, e := range err {
									errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
								}
							}
						case reflect.Float32, reflect.Float64:
							if _, err := parseFloat64Rules(v, tags); err != nil {
								for _, e := range err {
									errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
								}
							}
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							if _, err := parseIntRules(v, tags); err != nil {
								for _, e := range err {
									errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
								}
							}
						}
					}
				} else {
					errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: ErrCommaSeparatedInvalid})
				}
			case reflect.String:
				if _, err := parseStringRules(value, tags); err != nil {
					for _, e := range err {
						errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
					}
				}
			case reflect.Float32, reflect.Float64:
				if _, err := parseFloat64Rules(value, tags); err != nil {
					for _, e := range err {
						errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
					}
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if _, err := parseIntRules(value, tags); err != nil {
					for _, e := range err {
						errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (l *ExcelLayout) ParseCells(r interface{}, cells []string) []Error {
	if l.uniques == nil {
		l.uniques = map[string]int{}
	}

	errors := []Error{}

	s := reflect.ValueOf(r)

	for i := 0; i < s.Elem().NumField(); i++ {
		rowIndex := int(s.Elem().FieldByName("Index").Int())
		tags, err := parseOptions(string(s.Elem().Type().Field(i).Tag))
		if err == nil {
			f := s.Elem().Field(i)
			col, _ := excelize.ColumnNameToNumber(tags.Column)
			col--
			if col >= 0 && col <= len(cells)-1 {

				value := cells[col]

				switch f.Kind() {
				case reflect.Slice:
					if tags.CommaSeparatedValue {
						values := strings.Split(value, ",")
						for _, v := range values {
							switch reflect.TypeOf(f.Interface()).Elem().Kind() {
							case reflect.String:
								if val, err := parseStringRules(v, tags); err != nil {
									for _, e := range err {
										errors = append(errors, Error{RowIndex: rowIndex, Column: tags.Column, Error: e})
									}
								} else {
									f.Set(reflect.Append(f, reflect.ValueOf(val)))
								}
							case reflect.Float32, reflect.Float64:
								if val, err := parseFloat64Rules(v, tags); err != nil {
									for _, e := range err {
										errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
									}
								} else {
									f.Set(reflect.Append(f, reflect.ValueOf(val)))
								}
							case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
								if val, err := parseIntRules(v, tags); err != nil {
									for _, e := range err {
										errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
									}
								} else {
									f.Set(reflect.Append(f, reflect.ValueOf(val)))
								}
							}
						}
					} else {
						errors = append(errors, Error{RowIndex: rowIndex, Column: tags.Column, Error: ErrCommaSeparatedInvalid})
					}
				case reflect.String:
					if val, err := parseStringRules(value, tags); err != nil {
						for _, e := range err {
							errors = append(errors, Error{RowIndex: rowIndex, Error: e, Column: tags.Column})
						}
					} else {
						f.SetString(val)
					}
				case reflect.Float32, reflect.Float64:
					if val, err := parseFloat64Rules(value, tags); err != nil {
						for _, e := range err {
							errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
						}
					} else {
						f.SetFloat(float64(val))
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if val, err := parseIntRules(value, tags); err != nil {
						for _, e := range err {
							errors = append(errors, Error{RowIndex: 0, Column: tags.Column, Error: e})
						}
					} else {
						f.SetInt(int64(val))
					}
				}

				if tags.Unique {
					if _, exists := l.uniques[tags.Column]; exists {
						errors = append(errors, Error{RowIndex: rowIndex, Error: ErrNotUnique, Column: tags.Column})
					} else {
						l.uniques[tags.Column] = rowIndex
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (l *ExcelLayout) ReadFile(rowType interface{}, filePath string) error {

	hasErrors := false
	elType := reflect.TypeOf(rowType)
	elSlice := []interface{}{}

	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer func() {
		xlsx.Close()
	}()
	sheets := xlsx.GetSheetList()

	if len(sheets) == 0 {
		return ErrNoSheetFound
	}

	// Get all the rows in the Sheet1.
	rows, err := xlsx.GetRows(sheets[0])
	if err != nil {
		return err
	}
	for i, row := range rows {
		if i > 0 {
			elItem := reflect.New(elType).Interface()
			f := reflect.Indirect(reflect.ValueOf(elItem)).FieldByName("Index")
			f.SetInt(int64(i) + 1)

			if err := l.ParseCells(elItem, row); err != nil {
				hasErrors = true
				l.errors = append(l.errors, err...)
			}
			elSlice = append(elSlice, elItem)
		}
	}
	l.rows = elSlice
	if hasErrors {
		return ErrValidationFail
	}
	return nil
}
