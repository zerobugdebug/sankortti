package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

const TableName = "GameStates"

// Card represents a card in the game.
type Card struct {
	ID      string `json:"id"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	Power   int    `json:"power"`
	Speed   int    `json:"speed"`
}

// GameState holds the current state of the game.
type GameState struct {
	GameID        string `json:"gameId"`
	Player1       string `json:"player1"`
	Player2       string `json:"player2"`
	Player1Hand   []Card `json:"player1Hand"`
	Player2Hand   []Card `json:"player2Hand"`
	Player1HP     int    `json:"player1HP"`
	Player2HP     int    `json:"player2HP"`
	CurrentPlayer string `json:"currentPlayer"`
}

// GameAction represents an action performed by a player.
type GameAction struct {
	GameID   string `json:"gameId"`
	PlayerID string `json:"playerId"`
	Action   string `json:"action"`
	CardID   string `json:"cardId"`
}

// HandleGameAction processes a game's action and updates the game state.
func HandleGameAction(ctx context.Context, request events.APIGatewayWebsocketProxyRequest, dynamoDB *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	var action GameAction
	err := json.Unmarshal([]byte(request.Body), &action)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid action format",
		}, nil
	}

	if action.Action == "new_game" {
		handleNewGame(ctx, request, dynamoDB)

	}

	// Load the game state from DynamoDB.
	gameState, err := loadGameState(dynamoDB, action.GameID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error loading game state",
		}, nil
	}

	// Update the game state based on the action.
	switch action.Action {
	case "card_clicked":
		handleCardClicked(gameState, action)
	case "card_confirmed":
		handleCardConfirmed(gameState, action)
	case "new_game":
		handleNewGame(ctx, request, dynamoDB)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid action",
		}, nil
	}

	// Save the updated game state to DynamoDB.
	err = saveGameState(dynamoDB, gameState)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error saving game state",
		}, nil
	}

	// Create an API Gateway management API client
	sess := session.Must(session.NewSession())
	apiGatewayManagementApi := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(request.RequestContext.DomainName+"/"+request.RequestContext.Stage))

	// Send a message to the frontend with the confirmation
	responseData, err := json.Marshal(map[string]string{
		"command": "action_confirmed",
		"cardId":  action.CardID,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error marshaling response data",
		}, nil
	}

	_, err = apiGatewayManagementApi.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(request.RequestContext.ConnectionID),
		Data:         responseData,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error sending WebSocket message",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Game action processed",
	}, nil
}

// loadGameState is a function to load the game state from DynamoDB.
func loadGameState(dynamoDB *dynamodb.DynamoDB, gameID string) (*GameState, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"GameID": {
				S: aws.String(gameID),
			},
		},
	}

	result, err := dynamoDB.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, fmt.Errorf("Game state not found for GameID: %s", gameID)
	}

	var gameState GameState
	err = dynamodbattribute.UnmarshalMap(result.Item, &gameState)
	if err != nil {
		return nil, err
	}

	return &gameState, nil
}

// saveGameState is a function to save the game state to DynamoDB.
func saveGameState(dynamoDB *dynamodb.DynamoDB, gameState *GameState) error {
	item, err := dynamodbattribute.MarshalMap(gameState)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      item,
	}

	_, err = dynamoDB.PutItem(input)
	return err
}

func handleCardClicked(gameState *GameState, action GameAction) {
	fmt.Printf("action: %v\n", action)
	// Implement your logic for handling the card_clicked action.
	// ...
}

func handleCardConfirmed(gameState *GameState, action GameAction) {
	fmt.Printf("action: %v\n", action)
	// Implement your logic for handling the card_confirmed action.
	// ...
}

func handleNewGame(ctx context.Context, request events.APIGatewayWebsocketProxyRequest, dynamoDB *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	gameID := uuid.New().String()

	gameState := &GameState{
		GameID:        gameID,
		Player1:       "",              // Set this to the player's ID or a default value
		Player2:       "",              // Set this to the opponent's ID or a default value
		Player1Hand:   make([]Card, 4), // Initialize the player's hand with 4 cards
		Player2Hand:   make([]Card, 4), // Initialize the opponent's hand with 4 cards
		Player1HP:     100,             // Set the initial hit points for the player
		Player2HP:     100,             // Set the initial hit points for the opponent
		CurrentPlayer: "",              // Set this to the current player's ID or a default value
	}

	err := saveGameState(dynamoDB, gameState)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error saving game state",
		}, nil
	}

	response, err := json.Marshal(gameState)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error marshaling game state",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(response),
	}, nil
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		//Region: aws.String("us-west-2"), // Update this to your preferred AWS region
	})
	if err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err)
	}

	dynamoDB := dynamodb.New(sess)

	lambda.Start(func(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		return HandleGameAction(ctx, request, dynamoDB)
	})
}
