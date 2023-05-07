import React, { useState } from "react";
import "./App.css";
import GameBoard from "./components/GameBoard";

function App() {
  const [gameName] = useState("My Card Game");
  const [turnNumber, setTurnNumber] = useState(1);

  const [player] = useState({
    name: "Player 1",
    hitPoints: 30,
    cards: [
      { id: 1, attack: 5, defense: 3, power: 2, speed: 1 },
      { id: 2, attack: 3, defense: 5, power: 1, speed: 2 },
      { id: 3, attack: 4, defense: 4, power: 2, speed: 2 },
      { id: 4, attack: 6, defense: 2, power: 3, speed: 1 },
    ],
  });

  const [opponent] = useState({
    name: "Opponent 1",
    hitPoints: 30,
    cards: [
      { id: 5, attack: 4, defense: 4, power: 2, speed: 2 },
      { id: 6, attack: 5, defense: 3, power: 1, speed: 1 },
      { id: 7, attack: 6, defense: 2, power: 2, speed: 2 },
      { id: 8, attack: 3, defense: 5, power: 3, speed: 1 },
    ],
  });

  const [notifications] = useState([
    "Player attacked with card1, opponent defended with card2.",
    "Player wins and opponent received 5 damage.",
  ]);

  return (
    <div className="App">
      <GameBoard
        gameName={gameName}
        turnNumber={turnNumber}
        player={player}
        opponent={opponent}
        notifications={notifications}
      />
    </div>
  );
}

export default App;