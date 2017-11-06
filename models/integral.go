package models

type Integral struct {
	Expressao string  `json:"expressao"`
	Parametro string  `json:"parametro"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}
