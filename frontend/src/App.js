// App.js
import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './components/Home';
import Login from './components/Login';
import Signup from './components/Signup';
import WebSocket from './components/WebSocketExample';

function App() {
  // Create a WebSocket connection
    const ws = new WebSocket('ws://localhost:8080/ws');

    return (
      <Router>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/login">
            <Login ws={ws} />
          </Route>
          <Route path="/signup">
            <Signup ws={ws} />
          </Route>
        </Switch>
        <WebSocketExample ws={ws} />
      </Router>
    );
}

export default App;
