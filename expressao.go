package metodos

import (
	"github.com/pkg/errors"

	"math"

	"github.com/Knetic/govaluate"
)

// Expressao ...
type Expressao struct {
	Corpo     string  `json:"corpo"`
	Parametro string  `json:"parametro"`
	A         float64 `json:"a,string"`
	B         float64 `json:"b,string"`
}

type ExpressaoAvaliavel struct {
	eval *govaluate.EvaluableExpression
	expr Expressao
}

func (e *ExpressaoAvaliavel) Avaliar(params map[string]interface{}) (float64, error) {
	if e == nil {
		return 0.0, errors.New("tentativa de avaliar uma expressão nula")
	}
	return evaluateEvaluableExpression(e.eval, params)
}

func NewExpressaoAvaliavel(expr Expressao) (ExpressaoAvaliavel, error) {
	e, err := newEvaluableExpression(expr.Corpo)
	if err != nil {
		return ExpressaoAvaliavel{}, errors.Wrap(err, "expressão inválida")
	}
	return ExpressaoAvaliavel{e, expr}, nil
}

func evaluateEvaluableExpression(expr *govaluate.EvaluableExpression, params map[string]interface{}) (float64, error) {
	params["e"] = math.E
	params["pi"] = math.Pi
	t, err := expr.Evaluate(params)
	if err != nil {
		return 0.0, err
	}
	result, ok := t.(float64)
	if !ok {
		return 0.0, errors.New("impossível converter o resultado para float64")
	}
	return result, nil
}

func newEvaluableExpression(expr string) (*govaluate.EvaluableExpression, error) {
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
			if c == math.NaN() {
				return 0.0, nil
			}
			return (float64)(c), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
			c := math.Log10(args[0].(float64))
			if c == math.NaN() {
				return 0.0, nil
			}
			return (float64)(c), nil
		},
		"logn": func(args ...interface{}) (interface{}, error) {
			c := math.Log(args[0].(float64))
			if c == math.NaN() {
				return 0.0, nil
			}
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
