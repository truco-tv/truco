package main

import (
	"context"
	"strings"

	"strconv"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/truco-tv/truco"
)

/*
GOOS=linux go build main.go
zip function.zip main
*/

type TrucoResponse struct {
	Cards       []*truco.Card
	HasEnvido   bool
	HasFlor     bool
	EnvidoValue int
	FlorValue   int
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	deck := truco.NewDeck()
	value, _ := strconv.Atoi(request.QueryStringParameters["value"])
	suit, _ := strconv.Atoi(request.QueryStringParameters["suit"])
	vira := &truco.Card{Value: value, Suit: suit}
	cards := truco.NewCheatSheet(vira, deck)
	player := truco.NewPlayer()

	for _, v := range request.MultiValueQueryStringParameters["cards"] {
		cardValue := strings.Split(v, "-")
		value, _ := strconv.Atoi(cardValue[0])
		suit, _ := strconv.Atoi(cardValue[1])

		card := truco.NewCard(value, suit)
		card = player.TransformPerica(vira, card)
		card = player.TransformPerico(vira, card)

		player.AppendCard(card)
	}

	trucoResponse := &TrucoResponse{Cards: cards, HasEnvido: player.HasEnvido(), EnvidoValue: player.EnvidoValue}
	trucoResponseJson, _ := json.Marshal(trucoResponse)

	return events.APIGatewayProxyResponse{Body: string(trucoResponseJson), StatusCode: 200}, nil
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
