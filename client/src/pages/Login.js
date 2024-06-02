import React, { useEffect, useState } from "react";
import LobbyCard from "../components/LobbyCard";
import { fetchActiveRooms } from "../utils/ActiveRooms"; // Adjust the path as necessary

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
      const loadRooms = async () => {
          const rooms = await fetchActiveRooms();
          setActiveRooms(rooms);
      };
      loadRooms();
  }, []);

  const handleRoomKeyDown = (e) => {
    if (e.key === "Enter") {
      handleSubmit();
    }
  };

  const handleRoomSelect = (selectedRoom) => {
    setRoom(selectedRoom);
    // Automatically submit the form when a room is selected
    handleSubmit();
  };

  return (
    <div>
      <h1 className="p-5 font-bold text-5xl">Transetta Stone</h1>
      <input
        type="text"
        placeholder="Enter username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        onKeyDown={handleRoomKeyDown}
        style={inputStyle}
      />
      <select
        value={language}
        onChange={(e) => setLanguage(e.target.value)}
        style={inputStyle}
      >
        <option value="en">English</option>
        <option value="zh">Chinese</option>
        <option value="fr">French</option>
        <option value="ja">Japanese</option>
        <option value="ko">Korean</option>
        <option value="pl">Polish</option>
        <option value="es">Spanish</option>
      </select>

      <input
        type="text"
        placeholder="Enter room name"
        value={room}
        onChange={(e) => setRoom(e.target.value)}
        onKeyDown={handleRoomKeyDown}
        style={inputStyle}
      />
      <button
        onClick={handleSubmit}
        style={buttonStyle}
      >
        Enter Chat
      </button>
      <h2 className="p-5 font-bold">Active Rooms</h2>
      {activeRooms.length > 0 ? (
        activeRooms.map((room) => (
          <LobbyCard
            key={room.roomName}
            name={room.roomName}
            numOfUsers={room.numClients}
            onClick={handleRoomSelect}
            currentRoom={false}
          />
        ))
      ) : (
        <p className="p-4">No active rooms available.</p>
      )}
    </div>
  );
};

const inputStyle = {
  background: "white",
  border: "1px solid gray",
  borderRadius: "5px",
  padding: "10px",
  margin: "10px",
  boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
  width: "200px",
};

const buttonStyle = {
  ...inputStyle,
  cursor: "pointer",
};

export default Login;
