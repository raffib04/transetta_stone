import React from "react";

const LobbyCard = ({ name, numOfUsers, onClick, currentRoom }) => {
    return (
        <button
            className='card'
            style={{
                background: "white",
                border: "1px solid gray",
                borderRadius: "5px",
                padding: "10px",
                margin: "10px",
                boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
                width: "200px",
                textAlign: "left",
                fontWeight: currentRoom ? "bold" : "normal",
            }}
            onClick={() => onClick(name)}
        >
            <h2>{name}</h2>
            <p>Number of Users: {numOfUsers}</p>
        </button>
    );
};

export default LobbyCard;
