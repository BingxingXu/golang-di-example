package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type AppController struct {
	AppService AppService
}

func ProviderAppController(s AppService) AppController {
	return AppController{AppService: s}
}

// findAll godoc
// @Summary 123
// @Description 123
// @Tags apps
// @Accept json
// @Produce json
// @param order path string true "Order Type Desc or INC"
// @Success 200 {array} app.AppDTO
// @Router /apps [get]
func (s *AppController) FindAll(c *gin.Context) {
	apps := s.AppService.FindAll()

	c.JSON(http.StatusOK, gin.H{"apps": ToAppDTOs(apps)})
}
