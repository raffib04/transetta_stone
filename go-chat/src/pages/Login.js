import React, { useEffect, useState } from "react";
import LobbyCard from "../components/LobbyCard";

const Login = ({ onLogin }) => {
  const [username, setUsername] = useState("");
  const [language, setLanguage] = useState("en");
  const [room, setRoom] = useState("");
  const [activeRooms, setActiveRooms] = useState([]);

  const handleSubmit = () => {
    if (username && language && room) {
      onLogin(username, language, room);
    }
  };

  useEffect(() => {
    fetch(`${process.env.REACT_APP_BACKEND_URL}/rooms`)
      .then((res) => {
        if (!res.ok) {
          throw new Error("Network response was not ok " + res.statusText);
        }
        return res.json();
      })
      .then((data) => {
        if (Array.isArray(data)) {
          setActiveRooms(data);
        } else {
          setActiveRooms([]);
        }
      })
      .catch((error) => console.error("Error fetching rooms:", error));
  }, []);

  return (
    <div>
      <input
        type="text"
        placeholder="Enter username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        style={{
          background: "white",
          border: "1px solid gray",
          borderRadius: "5px",
          padding: "10px",
          margin: "10px",
          boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
          width: "200px",
        }}
      />
      <select
        value={language}
        onChange={(e) => setLanguage(e.target.value)}
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
        <option value="en">English</option>
        <option value="es">Spanish</option>
        <option value="pl">Polish</option>
        <option value="zh">Chinese</option>
        <option value="ko">Korean</option>
        <option value="ja">Japanese</option>
        <option value="fr">French</option>
      </select>

      <input
        type="text"
        placeholder="Enter room name"
        value={room}
        onChange={(e) => setRoom(e.target.value)}
        style={{
          background: "white",
          border: "1px solid gray",
          borderRadius: "5px",
          padding: "10px",
          margin: "10px",
          boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
          width: "200px",
        }}
      />
      <button
        onClick={handleSubmit}
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
        Enter Chat
      </button>
      <h2 style={{ margin: "10px" }}>Active Rooms</h2>
      {activeRooms.length > 0 ? (
        activeRooms.map((room) => (
          <LobbyCard
            key={room.roomName}
            name={room.roomName}
            numOfUsers={room.numClients}
          />
        ))
      ) : (
        <p>No active rooms available.</p>
      )}
    </div>
  );
};

export default Login;
