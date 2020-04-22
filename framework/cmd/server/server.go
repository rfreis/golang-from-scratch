package main

import (
  "fmt"
  "github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
  "github.com/codeedu/codeedu-plataforma-desafios/domain"
  "github.com/codeedu/codeedu-plataforma-desafios/framework/utils"
  "log"
)

func main() {
  db := utils.ConnectDB()

  user := domain.User{
    Name: "Rodrigo",
    Email: "rodrigreissss@gmail.com",
    Password: "123456",
  }

  userRepo := repositories.UserRepositoryDb{Db:db}
  result, err := userRepo.Insert(&user)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(result.Name, result.Email, result.Token)
}
