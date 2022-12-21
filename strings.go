package strings

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Operation struct {
	Left     float64
	Right    float64
	Operator string
}

var operators []string

func init() {
	operators = append(operators, "+")
	operators = append(operators, "-")
	operators = append(operators, "*")
	operators = append(operators, "/")
	operators = append(operators, "^")
}

func Eval(s string) (float64, error) {
	ns := strings.ReplaceAll(s, " ", "")
	ns = removeBracket(ns)
	var res float64

	nns, err := normalizeBrackets(ns)
	if err != nil {
		return res, err
	}

	numbs := extractNumbers(nns)
	ops := extractOperators(nns)

	if ln, lo := len(numbs), len(ops); ln < 2 || lo >= ln {
		return res, errors.New("invalid string")
	}

	com := combine(numbs, ops)
	res, err = calculate(com)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func GetOperators() []string {
	return operators
}

func extractNumbers(s string) []string {
	var res []string
	numbs := strings.FieldsFunc(s, splitter)
	for i := 0; i < len(numbs); i++ {
		if !strings.Contains(numbs[i], "(") && !isNegative(numbs[i]) {
			esp := strings.Split(numbs[i], "-")
			res = append(res, esp...)
		} else {
			res = append(res, fmt.Sprintf("(%s)", numbs[i]))
		}
	}
	return res
}

func extractOperators(s string) []string {
	var res []string
	b := []byte(s)
	for i := 1; i < len(b); i++ {
		c := string(s[i])
		prev := string(s[i-1])
		if c == "(" {
			i++
		} else if inArray(c, operators) && !inArray(prev, operators) {
			res = append(res, c)
		}
	}
	return res
}

func (o Operation) calc() (float64, error) {
	switch o.Operator {
	case "+":
		return o.Left + o.Right, nil
	case "-":
		return o.Left - o.Right, nil
	case "*":
		return o.Left * o.Right, nil
	case "/":
		return o.Left / o.Right, nil
	case "^":
		return math.Pow(o.Left, o.Right), nil
	}
	return 0, errors.New("operator is not supported")
}

func indexHighestOperator(s []string) (int, error) {
	var index int
	counter := 0
	for i, e := range s {
		if e == "^" || e == "√" {
			index = i
			break
		}
		if e == "*" || e == "/" {
			index = i
			break
		}
		if e == "+" || e == "-" {
			counter++
			index = i
		}
	}

	if on := operatorNumber(s); on == counter {
		return 1, nil
	}

	if index == 0 {
		return 0, errors.New("operator is not supported")
	}

	return index, nil
}

func removeOperator(s string) []string {
	var new []string
	for _, e := range operators {
		if e != s {
			new = append(new, e)
		}
	}
	return new
}

func isNegative(s string) bool {

	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}

	if res > 0 {
		return false
	}

	return true
}

func normalizeBrackets(s string) (string, error) {
	start := 0
	end := 0
	newString := ""
	var mid string
	for i, e := range s {
		if fmt.Sprintf("%c", e) == "(" {
			start = i
		}

		if fmt.Sprintf("%c", e) == ")" {
			end = i
			break
		}
	}

	if start == 0 && end == 0 {
		return s, nil
	}

	if start > 0 && end == 0 {
		fmt.Println("err")
	}

	ns := s[start+1 : end]
	if isNegative(ns) {
		newString = fmt.Sprintf("%s%s%s", s[0:start], ns, s[end+1:])
		return normalizeBrackets(newString)
	}

	op, a, b := GetValues(ns)

	var o Operation

	o.Left = a
	o.Right = b
	o.Operator = op

	res, err := o.calc()

	if err != nil {
		return "", err
	}

	if res < 0 {
		mid = fmt.Sprintf("(%f)", res)
	} else {
		mid = fmt.Sprintf("%f", res)
	}

	left := s[0:start]
	right := s[end+1:]

	if !inArray(left[len(left)-1:], operators) && left[len(left)-1:] != "(" {
		newString = fmt.Sprintf("%s*%s%s", left, mid, right)
	} else {
		newString = fmt.Sprintf("%s%s%s", left, mid, right)
	}

	return normalizeBrackets(newString)
}

func GetValues(s string) (string, float64, float64) {
	op := findOperator(s)
	nns := strings.Split(s, op)
	var a, b float64

	if nns[0] == "" {
		a, _ = strconv.ParseFloat(fmt.Sprintf("-%s", nns[1]), 64)
		b, _ = strconv.ParseFloat(nns[2], 64)
	} else {
		a, _ = strconv.ParseFloat(nns[0], 64)
		b, _ = strconv.ParseFloat(nns[1], 64)
	}

	return op, a, b
}

func findOperator(s string) string {
	for i, o := range operators {
		if i == 0 && o == "-" {
			continue
		} else if strings.Contains(s, o) {
			return o
		}
	}
	return ""
}

func calculate(com []string) (float64, error) {

	opIndex, err := indexHighestOperator(com)
	if err != nil {
		return 0, err
	}

	a, _ := strconv.ParseFloat(removeBracket(com[opIndex-1]), 64)
	b, _ := strconv.ParseFloat(removeBracket(com[opIndex+1]), 64)

	var o Operation

	o.Left = a
	o.Right = b
	o.Operator = com[opIndex]

	res, err := o.calc()
	if err != nil {
		return 0, err
	}

	if operatorNumber(com) == 1 {
		return res, nil
	}

	com[opIndex-1] = fmt.Sprintf("%f", res)
	newcom := removeIndex(com, opIndex)
	newcom = removeIndex(newcom, opIndex)

	return calculate(newcom)
}

func removeBracket(s string) string {
	l := len(s)
	if s[0:1] == "(" {
		return s[1 : l-1]
	}
	return s
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func operatorNumber(s []string) int {
	counter := 0
	for _, e := range s {
		if inArray(e, operators) {
			counter++
		}
	}
	return counter
}

func inArray(s string, as []string) bool {
	for _, e := range as {
		if s == e {
			return true
		}
	}
	return false
}

func splitter(r rune) bool {
	arr := removeOperator("-")
	o := fmt.Sprintf("%c", r)
	return inArray(o, arr)
}

func combine(number []string, operator []string) []string {
	var res []string
	for i := 0; i < len(operator); i++ {
		res = append(res, number[i])
		res = append(res, operator[i])
		if i == len(operator)-1 {
			res = append(res, number[i+1])
		}
	}
	return res
}
