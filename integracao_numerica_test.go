package metodos

import (
	"testing"
)

var expr Expressao
var k int

func init() {
	expr = Expressao{
		A:         1,
		B:         4,
		Corpo:     "x**2",
		Parametro: "x",
	}
	k = 5
}

func BenchmarkRegraDeSimpson38Repetida(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		v, _ := RegraDeSimpson38Repetida(expr, k)
		r = v
	}
	_ = r
}


func BenchmarkRegraDeSimpson13Repetida(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		v, _ := RegraDeSimpson13Repetida(expr, k)
		r = v
	}
	_ = r
}


func BenchmarkRegraDosTrapeziosRepetida(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		v, _ := RegraDosTrapeziosRepetida(expr, k)
		r = v
	}
	_ = r
}

func BenchmarkRegraNewtonCotes4(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		v, _ := RegraNewtonCotes4(expr, k)
		r = v
	}
	_ = r
}


