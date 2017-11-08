package metodos

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/fuzzyqu/trabalho-metodos/models"
)

func RegraDosTrapeziosRepetida(integral models.Integral, n int) (float64, error) {

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	expr, err := govaluate.NewEvaluableExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	// a
	params[integral.Parametro] = integral.A
	r, err := evaluate(expr, params)
	if err != nil {
		return r, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = evaluate(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 1 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluate(expr, params)
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

	expr, err := govaluate.NewEvaluableExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.Parametro] = integral.A
	r, err := evaluate(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = evaluate(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 2 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluate(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 4
	}

	// intervalo
	for i := 2; i < n-1; i += 2 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluate(expr, params)
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

	expr, err := govaluate.NewEvaluableExpression(integral.Expressao)
	if err != nil {
		return 0.0, err
	}

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.Parametro] = integral.A
	r, err := evaluate(expr, params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = evaluate(expr, params)
	if err != nil {
		return r, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 3 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluate(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 3
	}

	// intervalo
	for i := 2; i < n-1; i += 3 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluate(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 3
	}

	// intervalo
	for i := 3; i < n-2; i += 3 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := evaluate(expr, params)
		if err != nil {
			return 0.0, err
		}
		result += r * 2
	}

	return result * step * 3.0 / 8.0, nil
}

func evaluate(expr *govaluate.EvaluableExpression, params map[string]interface{}) (float64, error) {
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
