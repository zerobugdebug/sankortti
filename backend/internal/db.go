package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/zerobugdebug/sankortti/backend/internal/game"
)

const (
	gameTable = "CardGame"
)

type DBHandler struct {
	client *dynamodb.DynamoDB
}

func NewDBHandler(client *dynamodb.DynamoDB) *DBHandler {
	return &DBHandler{client: client}
}

func (db *DBHandler) GetGame(gameID string) (*game.Game, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(gameTable),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(gameID),
			},
		},
	}

	result, err := db.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	var g game.Game
	err = dynamodbattribute.UnmarshalMap(result.Item, &g)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (db *DBHandler) SaveGame(g *game.Game) error {
	av, err := dynamodbattribute.MarshalMap(g)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(gameTable),
		Item:      av,
	}

	_, err = db.client.Put(input)
	return err
}
