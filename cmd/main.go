package main
import (
	"github.com/spear-app/spear-go/pkg/driver"
	"github.com/joho/godotenv"
	"log"
)
func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbconn:= driver.GetDbConnetion()
	print(dbconn)
}