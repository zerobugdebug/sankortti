package game

import (
	"errors"
	"math/rand"
)

// Game struct represents the game state and associated metadata.
type Game struct {
	ID           string
	Player1      Player
	Player2      Player
	CurrentTurn  string
	ActivePlayer *Player
	GameOver     bool
}

// Player struct represents a player in the game, with cards, hit points, and other attributes.
type Player struct {
	ID        string
	HP        int
	Cards     []Card
	Selected  *Card
	Confirmed bool
}

// Card struct represents a card with attributes such as attack, defense, power, and speed.
type Card struct {
	ID       string
	Attack   int
	Defense  int
	Power    int
	Speed    int
	Selected bool
}

const (
	playerHP = 100
)

var (
	ErrInvalidAction = errors.New("invalid action")
)

// NewGame creates a new Game instance with initial state.
func NewGame(playerID string, computer bool) *Game {
	gameID := generateGameID()
	player1 := Player{
		ID:    playerID,
		HP:    playerHP,
		Cards: generateCards(),
	}
	player2ID := "computer"
	if !computer {
		player2ID = "" // Assign a different ID when a second player joins.
	}
	player2 := Player{
		ID:    player2ID,
		HP:    playerHP,
		Cards: generateCards(),
	}
	game := &Game{
		ID:           gameID,
		Player1:      player1,
		Player2:      player2,
		CurrentTurn:  player1.ID,
		ActivePlayer: &player1,
		GameOver:     false,
	}
	return game
}

func generateGameID() string {
	// Generate a unique game ID (consider using a UUID package for better uniqueness)
	return "gameID"
}

func generateCards() []Card {
	// Generate 4 cards for a player with randomized attributes
	cards := make([]Card, 4)
	for i := range cards {
		cards[i] = Card{
			ID:      "cardID",
			Attack:  rand.Intn(10) + 1,
			Defense: rand.Intn(10) + 1,
			Power:   rand.Intn(10) + 1,
			Speed:   rand.Intn(10) + 1,
		}
	}
	return cards
}
