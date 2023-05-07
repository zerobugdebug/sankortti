/* import React from "react";
import './Card.css';

const Card = ({ card, onClick, isSelected }) => {
    const cardClass = isSelected ? "card selected" : "card";

    return (
        <div className={cardClass} onClick={onClick}>
            <div className="attack">Attack: {card.attack}</div>
            <div className="defense">Defense: {card.defense}</div>
            <div className="power">Power: {card.power}</div>
            <div className="speed">Speed: {card.speed}</div>
        </div>
    );
};

export default Card; */

import React from "react";
import "./Card.css";

function Card({ card, type }) {
    const { attack, defense, power, speed } = card;

    return (
        <div className={`card ${type}`}>
            <div className="card-stats">
                <div className="attack">Attack: {attack}</div>
                <div className="defense">Defense: {defense}</div>
                <div className="power">Power: {power}</div>
                <div className="speed">Speed: {speed}</div>
            </div>
        </div>
    );
}

export default Card;