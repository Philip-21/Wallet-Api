// App.js
import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './components/user/home';
import Login from './components/user/login';
import Signup from './components/user/signup';
import WebSocket from './components/websockets/websocket';


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
        <WebSocket ws={ws} />
      </Router>
    );
}

export default App;
