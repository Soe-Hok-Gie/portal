package main

import (
	"fmt"
	"log"
	"medsos/app"
	"medsos/controller"
	"medsos/repository"
	"medsos/service"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	//setENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	userDB := os.Getenv("DB_USER")
	passDB := os.Getenv("DB_PASS")
	hostDB := os.Getenv("DB_HOST")
	portDB := os.Getenv("DB_PORT")
	nameDB := os.Getenv("DB_NAME")

	fmt.Println("dsn:", userDB, passDB, hostDB, portDB, nameDB)

	//setDB
	db := app.NewDB(userDB, passDB, hostDB, portDB, nameDB)

	//pattern
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	//patern post
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	r := mux.NewRouter()
	r.HandleFunc("/user", userController.Create).Methods("POST")
	r.HandleFunc("/user/{id}", userController.Update).Methods("PUT")
	r.HandleFunc("/user/{id}", userController.FindById).Methods("GET")
	r.HandleFunc("/user/{id}", userController.Delete).Methods("DELETE")
	r.HandleFunc("/user", userController.FindAll).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
