package metodos

import (
	"fmt"

	"math"

	"github.com/Knetic/govaluate"
	"github.com/fuzzyqu/trabalho-metodos/models"
)

func RegraDosTrapeziosRepetida(integral models.Integral, n int) (float64, error) {

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	expr, err := newExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	// a
	params[integral.Parametro] = integral.A
	r, err := evaluateExpression(expr, params)
	if err != nil {
		return r, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 1 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluateExpression(expr, params)
		if err != nil {
			return r, err
		}
		result += r * 2
	}

	return result * (step / 2.0), nil
}

func RegraDeSimpson13Repetida(integral models.Integral, n int) (float64, error) {

	if n&1 != 0 {
		return 0.0, fmt.Errorf("n must be even n: %d", n)
	}

	expr, err := newExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.Parametro] = integral.A
	r, err := evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 2 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluateExpression(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 4
	}

	// intervalo
	for i := 2; i < n-1; i += 2 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluateExpression(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 2
	}

	return result * (step / 3.0), nil
}

func RegraDeSimpson38Repetida(integral models.Integral, n int) (float64, error) {

	if n%3 != 0 {
		return 0.0, fmt.Errorf("n must be multiple of 3 n: %d", n)
	}

	expr, err := newExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.Parametro] = integral.A
	r, err := evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 3 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluateExpression(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 3.0
	}

	// intervalo
	for i := 2; i < n-1; i += 3 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluateExpression(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 3.0
	}

	// intervalo
	for i := 3; i < n-2; i += 3 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluateExpression(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 2.0
	}

	return result * step * 3.0 / 8.0, nil
}

func RegraNewtonCotes4(integral models.Integral) (float64, error) {

	expr, err := newExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	step := (integral.B - integral.A) / 4.0

	var result float64
	params := make(map[string]interface{}, 1)

	expr, err = newExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	params[integral.Parametro] = integral.A
	r, err := evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r * 7.0

	params[integral.Parametro] = integral.A + step*1
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r * 32.0

	params[integral.Parametro] = integral.A + step*2
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r * 12.0

	params[integral.Parametro] = integral.A + step*3
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r * 32.0

	params[integral.Parametro] = integral.A + step*4
	r, err = evaluateExpression(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r * 7.0

	return result * (step * (2.0 / 45.0)), nil
}

func evaluateExpression(expr *govaluate.EvaluableExpression, params map[string]interface{}) (float64, error) {
	t, err := expr.Evaluate(params)
	if err != nil {
		return 0.0, err
	}
	result, ok := t.(float64)
	if !ok {
		return 0.0, err
	}
	return result, nil
}

func newExpression(expr string) (*govaluate.EvaluableExpression, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"cos": func(args ...interface{}) (interface{}, error) {
			c := math.Cos(args[0].(float64))
			return (float64)(c), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			c := math.Sin(args[0].(float64))
			return (float64)(c), nil
		},
		"abs": func(args ...interface{}) (interface{}, error) {
			c := math.Abs(args[0].(float64))
			return (float64)(c), nil
		},
		"log2": func(args ...interface{}) (interface{}, error) {
			c := math.Log2(args[0].(float64))
			return (float64)(c), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			c := math.Log10(args[0].(float64))
			return (float64)(c), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
			c := math.Log(args[0].(float64))
			return (float64)(c), nil
		},
		"tan": func(args ...interface{}) (interface{}, error) {
			c := math.Tan(args[0].(float64))
			return (float64)(c), nil
		},
	}
	evaluable, err := govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
	if err != nil {
		return &govaluate.EvaluableExpression{}, err
	}
	return evaluable, nil
}
