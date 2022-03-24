package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)
type Response struct {
	Role  string `json:"given_name"`
	Email string `json:"email"`
	Id    string `json:"nickname"`
}

func authMid(c *gin.Context) {

	url := "https://dev-rgmfg73e.us.auth0.com/userinfo"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", c.Request.Header["Authorization"][0])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := Response{}
	json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Email)
	fmt.Println(response.Role)
	fmt.Println(response.Id)
	c.Params = []gin.Param{
		{
			Key:   "email",
			Value: response.Email,
		},
		{
			Key:   "role",
			Value: response.Role,
		},
		{
			Key:   "id",
			Value: response.Id,
		},
		
	}
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://87cb609bebb8450283fb75d18f14aa28@o1176298.ingest.sentry.io/6273809",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	secureGroup := r.Group("/secure/", authMid)

	// Routes
	secureGroup.GET("/reports", controllers.FindReports)
	secureGroup.GET("/reports/:id", controllers.FindReport)
	secureGroup.POST("/reports", controllers.CreateReport)
	secureGroup.PATCH("/reports/:id", controllers.UpdateReport)
	secureGroup.DELETE("/reports/:id", controllers.DeleteReport)

	// Run the server
	r.Run()
}
