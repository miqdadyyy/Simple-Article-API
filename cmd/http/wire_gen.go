// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/miqdadyyy/jetdevs-assesment/internal/handler"
	"github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/implementation"
	implementation2 "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/implementation"
	implementation3 "github.com/miqdadyyy/jetdevs-assesment/internal/usecase/article/implementation"
	"github.com/miqdadyyy/jetdevs-assesment/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Injectors from wire.go:

func InitHandler() *handler.Handler {
	config := ProvideConfig()
	db := ProvideDatabase(config)
	articleRepositoryInterface := implementation.NewArticleRepository(db)
	commentRepositoryInterface := implementation2.NewCommentRepository(db)
	articleUsecaseInterface := implementation3.NewArticleUsecase(articleRepositoryInterface, commentRepositoryInterface)
	handlerHandler := handler.NewHandler(articleUsecaseInterface)
	return handlerHandler
}

// wire.go:

var (
	PkgSet = wire.NewSet(
		ProvideConfig,
		ProvideDatabase,
	)

	RepoSet = wire.NewSet(implementation.NewArticleRepository, implementation2.NewCommentRepository)

	UsecaseSet = wire.NewSet(implementation3.NewArticleUsecase)

	AllSet = wire.NewSet(
		PkgSet,
		RepoSet,
		UsecaseSet, handler.NewHandler,
	)
)

func ProvideConfig() config.Config {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "local"
	}

	return config.GetConfig(environment)
}

func ProvideDatabase(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}
