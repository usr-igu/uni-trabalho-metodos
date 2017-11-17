package metodos

import (
	"math"

	"github.com/Knetic/govaluate"
)

func evaluateExpression(expr *govaluate.EvaluableExpression, params map[string]interface{}) (float64, error) {
	params["e"] = math.E
	params["pi"] = math.Pi
	params["phi"] = math.Phi
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
			if c == math.NaN() {
				return 0.0, nil
			}
			return (float64)(c), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			c := math.Log10(args[0].(float64))
			if c == math.NaN() {
				return 0.0, nil
			}
			return (float64)(c), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
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