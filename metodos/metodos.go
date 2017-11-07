package metodos

import (
	"fmt"

	"github.com/fuzzyqu/trabalho-metodos/models"
	"github.com/soudy/mathcat"
)

func RegraDosTrapezios(integral models.Integral, n int64) (float64, error) {

	deltaX := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]float64, 8)

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
	for i := integral.A + deltaX; i < integral.B; i += deltaX {
		params[integral.Parametro] = i
		r, err := mathcat.Exec(integral.Expressao, params)
		if err != nil {
			return r, err
		}
		result += r * 2
	}

	result *= (deltaX / 2)

	return result, nil
}

func RegraDeSimpson13(integral models.Integral, n int64) (float64, error) {

	if n&1 != 0 {
		return 0.0, fmt.Errorf("n must be even n: %d", n)
	}

	step := (integral.B - integral.A) / float64(n)

	var result float64
	params := make(map[string]float64, 8)

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
	even := true
	for i := integral.A + step; i < integral.B; i += step {
		params[integral.Parametro] = i
		r, err := mathcat.Exec(integral.Expressao, params)
		if err != nil {
			return r, err
		}
		if even {
			result += r * 4
			even = !even
		} else {
			result += r * 2
			even = !even
		}

	}

	result *= (step / 3)

	return result, nil
}
