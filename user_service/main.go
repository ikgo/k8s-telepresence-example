package main;

import (
	"os"
	"log"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	port :=os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT not set")
	}
	e := echo.New()
	e.GET("/", getUsers)
	e.Logger.Fatal(e.Start(":" + port))
}

func getUsers(ctx echo.Context) error {
	u := make([]map[string]interface{}, 4)
	u[0] = map[string]interface{}{"name": "Matt"}
	u[1] = map[string]interface{}{"name": "Steve"}
	u[2] = map[string]interface{}{"name": "Greg"}
	u[3] = map[string]interface{}{"name": "Mark"}
	return ctx.JSON(http.StatusOK, u)
}