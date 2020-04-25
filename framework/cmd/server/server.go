package main

import (
	"fmt"
	"log"

	"github.com/andrewmak/poc_golang/application/repositories"
	"github.com/andrewmak/poc_golang/domain"
	"github.com/andrewmak/poc_golang/framework/utils"
	_ "github.com/lib/pq"
)

func main() {
	db := utils.ConnectDb()
	user := domain.User{
		Name:     "andrew",
		Email:    "andrew@teste.com",
		Password: "1234",
	}

	useRepo := repositories.UserRepositoryDb{Db: db}
	result, err := useRepo.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
