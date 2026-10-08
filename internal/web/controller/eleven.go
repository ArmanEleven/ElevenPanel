package controller

import (
	"net/http"

	v1 "github.com/mhsanaei/3x-ui/v3/internal/eleven/api/v1"

	"github.com/gin-gonic/gin"
)

type ElevenController struct{}

func NewElevenController(g *gin.RouterGroup) *ElevenController {
	c := &ElevenController{}

	eleven := g.Group("/api/v1/eleven")
	eleven.GET("/health", c.health)

	return c
}

func (c *ElevenController) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, v1.HealthResponse{
		Status:  "ok",
		Version: "0.1.0",
	})
}
