
import React, { useState } from 'react';
import axios from 'axios';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  //setWs allows you to update the value of the ws state variable.
  const [ws, setWs] = useState(null);

  const handleSignIn = async () => {
    try {
      const response = await axios.post('http://localhost:8080/backend/signin', {
        username: username,
        password: password
      });
      setMessage(response.data.message);
      if (ws) {
        ws.send('Login Successful');
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <h2>Login</h2>
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
      <button onClick={handleSignIn}>Sign In</button>
      <p>{message}</p>
    </div>
  );
};

export default Login;
