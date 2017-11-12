package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fuzzyqu/trabalho-metodos/metodos"
	"github.com/fuzzyqu/trabalho-metodos/models"
	"github.com/gin-gonic/gin"
)

type cacheLine struct {
	method   string
	integral models.Integral
	erro     int
}

var cache map[cacheLine]float64

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "6565"
	}

	cache = make(map[cacheLine]float64, 32)

	router := gin.Default()

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router.Static("/", "view/")

	router.POST("/trapezio/:erro", trapezio)

	router.POST("/simpson13/:erro", simpson13)

	router.POST("/simpson38/:erro", simpson38)

	router.POST("/newtoncotes4/:erro", newtoncotes4)

	s.ListenAndServe()
}

func trapezio(c *gin.Context) {
	t := c.Param("erro")
	erro, err := strconv.Atoi(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			if r, ok := cache[cacheLine{"trapezio", integral, erro}]; ok { // está na cache ?
				c.JSON(http.StatusOK, gin.H{"result": r})
			} else { // não foi computado ainda
				result, err := metodos.RegraDosTrapeziosRepetida(integral, erro)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				} else {
					cache[cacheLine{"trapezio", integral, erro}] = result
					c.JSON(http.StatusOK, gin.H{"result": result})
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}

func simpson13(c *gin.Context) {
	t := c.Param("erro")
	erro, err := strconv.Atoi(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			if r, ok := cache[cacheLine{"simpson13", integral, erro}]; ok { // está na cache ?
				c.JSON(http.StatusOK, gin.H{"result": r})
			} else { // não foi computado ainda
				result, err := metodos.RegraDeSimpson13Repetida(integral, erro)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				} else {
					cache[cacheLine{"simpson13", integral, erro}] = result
					c.JSON(http.StatusOK, gin.H{"result": result})
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}

func simpson38(c *gin.Context) {
	t := c.Param("erro")
	erro, err := strconv.Atoi(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			if r, ok := cache[cacheLine{"simpson38", integral, erro}]; ok { // está na cache ?
				c.JSON(http.StatusOK, gin.H{"result": r})
			} else { // não foi computado ainda
				result, err := metodos.RegraDeSimpson38Repetida(integral, erro)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				} else {
					cache[cacheLine{"simpson38", integral, erro}] = result
					c.JSON(http.StatusOK, gin.H{"result": result})
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}

func newtoncotes4(c *gin.Context) {
	var integral models.Integral
	if err := c.ShouldBindJSON(&integral); err == nil {
		if r, ok := cache[cacheLine{"newtoncotes4", integral, 0.0}]; ok { // está na cache ?
			c.JSON(http.StatusOK, gin.H{"result": r})
		} else { // não foi computado ainda
			result, err := metodos.RegraNewtonCotes4(integral, 0.0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				cache[cacheLine{"newtoncotes4", integral, 0.0}] = result
				c.JSON(http.StatusOK, gin.H{"result": result})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
