package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"madhav/hardik/models"
	"fmt"
)

var key = "a86bcd0ea3d1d9691566aa2a37359c119619969485e418ee997f2dfb1a6c2470"

func Redirect (ctx *gin.Context) {
	key := ctx.Param("id")
	value, doesExist := maps[key]

	if !doesExist {
		value = "/error"
	}

	ctx.Redirect(http.StatusMovedPermanently, value)
}

func Modify(ctx *gin.Context) {
	var body models.AddEntry
	err := ctx.BindJSON(&body)

	if err != nil {
		fmt.Println(err)
	}
}

func Verify(ctx *gin.Context) bool {
	if ctx.Param("key") == key {
		ctx.JSON(401, gin.H{
			"message": "Nottie genius, no you cant hack :)",
		})
		return false
	}
	return true
}