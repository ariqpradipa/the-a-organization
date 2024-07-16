package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "true",
		})
	})
	return r
}

func TestHealthCheckRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"true"}`, w.Body.String())
}
