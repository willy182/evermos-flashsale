package main

import (
	"fmt"

	configEnv "github.com/joho/godotenv"
	"github.com/willy182/evermos-flashsale/config"
	"github.com/willy182/evermos-flashsale/delivery"
	"github.com/willy182/evermos-flashsale/repository"
	"github.com/willy182/evermos-flashsale/usecase"
)

// main function, entry function for this services
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %v", r)
		}
	}()

	// load environment variables
	err := configEnv.Load(".env")
	if err != nil {
		panic(err)
	}

	conf := config.Load()
	repo := repository.NewRepositorySQL(conf.DB)

	useCase := usecase.NewUseCase(repo)

	handler := delivery.NewDelivery(useCase)
	handler.Checkout()
	return
}
