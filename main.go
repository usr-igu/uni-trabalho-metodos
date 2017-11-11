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

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "6565"
	}

	router := gin.Default()

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	cache := make(map[cacheLine]float64, 32)

	router.Static("/", "view/")

	router.POST("/trapezio/:erro", func(c *gin.Context) {
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
	})

	router.POST("/simpson13/:erro", func(c *gin.Context) {
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
	})

	router.POST("/simpson38/:erro", func(c *gin.Context) {
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
	})

	router.POST("/newtoncotes4/:erro", func(c *gin.Context) {
		t := c.Param("erro")
		erro, err := strconv.Atoi(t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			var integral models.Integral
			if err := c.ShouldBindJSON(&integral); err == nil {
				if r, ok := cache[cacheLine{"newtoncotes4", integral, erro}]; ok { // está na cache ?
					c.JSON(http.StatusOK, gin.H{"result": r})
				} else { // não foi computado ainda
					result, err := metodos.RegraNewtonCotes4(integral, erro)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					} else {
						cache[cacheLine{"newtoncotes4", integral, erro}] = result
						c.JSON(http.StatusOK, gin.H{"result": result})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	})

	s.ListenAndServe()
}
