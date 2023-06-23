import React, { useEffect, useRef } from 'react';

const WebSocketExample = () => {
  const socketRef = useRef(null);

  useEffect(() => {
    // Establish a WebSocket connection when the component mounts
    //establishes to the backend server
    socketRef.current = new WebSocket('ws://localhost:8080/ws');

    // Handle incoming messages from the server
    socketRef.current.onmessage = (event) => {
      const message = event.data;
      console.log('Received message:', message);
    };

    // Clean up the WebSocket connection when the component unmounts
    return () => {
      socketRef.current.close();
    };
  }, []);

  const sendMessage = () => {
    const message = 'Hello, server!';
    socketRef.current.send(message);
    console.log('Sent message:', message);
  };

  return (
    <div>
      <button onClick={sendMessage}>Send Message</button>
    </div>
  );
};

export default WebSocketExample;
