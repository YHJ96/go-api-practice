package controllers

import (
	"doc/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func HomeController(ctx *gin.Context) {
	var response models.Home
	file, _ := os.ReadFile("./mocks/home.json")
	json.Unmarshal(file, &response)
	ctx.JSON(200, response)
}

func UserController(ctx *gin.Context) {
	type Reponse struct {
		Data []models.User `json:"users"`
	}

	var response Reponse
	file, _ := os.ReadFile("./mocks/user.json")
	json.Unmarshal(file, &response)
	fmt.Println(response)

	ctx.JSON(200, response)
}