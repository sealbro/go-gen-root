// Composition root code generated by github.com/sealbro/go-gen-root. DO NOT EDIT.
package main

import (
	businesslogic "github.com/sealbro/go-gen-root/examples/app-with-deps/internal/businesslogic"
	services "github.com/sealbro/go-gen-root/examples/app-with-deps/internal/services"
)

type App struct {
	serviceTwo        *services.ServiceTwo
	calculatorService *businesslogic.CalculatorService
	serviceOne        *services.ServiceOne
}

func NewApp() *App {
	serviceTwo := services.NewServiceTwo()
	serviceOne := services.NewServiceOne()
	calculatorService := businesslogic.NewCalculatorService(serviceOne, serviceTwo)
	return &App{
		calculatorService: calculatorService,
		serviceOne:        serviceOne,
		serviceTwo:        serviceTwo,
	}
}