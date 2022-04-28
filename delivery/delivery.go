package delivery

import (
	"assignment3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpDelivery struct {
	models.StatusCase
}

func NewHttpDelivery(r *gin.Engine, UseCase models.StatusCase) {
	handler := HttpDelivery{
		StatusCase: UseCase,
	}

	r.GET("/", handler.updates)
}

func (h *HttpDelivery) updates(c *gin.Context) {
	data, err := h.UpdateStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	datas := models.Danger{
		Water:       data.Water,
		Wind:        data.Wind,
		StatusWater: data.StatusWater,
		StatusWind:  data.StatusWind,
	}

	c.HTML(
		http.StatusOK,
		"index.html", gin.H{
			"title": "Assignment 3",
			"datas": datas,
		},
	)

}
