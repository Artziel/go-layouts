package Layouts

import "fmt"

/**
 * Row Layout Structure
 */
type Row struct {
	Index int
}

/**
 * Row error structure
 */
type Error struct {
	RowIndex int
	Error    error
	Column   string
}

/**
 * Return redeable error type version
 */
func ErrToMessage(e *Error) string {

	message := ""

	switch e.Error {
	case ErrRequiredValueRuleFail:
		message = fmt.Sprintf("El valor de la columna \"%s\" es requerido", e.Column)
	case ErrMinValueRuleFail:
		message = fmt.Sprintf("El valor de la columna \"%s\" es menor al mínimo permitido", e.Column)
	case ErrMaxValueRuleFail:
		message = fmt.Sprintf("El valor de la columna \"%s\" es mayor al máximo permitido", e.Column)
	case ErrUrlValueRuleFail:
		message = fmt.Sprintf("El valor de la columna \"%s\" no es una URL válida", e.Column)
	case ErrEmailValueRuleFail:
		message = fmt.Sprintf("El valor de la columna \"%s\" no es un correo electrónico válido", e.Column)
	case ErrRegexRuleFail:
		message = fmt.Sprintf("El valor de la columna \"%s\" no cumple con la expresión regular", e.Column)
	case ErrRegexInvalid:
		message = fmt.Sprintf("La expresión regular definida para la columna \"%s\" es inválida", e.Column)
	case ErrIntegerInvalid:
		message = fmt.Sprintf("El valor de la columna \"%s\" no es un valor entero válido", e.Column)
	case ErrDecimalInvalid:
		message = fmt.Sprintf("El valor de la columna \"%s\" no es un valor decimal válido", e.Column)
	case ErrNotUnique:
		message = fmt.Sprintf("El valor de la columna \"%s\" debe ser único por archivo", e.Column)
	case ErrCommaSeparatedInvalid:
		message = fmt.Sprintf("El valor de la columna \"%s\" no es un valor asignable a un arreglo", e.Column)
	case ErrMaxLengthValueRuleFail:
		message = fmt.Sprintf("La longitud del valor de la columna \"%s\" es mayor a la permitida", e.Column)
	case ErrMinLengthValueRuleFail:
		message = fmt.Sprintf("La longitud del valor de la columna \"%s\" es menor a la permitida", e.Column)
	case ErrTagMinForbidden:
		message = fmt.Sprintf("Error de definición en la columna \"%s\". No se puede definir un valor mínimo para cadenas de caracteres", e.Column)
	case ErrTagMaxForbidden:
		message = fmt.Sprintf("Error de definición en la columna \"%s\". No se puede definir un valor máximo para cadenas de caracteres", e.Column)
	case ErrTagMinLengthForbidden:
		message = fmt.Sprintf("Error de definición en la columna \"%s\". No se puede definir una longitud mínima para valores numéricos", e.Column)
	case ErrTagMaxLengthForbidden:
		message = fmt.Sprintf("Error de definición en la columna \"%s\". No se puede definir una longitud máxima para valores numéricos", e.Column)
	default:
		message = fmt.Sprintf("Ocurrió un error desconocido al evaluar el valor de la columna \"%s\"", e.Column)
	}

	return message
}

/**
 * Layout base structure
 */
type Layout struct {
	rows    []interface{}
	uniques map[string]int
	errors  []Error
}

func (l *Layout) CountRows() int {
	return len(l.rows)
}

func (l *Layout) GetRows() []interface{} {
	return l.rows
}

func (l *Layout) GetErrors() []Error {
	return l.errors
}
