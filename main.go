package main

import (
	"net/http"
	"strconv"

	"github.com/fuzzyqu/trabalho-metodos/metodos"
	"github.com/fuzzyqu/trabalho-metodos/models"
	"github.com/gin-gonic/gin"
)

type cacheLine struct {
	method   string
	integral models.Integral
	n        int
}

func main() {
	router := gin.Default()

	cache := make(map[cacheLine]float64, 32)

	router.POST("/trapezio/:n", func(c *gin.Context) {
		n := c.Param("n")
		nint, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			var integral models.Integral
			if err := c.ShouldBindJSON(&integral); err == nil {
				if r, ok := cache[cacheLine{"trapezio", integral, nint}]; ok { // está na cache ?
					c.JSON(http.StatusOK, gin.H{"result": r})
				} else { // não foi computado ainda
					result, err := metodos.RegraDosTrapeziosRepetida(integral, nint)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					} else {
						cache[cacheLine{"trapezio", integral, nint}] = result
						c.JSON(http.StatusOK, gin.H{"result": result})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	})

	router.POST("/simpson13/:n", func(c *gin.Context) {
		n := c.Param("n")
		nint, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			var integral models.Integral
			if err := c.ShouldBindJSON(&integral); err == nil {
				if r, ok := cache[cacheLine{"simpson13", integral, nint}]; ok { // está na cache ?
					c.JSON(http.StatusOK, gin.H{"result": r})
				} else { // não foi computado ainda
					result, err := metodos.RegraDeSimpson13Repetida(integral, nint)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					} else {
						cache[cacheLine{"simpson13", integral, nint}] = result
						c.JSON(http.StatusOK, gin.H{"result": result})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	})

	router.POST("/simpson38/:n", func(c *gin.Context) {
		n := c.Param("n")
		nint, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			var integral models.Integral
			if err := c.ShouldBindJSON(&integral); err == nil {
				if r, ok := cache[cacheLine{"simpson38", integral, nint}]; ok { // está na cache ?
					c.JSON(http.StatusOK, gin.H{"result": r})
				} else { // não foi computado ainda
					result, err := metodos.RegraDeSimpson38Repetida(integral, nint)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					} else {
						cache[cacheLine{"simpson38", integral, nint}] = result
						c.JSON(http.StatusOK, gin.H{"result": result})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	})

	router.POST("/newtoncotes4/", func(c *gin.Context) {
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			if r, ok := cache[cacheLine{"newtoncotes4", integral, 0}]; ok { // está na cache ?
				c.JSON(http.StatusOK, gin.H{"result": r})
			} else { // não foi computado ainda
				result, err := metodos.RegraNewtonCotes4(integral)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				} else {
					cache[cacheLine{"newtoncotes4", integral, 0}] = result
					c.JSON(http.StatusOK, gin.H{"result": result})
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run(":6565")
}
