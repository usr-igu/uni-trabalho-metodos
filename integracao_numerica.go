package metodos

import (
	"context"
	"fmt"
	"math"
	"time"
)

// RegraDosTrapeziosRepetida ...
func RegraDosTrapeziosRepetida(integral Expressao, k int) (float64, error) {
	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	wantedPrecision := math.Pow10(-k)

	expr, err := NewExpressaoAvaliavel(integral)
	if err != nil {
		return 0.0, err
	}

	lastR, err := regraDosTrapeziosRepetida(ctx, expr, 1)
	if err != nil {
		return 0.0, err
	}

	for i := 2; ; i *= 2 {
		r, err := regraDosTrapeziosRepetida(ctx, expr, i)
		if err != nil {
			if err == context.DeadlineExceeded {
				return lastR, nil
			}
			return 0.0, err
		}
		if relativeError := math.Abs(r-lastR) / math.Abs(r); relativeError < wantedPrecision {
			return r, nil
		}
		lastR = r
	}
}

// RegraDeSimpson13Repetida ...
func RegraDeSimpson13Repetida(integral Expressao, k int) (float64, error) {
	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	expr, err := NewExpressaoAvaliavel(integral)
	if err != nil {
		return 0.0, err
	}

	wantedPrecision := math.Pow10(-k)
	lastR, err := regraDeSimpson13Repetida(ctx, expr, 2)
	if err != nil {
		return 0.0, err
	}

	for i := 4; ; i *= 2 {
		r, err := regraDeSimpson13Repetida(ctx, expr, i)
		if err != nil {
			if err == context.DeadlineExceeded {
				return lastR, nil
			}
			return 0.0, err
		}

		if relativeError := math.Abs(r-lastR) / math.Abs(r); relativeError < wantedPrecision {
			return r, nil
		}
		lastR = r
	}
}

// RegraDeSimpson38Repetida ...
func RegraDeSimpson38Repetida(integral Expressao, k int) (float64, error) {
	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	expr, err := NewExpressaoAvaliavel(integral)
	if err != nil {
		return 0.0, err
	}

	wantedPrecision := math.Pow10(-k)
	lastR, err := regraDeSimpson38Repetida(ctx, expr, 3)
	if err != nil {
		return 0.0, err
	}
	for i := 6; ; i *= 3 {
		r, err := regraDeSimpson38Repetida(ctx, expr, i)
		if err != nil {
			if err == context.DeadlineExceeded {
				return lastR, nil
			}
			return 0.0, err
		}
		if relativeError := math.Abs(r-lastR) / math.Abs(r); relativeError < wantedPrecision {
			return r, nil
		}
		lastR = r
	}
}

// RegraNewtonCotes4 ...
func RegraNewtonCotes4(integral Expressao, k int) (float64, error) {
	expr, err := NewExpressaoAvaliavel(integral)
	if err != nil {
		return 0.0, err
	}
	return regraNewtonCotes4(expr)
}

func regraDosTrapeziosRepetida(ctx context.Context, integral ExpressaoAvaliavel, n int) (float64, error) {

	step := (integral.expr.A - integral.expr.B) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.expr.Parametro] = integral.expr.A
	r, err := integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.expr.Parametro] = integral.expr.B
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i++ {
		params[integral.expr.Parametro] = integral.expr.A + float64(i)*step
		r, err := integral.Avaliar(params)
		if err != nil {
			return r, err
		}
		result += r * 2.0
		select {
		case <-ctx.Done():
			return 0.0, ctx.Err()
		default:
			continue
		}
	}

	return result * (step / 2.0), nil
}

func regraDeSimpson13Repetida(ctx context.Context, integral ExpressaoAvaliavel, n int) (float64, error) {

	if n&1 != 0 {
		return 0.0, fmt.Errorf("n must be even n: %d", n)
	}

	step := (integral.expr.B - integral.expr.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.expr.Parametro] = integral.expr.A
	r, err := integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.expr.Parametro] = integral.expr.B
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 2 {
		params[integral.expr.Parametro] = integral.expr.A + float64(i)*step
		r, err := integral.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		result += r * 4
		select {
		case <-ctx.Done():
			return 0.0, ctx.Err()
		default:
			continue
		}
	}

	// intervalo
	for i := 2; i < n-1; i += 2 {
		params[integral.expr.Parametro] = integral.expr.A + float64(i)*step
		r, err := integral.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		result += r * 2
		select {
		case <-ctx.Done():
			return 0.0, ctx.Err()
		default:
			continue
		}
	}

	return result * (step / 3.0), nil
}

func regraDeSimpson38Repetida(ctx context.Context, integral ExpressaoAvaliavel, n int) (float64, error) {

	if n%3 != 0 {
		return 0.0, fmt.Errorf("n must be multiple of 3 n: %d", n)
	}

	step := (integral.expr.B - integral.expr.A) / float64(n)

	var result float64
	params := make(map[string]interface{}, 1)

	// a
	params[integral.expr.Parametro] = integral.expr.A
	r, err := integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// b
	params[integral.expr.Parametro] = integral.expr.B
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r

	// intervalo
	for i := 1; i < n; i += 3 {
		params[integral.expr.Parametro] = integral.expr.A + float64(i)*step
		r, err := integral.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		result += r * 3.0
		select {
		case <-ctx.Done():
			return 0.0, ctx.Err()
		default:
			continue
		}
	}

	// intervalo
	for i := 2; i < n-1; i += 3 {
		params[integral.expr.Parametro] = integral.expr.A + float64(i)*step
		r, err := integral.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		result += r * 3.0
		select {
		case <-ctx.Done():
			return 0.0, ctx.Err()
		default:
			continue
		}
	}

	// intervalo
	for i := 3; i < n-2; i += 3 {
		params[integral.expr.Parametro] = integral.expr.A + float64(i)*step
		r, err := integral.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		result += r * 2.0
		select {
		case <-ctx.Done():
			return 0.0, ctx.Err()
		default:
			continue
		}
	}

	return result * step * (3.0 / 8.0), nil
}

func regraNewtonCotes4(integral ExpressaoAvaliavel) (float64, error) {

	step := (integral.expr.B - integral.expr.A) / 4.0

	var result float64
	params := make(map[string]interface{}, 1)

	params[integral.expr.Parametro] = integral.expr.A
	r, err := integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r * 7.0

	params[integral.expr.Parametro] = integral.expr.A + step*1
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r * 32.0

	params[integral.expr.Parametro] = integral.expr.A + step*2
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r * 12.0

	params[integral.expr.Parametro] = integral.expr.A + step*3
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r * 32.0

	params[integral.expr.Parametro] = integral.expr.A + step*4
	r, err = integral.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	result += r * 7.0

	return result * (step * (2.0 / 45.0)), nil
}
