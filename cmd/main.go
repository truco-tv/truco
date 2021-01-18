package main

import (
	"context"

	"strconv"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/truco-tv/truco"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	deck := truco.NewDeck()
	value, _ := strconv.Atoi(request.QueryStringParameters["value"])
	suit, _ := strconv.Atoi(request.QueryStringParameters["suit"])
	vira := &truco.Card{Value: value, Suit: suit}
	cards := truco.NewCheatSheet(vira, deck)
	cardsJson, _ := json.Marshal(cards)

	return events.APIGatewayProxyResponse{Body: string(cardsJson), StatusCode: 200}, nil
}
func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/cheatsheet", func(c *gin.Context) {
	// 	deck := truco.NewDeck()
	// 	vira := &truco.Card{Value: 7, Suit: truco.Cup}

	// 	cards := truco.NewCheatSheet(vira, deck)
	// 	c.JSON(200, cards)
	// })

	// r.Run()
	lambda.Start(HandleRequest)
}
