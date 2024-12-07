package controller

import (
	"project-ta/service"
)

type LabaRugiController struct {
	LabaRugiService service.LabaRugiServiceInj
}

type LabaRugiControllerInj interface{}

func NewLabaRugiController(lrs service.LabaRugiServiceInj) LabaRugiControllerInj {
	return LabaRugiController{
		LabaRugiService: lrs,
	}
}
