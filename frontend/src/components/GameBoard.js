import React from "react";
import "./GameBoard.css";
import Card from "./Card";

function GameBoard({ gameName, turnNumber, player, opponent, notifications }) {
    return (
        <div className="gameboard">
            <header>
                <h1>{gameName}</h1>
                <p>Turn: {turnNumber}</p>
                <button>Exit Game</button>
            </header>

            <section className="opponent-hand">
                <h2>{opponent.name}</h2>
                <p>Hit Points: {opponent.hitPoints}</p>
                {opponent.cards.map((card) => (
                    <Card key={card.id} card={card} type="opponent" />
                ))}
            </section>

            <section className="battle">
                <div className="attacker">
                    {/* Render attacker card content */}
                </div>
                <div className="defender">
                    {/* Render defender card content */}
                </div>
            </section>

            <section className="player-hand">
                <h2>{player.name}</h2>
                <p>Hit Points: {player.hitPoints}</p>
                {player.cards.map((card) => (
                    <Card key={card.id} card={card} type="player" />
                ))}
                <button>End Turn</button>
            </section>

            <section className="notifications">
                <ul>
                    {notifications.map((notification, index) => (
                        <li key={index}>{notification}</li>
                    ))}
                </ul>
            </section>
        </div>
    );
}

export default GameBoard;