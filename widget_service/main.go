package main;

import (
	"os"
	//"log"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
	port :=os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT not set")
	}
	e := echo.New()
	e.GET("/", getWidgets)
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Error(e.Start(":" + port))

}

func getWidgets(ctx echo.Context) error {
	users, err := getUsers()
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	w := make([]map[string]interface{}, 3)
	w[0] = map[string]interface{}{"name": "Thingamajig"}
	w[1] = map[string]interface{}{"name": "Thingamabob"}
	w[2] = map[string]interface{}{"name": "Plumbus"}

	var rslt = make(map[string]interface{})
	rslt["users"] = users
	rslt["widgets"] = w

	return ctx.JSON(http.StatusOK, rslt)
}

func getUsers() ([]map[string]interface{}, error) {
	resp, err := http.Get("http://user-service.default:8080/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsn = make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}