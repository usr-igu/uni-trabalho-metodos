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
	n        int64
}

func main() {
	router := gin.Default()

	cache := make(map[cacheLine]float64, 128)

	router.POST("/trapezio/:n", func(c *gin.Context) {
		n := c.Param("n")
		nint, err := strconv.ParseInt(n, 0, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			if r, ok := cache[cacheLine{"trapezio", integral, nint}]; ok { // está na cache ?
				c.JSON(http.StatusOK, gin.H{"result": r})
			} else { // não foi computado ainda
				result, err := metodos.RegraDosTrapezios(integral, nint)
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
	})

	router.POST("/simpson13/:n", func(c *gin.Context) {
		n := c.Param("n")
		nint, err := strconv.ParseInt(n, 0, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			if r, ok := cache[cacheLine{"simpson13", integral, nint}]; ok { // está na cache ?
				c.JSON(http.StatusOK, gin.H{"result": r})
			} else { // não foi computado ainda
				result, err := metodos.RegraDeSimpson13(integral, nint)
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
	})

	router.Run(":6565")
}
