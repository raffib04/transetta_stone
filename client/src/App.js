import React, { useState } from "react";
import Login from "./pages/Login";
import Chat from "./pages/Chat";

const App = () => {
  const [user, setUser] = useState(null);

  const handleLogin = (username, language, room) => {
    setUser({ username, language, room });
  };

  return (
    <div>
      {!user ? (
        <Login onLogin={handleLogin} />
      ) : (
        <Chat
          username={user.username}
          language={user.language}
          room={user.room}
        />
      )}
    </div>
  );
};

export default App;
