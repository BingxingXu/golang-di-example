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

func (s *AppController) FindAll(c *gin.Context) {
	apps := s.AppService.FindAll()

	c.JSON(http.StatusOK, gin.H{"apps": ToAppDTOs(apps)})
}
