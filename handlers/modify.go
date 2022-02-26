package handlers

import (
	"fmt"
	"madhav/hardik/models"
	"github.com/gin-gonic/gin"
)

var maps = map[string]string{}

func Add(ctx *gin.Context) {
	Verify(ctx)

	var entry models.AddEntry
	err := ctx.BindJSON(&entry)
	
	if err != nil {
		fmt.Println(err)
	}
	maps[entry.Key] = entry.Value
}

func Delete(ctx *gin.Context) {
	Verify(ctx)

	var entry models.DeleteEntry
	err := ctx.BindJSON(&entry)

	if err != nil {
		fmt.Println(err)
	}
	delete(maps, entry.Key)
}