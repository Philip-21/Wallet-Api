import React, { useState } from "react";
import { Link } from 'react-router-dom';
import WebSocket from '../websockets/websocket';

function Home() {
  const [receivedMessage, setReceivedMessage] = useState('');

  const handleWebSocketMessage = (message) => {
    // Handle the received message as needed
    setReceivedMessage(message);
  };

  return (
    <div className="container">
      <div className="row">
        <div className="col-md-8 offset-md-2">
          <h2 className="mt-5">Wallet Payment App</h2>
          <hr />
          <Link className="btn btn-outline-secondary" to="/login">Login</Link>
          <Link className="btn btn-outline-secondary" to="/register">Register</Link>
          <p>Received message: {receivedMessage}</p>
        </div>
      </div>
      <WebSocket onMessageReceived={handleWebSocketMessage} />
    </div>
  );
}

export default Home;
