package metodos

import (
	"context"
	"math"
	"time"

	"github.com/pkg/errors"
)

func Bisseccao(funcao Expressao, k int) (float64, error) {
	expr, err := NewExpressaoAvaliavel(funcao)
	if err != nil {
		return 0.0, err
	}

	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	precisaoEsperada := math.Pow10(-k)

	n := math.Ceil((math.Log10(funcao.B-funcao.A) - math.Log10(precisaoEsperada)) / math.Log10(2))
	return bisseccao(ctx, expr, int(n))
}

func PosicaoFalsa(funcao Expressao, k int) (float64, error) {
	expr, err := NewExpressaoAvaliavel(funcao)
	if err != nil {
		return 0.0, err
	}

	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	return posicaoFalsa(ctx, expr, k)
}

func NewtonRalphson(funcao, derivada Expressao, k int) (float64, error) {
	expr, err := NewExpressaoAvaliavel(funcao)
	if err != nil {
		return 0.0, err
	}

	derivadaExpr, err := NewExpressaoAvaliavel(derivada)
	if err != nil {
		return 0.0, err
	}

	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	return newtonRalphson(ctx, expr, derivadaExpr, k)
}

func Secante(funcao Expressao, k int) (float64, error) {
	expr, err := NewExpressaoAvaliavel(funcao)
	if err != nil {
		return 0.0, err
	}

	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	//precisaoEsperada := math.Pow10(-k)

	return secante(ctx, expr, k)
}

func bisseccao(ctx context.Context, funcao ExpressaoAvaliavel, n int) (float64, error) {
	params := make(map[string]interface{}, 1)

	params[funcao.expr.Parametro] = funcao.expr.A
	fa, err := funcao.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	if fa == 0 {
		return funcao.expr.A, nil
	}

	params[funcao.expr.Parametro] = funcao.expr.B
	fb, err := funcao.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	if fb == 0 {
		return funcao.expr.B, nil
	}

	if fa*fb > 0 {
		return 0.0, errors.New("os sinais de f(a) e f(b) não são opostos")
	}

	a := funcao.expr.A
	b := funcao.expr.B
	var p, fp float64

	for i := 0; i < n; i++ {
		p = (a + b) / 2.0
		params[funcao.expr.Parametro] = p
		fp, err = funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		if fa*fp < 0 { //fa e fp tem sinais opostos
			b = p
		} else { // fp e fb tem sinais opostos
			a = p
		}
		select {
		case <-ctx.Done():
			return p, ctx.Err()
		default:
			continue
		}
	}
	return p, nil
}

func posicaoFalsa(ctx context.Context, funcao ExpressaoAvaliavel, k int) (float64, error) {
	params := make(map[string]interface{}, 1)
	precisaoEsperada := math.Pow10(-k)
	params[funcao.expr.Parametro] = funcao.expr.A
	fa, err := funcao.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	if fa == 0 {
		return funcao.expr.A, nil
	}

	params[funcao.expr.Parametro] = funcao.expr.B
	fb, err := funcao.Avaliar(params)
	if err != nil {
		return 0.0, err
	}
	if fb == 0 {
		return funcao.expr.B, nil
	}

	if fa*fb > 0 {
		return 0.0, errors.New("os sinais de f(a) e f(b) não são opostos")
	}

	a := funcao.expr.A
	b := funcao.expr.B
	for {
		xk := a*fb - b*fa/(fb-fa)
		params[funcao.expr.Parametro] = xk
		fxk, err := funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		if math.Abs(fxk) < precisaoEsperada {
			return xk, nil
		}

		if fa*fxk > 0 {
			a = xk
		} else {
			b = xk
		}

		params[funcao.expr.Parametro] = a
		fa, err = funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		params[funcao.expr.Parametro] = b
		fb, err = funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
	}
}

func newtonRalphson(ctx context.Context, funcao, derivada ExpressaoAvaliavel, k int) (float64, error) {
	params := make(map[string]interface{}, 1)
	//precisaoEsperada := math.Pow10(-k)
	var xn float64 = 1
	lastxn := xn
	for {
		params[funcao.expr.Parametro] = xn
		fx, err := funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		dx, err := derivada.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		xn = xn - fx/dx
		if xn == lastxn {
			return xn, nil
		}
		lastxn = xn
	}
}

func secante(ctx context.Context, funcao ExpressaoAvaliavel, k int) (float64, error) {
	params := make(map[string]interface{}, 1)
	precisaoEsperada := math.Pow10(-k)
	var xa float64 = 0
	var xb float64 = 1
	for {
		params[funcao.expr.Parametro] = xa
		fxa, err := funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		params[funcao.expr.Parametro] = xb
		fxb, err := funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}

		fxr := ((xa * fxb) - (xb * fxa)) / (fxb - fxa)

		params[funcao.expr.Parametro] = fxr
		r, err := funcao.Avaliar(params)
		if err != nil {
			return 0.0, err
		}
		if math.Abs(r) < precisaoEsperada {
			return fxr, nil
		}

		xa = fxr
	}
}
