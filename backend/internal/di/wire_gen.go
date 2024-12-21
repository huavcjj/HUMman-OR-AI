// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"Bot-or-Not/internal/app/handler"
	"Bot-or-Not/internal/app/router"
	"Bot-or-Not/internal/app/service"
	"Bot-or-Not/internal/infra/database"
	"Bot-or-Not/internal/infra/repository"
)

// Injectors from wire.go:

func New() *router.Root {
	db := database.New()
	iPlayerRepository := repository.NewPlayerRepository(db)
	iPlayerService := service.NewPlayerService(iPlayerRepository)
	iPlayerHandler := handler.NewPlayerHandler(iPlayerService)
	root := router.New(iPlayerHandler)
	return root
}