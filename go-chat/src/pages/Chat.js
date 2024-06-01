// Chat.js
import React, { useState, useEffect, useRef } from "react";
import { fetchActiveRooms } from "../utils/ActiveRooms";
import LobbyCard from "../components/LobbyCard";

const Chat = ({ username, language, room }) => {
    const [messages, setMessages] = useState([]);
    const [message, setMessage] = useState("");
    const [activeRooms, setActiveRooms] = useState([]);
    const [showActiveRooms, setShowActiveRooms] = useState(false);
    const [selectedRoom, setSelectedRoom] = useState("");
    const ws = useRef(null);
    const messagesEndRef = useRef(null);

    useEffect(() => {
        ws.current = new WebSocket(
        `${process.env.REACT_APP_BACKEND_URL_WEBSOCKET}/ws?room=${selectedRoom || room}`
        );

        ws.current.onopen = () => {
        ws.current.send(JSON.stringify({ username, language }));
        };

        ws.current.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setMessages((prevMessages) => [...prevMessages, data]);
        };

        return () => {
        ws.current.close();
        };
    }, [selectedRoom, room, username, language]);

    useEffect(() => {
        if (messagesEndRef.current) {
        messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
        }
    }, [messages]);

    const sendMessage = () => {
        if (message) {
        ws.current.send(JSON.stringify({ username, language, message }));
        setMessage("");
        }
    };

    const getRooms = async () => {
        if (!showActiveRooms) {
            const rooms = await fetchActiveRooms();
            setActiveRooms(rooms);
                
        }
        setShowActiveRooms(!showActiveRooms);
    };

    const handleRoomSelect = (selectedRoomName) => {
        setSelectedRoom(selectedRoomName);
        setShowActiveRooms(false);
    };

    return (
        <div>
            <div
                style={{
                height: "300px",
                width: "500px",
                border: "1px solid black",
                overflow: "auto",
                }}
            >
                {messages.map((msg, index) => (
                <div
                    key={index}
                    className={`ml-4 my-2 p-2 rounded-lg ${
                    msg.username === username ? "bg-blue-100 self-end" : "bg-gray-100"
                    }`}
                    style={{
                    textAlign: msg.username === username ? "right" : "left",
                    alignSelf: msg.username === username ? "flex-end" : "flex-start",
                    marginLeft: msg.username === username ? "auto" : "0",
                    }}
                >
                    <strong>{msg.username}</strong>: {msg.message}
                </div>
                ))}
                <div ref={messagesEndRef} />
            </div>
            <input
                type="text"
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={(e) => {
                if (e.key === "Enter") {
                    sendMessage();
                }
                }}
                style={{ width: "400px" }}
                className="border border-gray-400 rounded-lg p-2 m-2"
            />
            <button
                onClick={sendMessage}
                className="border border-gray-400 rounded-lg p-2 m-2"
            >
                Send
            </button>
            <button
                onClick={getRooms}
                className="border border-gray-400 rounded-lg p-2 m-2"
            >
                {showActiveRooms ? "Hide Rooms" : "Show Rooms"}
            </button>
            {showActiveRooms && (
                <div>
                    <h2 style={{ margin: "10px" }}>Active Rooms</h2>
                    {activeRooms.length > 0 ? (
                        activeRooms.map((r) => (
                            <LobbyCard
                                key={r.roomName}
                                name={r.roomName}
                                numOfUsers={r.numClients}
                                onClick={handleRoomSelect}
                                currentRoom={selectedRoom ? r.roomName === selectedRoom : r.roomName === room} // Compare room name with current room
                            />
                        ))
                    ) : (
                        <p>No active rooms available.</p>
                    )}
                </div>
            )}
        </div>
    );
};

export default Chat;
