package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	c "github.com/ockibagusp/hello/controllers"
	"github.com/ockibagusp/hello/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func truncateUsers(db *gorm.DB) {
	// db.Exec("TRUNCATE users")
}

func setupRouter() (router *echo.Echo) {
	router = echo.New()

	// controllers init
	controllers := c.Controller{DB: db}

	router.GET("/users", controllers.Users).Name = "users"

	return
}

const (
	userJSON  = `{"id":1,"name":"Jon Snow"}`
	usersJSON = `[{"id":1,"name":"Jon Snow"}]`
)

func TestEchoHandler(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()

	// HandlerFunc
	e.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	req := httptest.NewRequest(http.MethodGet, "/ok", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(http.StatusOK, rec.Code)
	assert.Equal("OK", rec.Body.String())
}

func TestUserControllerAPI(t *testing.T) {
	assert := assert.New(t)

	truncateUsers(db)

	tx := db.Begin()

	defer tx.Rollback()

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	assert.True(true)
}

func TestUserController(t *testing.T) {
	assert := assert.New(t)

	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()

	fmt.Println("req -> ", request)

	router := echo.New()

	// controllers init
	controllers := c.Controller{DB: db}
	router.GET("/users", controllers.Users)
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	fmt.Println("req -> ", req)
	fmt.Println("rec -> ", rec)

	assert.Equal(http.StatusOK, rec.Code)
	assert.Equal("OK", rec.Body.String())

	// ???
	// req := httptest.NewRequest(method, path, nil)
	// rec := httptest.NewRecorder()
	// e.ServeHTTP(rec, req)
	// return rec.Code, rec.Body.String()

	// u := router.NewContext(request, recorder)

	// h := &handler{mockDB}

	// // Assertions
	// if assert.NoError(h.getUser(c)) {
	// 	assert.Equal(http.StatusOK, recorder.Code)
	// 	assert.Equal(userJSON, recorder.Body.String())
	// }
	response := recorder.Result()
	defer response.Body.Close()

	// err := controllers.Users(u)

	// if assert.NoError(controllers.UsersAPI(u)) {
	// 	assert.Equal(http.StatusOK, response.StatusCode)
	// }

	// assert.Equal(200, response.StatusCode)

	// // h := &handler{mockDB}

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)

	// //assert.Equal(200, int(responseBody["code"].(float64)))
	// assert.Equal("OK", responseBody["status"])
}

func TestReadUserController(t *testing.T) {
	assert := assert.New(t)

	// // func trucate
	// db.Exec("TRUNCATE users")

	// f := make(url.Values)
	// f.Set("name", "Jon Snow")
	// f.Set("email", "jon@labstack.com")
	// request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	// request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	e := echo.New()

	// user.FindAll()
	user, _ := models.User{}.FindByID(db, 1)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodGet, "/users/read/"+strconv.Itoa(int(user.Model.ID)), requestBody)
	fmt.Println("request -> ", request)

	recorder := httptest.NewRecorder()

	c := e.NewContext(request, recorder)
	c.SetPath("/users/read/" + strconv.Itoa(int(user.Model.ID)))
	fmt.Println("c -> ", c.Path())
	// c.SetParamNames("email")
	// c.SetParamValues("jon@labstack.com")

	response := recorder.Result()

	assert.Equal(200, response.StatusCode)

	// h := &handler{mockDB}

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	//assert.Equal(200, int(responseBody["code"].(float64)))
	assert.Equal("OK", responseBody["status"])

}
