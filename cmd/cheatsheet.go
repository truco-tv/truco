package main

import (
	"github.com/gin-gonic/gin"
	"github.com/truco-tv/truco"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cheatsheet", func(c *gin.Context) {
		deck := truco.NewDeck()
		vira := &truco.Card{Value: 7, Suit: truco.Cup}

		cards := truco.NewCheatSheet(vira, deck)
		c.JSON(200, cards)
	})

	r.Run()
}
