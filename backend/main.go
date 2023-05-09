package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/zerobugdebug/sankortti/backend/internal/db"
	"github.com/zerobugdebug/sankortti/backend/internal/game"
	"github.com/zerobugdebug/sankortti/backend/internal/websocket"
)

var (
	dynamoDBClient *dynamodb.DynamoDB
	gameHandler    *game.GameHandler
	dbHandler      *db.DBHandler
	wsHandler      *websocket.WebsocketHandler
)

func init() {
	// Initialize AWS Session
	sess := session.Must(session.NewSession())

	// Initialize DynamoDB client
	dynamoDBClient = dynamodb.New(sess, aws.NewConfig())

	// Initialize game, DB, and WebSocket handlers
	gameHandler = game.NewGameHandler()
	dbHandler = db.NewDBHandler(dynamoDBClient)
	wsHandler = websocket.NewWebsocketHandler()
}

// Lambda function handler
func handler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Process WebSocket events using the WebSocket handler
	response, err := wsHandler.Handle(ctx, request, gameHandler, dbHandler)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return response, nil
}

func main() {
	// Start the Lambda function handler
	lambda.Start(handler)
}
