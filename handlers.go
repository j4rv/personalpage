package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Projects": projects,
	})
}

func notFoundHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", nil)
}