package app

import (
	"github.com/roziawilliams/npaftt/npaftt_game-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	//router.GET("/game", func(c *gin.Context) {
	//	HandleWebSocket(c.Writer, c.Request)
	//})

}
