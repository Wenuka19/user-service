package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/Wenuka19/user-service/internal/adapter/handler"
	"github.com/Wenuka19/user-service/internal/adapter/repository"
	"github.com/Wenuka19/user-service/internal/config"
	"github.com/Wenuka19/user-service/internal/infrastructure"
	"github.com/Wenuka19/user-service/pkg/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := infrastructure.InitDatabase()
	infrastructure.ValidateSchema(db)

	rdb := infrastructure.InitRedis()

	userRepo := repository.NewUserRepo(db, rdb)
	userHandler := handler.NewUserHandler(userRepo)

	router := gin.Default()
	api := router.Group("/api")
	api.Use(auth.EnsureValidTokenGin())

	userGroup := api.Group("/users")
	userHandler.RegisterRoutes(userGroup)

	log.Println("User service running on port 8081")
	err := router.Run(":" + config.AppConfig.Port)

	if err != nil {
		return
	}
}
