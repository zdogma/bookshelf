package main

import (
	"bookshelf/controllers"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/:id", func(controller *gin.Context) {
		idStr := controller.Param("id")
		id, error := strconv.Atoi(idStr)

		if error != nil {
			controller.JSON(400, error)
			return
		}
		if id <= 0 {
			controller.JSON(400, gin.H{"status": "id shoild be greater than 0"})
			return
		}

		user := controllers.NewUser().Find(id)

		if user == nil || reflect.ValueOf(user).IsNil() {
			controller.JSON(404, gin.H{})
			return
		}

		controller.JSON(200, user)
	})

	router.Run(":8080")
}
