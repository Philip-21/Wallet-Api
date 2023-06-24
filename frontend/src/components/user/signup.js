import React, { useState } from 'react';
import axios from 'axios';

const Signup = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [ws, setWs] = useState(null);

  const handleSignUp = async () => {
    try {
      //The backend URL is used in the axios requests made from the frontend components to communicate with the backend server.
      const response = await axios.post('http://localhost:8080/backend/signup', {
        username: username,
        password: password
      });
      setMessage(response.data.message);
      if (ws) {
        ws.send('Signup Successful');
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <h2>Signup</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={e => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={e => setPassword(e.target.value)}
      />
      <button onClick={handleSignUp}>Sign Up</button>
      <p>{message}</p>
    </div>
  );
};

export default Signup;
