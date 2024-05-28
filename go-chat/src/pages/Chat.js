import React, { useState, useEffect, useRef } from "react";

const Chat = ({ username, language, room }) => {
    const [messages, setMessages] = useState([]);
    const [message, setMessage] = useState("");
    const ws = useRef(null);

    useEffect(() => {
        ws.current = new WebSocket(`ws://localhost:8080/ws?room=${room}`);

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
    }, [room, username, language]);

    const sendMessage = () => {
        if (message) {
            ws.current.send(JSON.stringify({ username, language, message }));
            setMessage("");
        }
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
                    <div key={index}>
                        {msg.username}: {msg.message}
                    </div>
                ))}
            </div>
            <input
                type='text'
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
            <button onClick={sendMessage} className="border border-gray-400 rounded-lg p-2 m-2">Send</button>
            <button
                onClick={() => {
                    fetch("http://localhost:8080/rooms")
                        // .then((response) => response.json())
                        .then((data) => {
                            console.log(data.body);
                        })
                        .catch((error) => {
                            console.error("Error fetching rooms:", error);
                        });
                }}
                className="border border-gray-400 rounded-lg p-2 m-2"
            >
                Show Rooms
            </button>
        </div>
    );
};

export default Chat;
