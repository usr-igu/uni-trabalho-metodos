package metodos

import (
	"fmt"

	"github.com/fuzzyqu/trabalho-metodos/models"
	"github.com/soudy/mathcat"
)

func RegraDosTrapezios(integral models.Integral, n int) (float64, error) {

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]float64, 1)

	// a
	params[integral.Parametro] = integral.A
	r, err := mathcat.Exec(integral.Expressao, params)
	if err != nil {
		return r, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = mathcat.Exec(integral.Expressao, params)
	if err != nil {
		return r, err
	}
	result += r

	// intervalo
	for i := integral.A + step; i < integral.B; i += step {
		params[integral.Parametro] = i
		r, err := mathcat.Exec(integral.Expressao, params)
		if err != nil {
			return r, err
		}
		result += r * 2
	}

	result *= (step / 2)

	return result, nil
}

func RegraDeSimpson13(integral models.Integral, n int) (float64, error) {

	if n&1 != 0 {
		return 0.0, fmt.Errorf("n must be even n: %d", n)
	}

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]float64, 1)

	// a
	params[integral.Parametro] = integral.A
	r, err := mathcat.Exec(integral.Expressao, params)
	if err != nil {
		return r, err
	}
	result += r

	// b
	params[integral.Parametro] = integral.B
	r, err = mathcat.Exec(integral.Expressao, params)
	if err != nil {
		return r, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 2 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := mathcat.Exec(integral.Expressao, params)
		if err != nil {
			return r, err
		}
		result += r * 4
	}

	// intervalo
	for i := 2; i < n-1; i += 2 {
		params[integral.Parametro] = integral.A + float64(i)*step
		r, err := mathcat.Exec(integral.Expressao, params)
		if err != nil {
			return r, err
		}
		result += r * 2
	}

	return result * (step / 3), nil
}
