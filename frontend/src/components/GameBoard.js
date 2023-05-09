import React, { useEffect, useState } from "react";
import "./GameBoard.css";
import Card from "./Card";
//import WebSocketInstance from "../websocket";

function GameBoard({ gameName, turnNumber, player, opponent, notifications }) {
    useEffect(() => {
        // Replace "ws://your-websocket-url" with your actual WebSocket URL
        WebSocketInstance.connect("ws://your-websocket-url");

        // Add callback for receiving messages from the backend
        WebSocketInstance.addCallback("action_confirmed", (data) => {
            // Handle the confirmation from the backend, e.g., setSelectedCard(data.card)
        });
    }, []);



    const [zoomedCard, setZoomedCard] = useState(null);
    const [selectedCard, setSelectedCard] = useState(null);

    const handleCardClick = (card, type) => {
        if (type === "player") {
            WebSocketInstance.sendCommand("card_clicked", { card });
            setZoomedCard(card);
        }
    };

    const closeZoom = () => {
        setZoomedCard(null);
    };

    const handleConfirm = () => {
        WebSocketInstance.sendCommand("card_confirmed", { card: zoomedCard });
        setSelectedCard(zoomedCard);
        closeZoom();
    };

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
                <div className="cards">
                    {opponent.cards.map((card) => (
                        <Card key={card.id} card={card} type="opponent" />
                    ))}
                </div>
            </section>

            <section className="battle">
                {selectedCard && (
                    <div className="attacker-card">
                        <Card card={selectedCard} type="player" />
                    </div>
                )}
                <div className="defender">
                    {/* Render defender card content */}
                </div>
            </section>

            <section className="player-hand">
                <h2>{player.name}</h2>
                <p>Hit Points: {player.hitPoints}</p>
                <div className="cards">
                    {player.cards.map((card) => (
                        <Card key={card.id} card={card} type="player" onCardClick={handleCardClick} />
                    ))}
                </div>
                <button>End Turn</button>
            </section>

            <section className="notifications">
                <ul>
                    {notifications.map((notification, index) => (
                        <li key={index}>{notification}</li>
                    ))}
                </ul>
            </section>
            {zoomedCard && (
                <>
                    <div className="overlay" onClick={closeZoom}></div>
                    <div className="zoomed-card-container">
                        <Card card={zoomedCard} type="player zoomed" />
                        <button className="confirm-button" onClick={handleConfirm}>
                            Confirm
                        </button>
                    </div>
                </>
            )}

        </div>
    );
}

export default GameBoard;