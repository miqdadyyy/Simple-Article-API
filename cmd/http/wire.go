//go:build wireinject
// +build wireinject

package main

import (
	"os"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/miqdadyyy/jetdevs-assesment/internal/handler"
	articlerp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/implementation"
	commentrp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/implementation"
	articleuc "github.com/miqdadyyy/jetdevs-assesment/internal/usecase/article/implementation"
	"github.com/miqdadyyy/jetdevs-assesment/pkg/config"
)

var (
	PkgSet = wire.NewSet(
		ProvideConfig,
		ProvideDatabase,
	)

	RepoSet = wire.NewSet(
		articlerp.NewArticleRepository,
		commentrp.NewCommentRepository,
	)

	UsecaseSet = wire.NewSet(
		articleuc.NewArticleUsecase,
	)

	AllSet = wire.NewSet(
		PkgSet,
		RepoSet,
		UsecaseSet,
		handler.NewHandler,
	)
)

func InitHandler() *handler.Handler {
	wire.Build(AllSet)
	return nil
}

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
