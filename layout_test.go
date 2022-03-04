package Layouts

import (
	"errors"
	"testing"
)

type errTests struct {
	err      Error
	expected string
}

func TestErrToMessage(t *testing.T) {

	tests := []errTests{
		{Error{Error: ErrRequiredValueRuleFail, Column: "A"}, "El valor de la columna \"A\" es requerido"},
		{Error{Error: ErrMinValueRuleFail, Column: "A"}, "El valor de la columna \"A\" es menor al mínimo permitido"},
		{Error{Error: ErrMaxValueRuleFail, Column: "A"}, "El valor de la columna \"A\" es mayor al máximo permitido"},
		{Error{Error: ErrUrlValueRuleFail, Column: "A"}, "El valor de la columna \"A\" no es una URL válida"},
		{Error{Error: ErrEmailValueRuleFail, Column: "A"}, "El valor de la columna \"A\" no es un correo electrónico válido"},
		{Error{Error: ErrRegexRuleFail, Column: "A"}, "El valor de la columna \"A\" no cumple con la expresión regular"},
		{Error{Error: ErrRegexInvalid, Column: "A"}, "La expresión regular definida para la columna \"A\" es inválida"},
		{Error{Error: ErrIntegerInvalid, Column: "A"}, "El valor de la columna \"A\" no es un valor entero válido"},
		{Error{Error: ErrDecimalInvalid, Column: "A"}, "El valor de la columna \"A\" no es un valor decimal válido"},
		{Error{Error: ErrNotUnique, Column: "A"}, "El valor de la columna \"A\" debe ser único por archivo"},
		{Error{Error: ErrCommaSeparatedInvalid, Column: "A"}, "El valor de la columna \"A\" no es un valor asignable a un arreglo"},
		{Error{Error: ErrMaxLengthValueRuleFail, Column: "A"}, "La longitud del valor de la columna \"A\" es mayor a la permitida"},
		{Error{Error: ErrMinLengthValueRuleFail, Column: "A"}, "La longitud del valor de la columna \"A\" es menor a la permitida"},
		{Error{Error: ErrTagMinForbidden, Column: "A"}, "Error de definición en la columna \"A\". No se puede definir un valor mínimo para cadenas de caracteres"},
		{Error{Error: ErrTagMaxForbidden, Column: "A"}, "Error de definición en la columna \"A\". No se puede definir un valor máximo para cadenas de caracteres"},
		{Error{Error: ErrTagMinLengthForbidden, Column: "A"}, "Error de definición en la columna \"A\". No se puede definir una longitud mínima para valores numéricos"},
		{Error{Error: ErrTagMaxLengthForbidden, Column: "A"}, "Error de definición en la columna \"A\". No se puede definir una longitud máxima para valores numéricos"},
		{Error{Error: errors.New("unkown error"), Column: "A"}, "Ocurrió un error desconocido al evaluar el valor de la columna \"A\""},
	}

	for i, test := range tests {
		message := ErrToMessage(&test.err)
		if test.expected != message {
			t.Errorf("Test %d: Unexpected Message recived\n", i)
		}
	}
}

func TestLayoutStruct(t *testing.T) {

	l := Layout{
		rows: []interface{}{
			TestRow{},
			TestRow{},
			TestRow{},
			TestRow{},
			TestRow{},
		},
		errors: []Error{
			{Error: ErrRequiredValueRuleFail, Column: "A"},
			{Error: ErrRequiredValueRuleFail, Column: "B"},
			{Error: ErrRequiredValueRuleFail, Column: "C"},
			{Error: ErrRequiredValueRuleFail, Column: "D"},
			{Error: ErrRequiredValueRuleFail, Column: "E"},
		},
	}

	errs := l.GetErrors()
	if len(errs) != 5 {
		t.Errorf("Test 0: GetErrors should return 5 errors, Return %d \n", len(errs))
	}

	rows := l.GetRows()
	if len(rows) != 5 {
		t.Errorf("Test 1: GetRows should return 5 rows, Return %d\n", len(rows))
	}

	n := l.CountRows()
	if n != 5 {
		t.Errorf("Test 2: CountRows should return 5 rows, Return %d\n", n)
	}

}
