import React from "react";

const LobbyCard = ({ name, numOfUsers }) => {
    return (
        <div
            className='card'
            style={{
                background: "white",
                border: "1px solid gray",
                borderRadius: "5px",
                padding: "10px",
                margin: "10px",
                boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
                width: "200px",
            }}
        >
            <h2>{name}</h2>
            <p>Number of Users: {numOfUsers}</p>
        </div>
    );
};

export default LobbyCard;
