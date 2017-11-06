package main

import (
	"net/http"
	"strconv"

	"github.com/fuzzyqu/trabalho-metodos/metodos"
	"github.com/fuzzyqu/trabalho-metodos/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/trapezio/:n", func(c *gin.Context) {
		n := c.Param("n")
		nInteger, err := strconv.ParseInt(n, 0, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		var integral models.Integral
		if err := c.ShouldBindJSON(&integral); err == nil {
			result, err := metodos.RegraDosTrapezios(integral, nInteger)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": result})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run(":6565")
}
