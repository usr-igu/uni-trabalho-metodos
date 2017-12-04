package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"github.com/fuzzyqu/metodos"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {
	router := gin.Default()

	router.Static("/", "view/")

	router.POST("/simpson38/:erro", simpson38)
	router.POST("/simpson13/:erro", simpson13)
	router.POST("/trapezio/:erro", trapezio)
	router.POST("/newtoncotes4/:erro", newtoncotes4)
	router.POST("/bissecao/:erro", bissecao)
	router.POST("/posicaofalsa/:erro", posicaofalsa)
	router.POST("/newtonraphson/:erro", newtonraphson)
	router.POST("/secante/:erro", secante)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%v\n", err)
		}
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	<-exit
	log.Println("Desligando o servidor ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Erro ao desligar: ", err)
	}
	log.Println("Servidor desligado.")
}

func simpson38(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.RegraDeSimpson38Repetida(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func simpson13(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.RegraDeSimpson13Repetida(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func newtoncotes4(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.RegraNewtonCotes4(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func trapezio(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.RegraDosTrapeziosRepetida(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func bissecao(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.Bisseccao(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func posicaofalsa(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.PosicaoFalsa(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func newtonraphson(c *gin.Context) {
	expr, derivada, erro, err := parseNewtonRaphson(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.NewtonRalphson(expr, derivada, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func secante(c *gin.Context) {
	expr, erro, err := parseInput(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	result, err := metodos.Secante(expr, erro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func parseInput(c *gin.Context) (metodos.Expressao, int, error) {
	erro, err := extractError(c)
	if err != nil {
		return metodos.Expressao{}, 0, err
	}
	expr, err := extractJSON(c)
	if err != nil {
		return metodos.Expressao{}, 0, err
	}
	return expr, erro, nil
}

func extractJSON(c *gin.Context) (metodos.Expressao, error) {
	var integral metodos.Expressao
	err := c.ShouldBindJSON(&integral)
	if err != nil {
		return metodos.Expressao{}, errors.Wrap(err, "erro ao ler o json")
	}
	return integral, nil
}

func extractError(c *gin.Context) (int, error) {
	t := c.Param("erro")
	erro, err := strconv.Atoi(t)
	if err != nil {
		return 0.0, errors.Wrap(err, "valor de erro inválido")
	}
	return erro, nil
}

func parseNewtonRaphson(c *gin.Context) (metodos.Expressao, metodos.Expressao, int, error) {
	expr, erro, err := parseInput(c)
	if err != nil {
		return metodos.Expressao{}, metodos.Expressao{}, 0, err
	}
	derivada := c.Query("derivada")
	if derivada == "" {
		return metodos.Expressao{}, metodos.Expressao{}, 0, errors.New("é necessário passar a derivada de f(x)")
	}
	derivadaExpr := metodos.Expressao{Corpo: derivada}
	return expr, derivadaExpr, erro, nil
}
