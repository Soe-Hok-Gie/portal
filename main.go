package portal

import (
	"fmt"
	"log"
	"medsos/app"
	"os"

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
}
