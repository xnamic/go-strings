package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidString(t *testing.T) {
	s := "(1+1+)"
	_, err := Eval(s)
	assert.Equal(t, "invalid string", err.Error())
}

func TestValidString(t *testing.T) {
	s := "1+1+3"
	e := 1 + 1 + 3
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestBracket(t *testing.T) {
	s := "1+(1+3) * 4"
	e := 1 + (1+3)*4
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestNestedBracket(t *testing.T) {
	s := "1+(1+3) * ((4 - 3) - (3+3))"
	e := 1 + (1+3)*((4-3)-(3+3))
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestNestedBracketWithNegativeNumber(t *testing.T) {
	s := "1+(1+3) * ((4 - 100) - (3+3))"
	e := 1 + (1+3)*((4-100)-(3+3))
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestMultifyRightNegative(t *testing.T) {
	s := "4 * -10"
	e := 4 * -10
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestMultifyLeftNegative(t *testing.T) {
	s := "-10 * 5"
	e := -10 * 5
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestMultifyBothNegative(t *testing.T) {
	s := "-10 * -25"
	e := -10 * -25
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestDivideLeftNegative(t *testing.T) {
	s := "-10.0 / 25.0"
	e := -10.0 / 25.0
	res, _ := Eval(s)
	assert.Equal(t, e, res)
}

func TestDivideRIghtNegative(t *testing.T) {
	s := "10.0 / -25.0"
	e := 10.0 / -25.0
	res, _ := Eval(s)
	assert.Equal(t, e, res)
}

func TestDivideBothNegative(t *testing.T) {
	s := "-10.0 / -25.0"
	e := -10.0 / -25.0
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestBracketAsMultiply(t *testing.T) {
	s := "5 (10-2)"
	e := 5 * (10 - 2)
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}

func TestPositiveBracket(t *testing.T) {
	s := "5 + (10)"
	e := 5 + (10)
	res, _ := Eval(s)
	assert.Equal(t, float64(e), res)
}
