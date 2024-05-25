import React, { useState } from "react";

const Login = ({ onLogin }) => {
    const [username, setUsername] = useState("");
    const [language, setLanguage] = useState("en");
    const [room, setRoom] = useState("");

    const handleSubmit = () => {
        if (username && language && room) {
            onLogin(username, language, room);
        }
    };

    return (
        <div>
            <input
                type='text'
                placeholder='Enter username'
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <select
                value={language}
                onChange={(e) => setLanguage(e.target.value)}
            >
                <option value='en'>English</option>
                <option value='es'>Spanish</option>
                <option value='pl'>Polish</option>
                <option value='zh'>Chinese</option>
                <option value='ko'>Korean</option>
                <option value='ja'>Japanese</option>
                {/* Add more languages as needed */}
            </select>
            <input
                type='text'
                placeholder='Enter room name'
                value={room}
                onChange={(e) => setRoom(e.target.value)}
            />
            <button onClick={handleSubmit}>Enter Chat</button>
        </div>
    );
};

export default Login;
