package main

import (
	"api/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	controllers.TweetRouter(app)

	app.Run()

}