import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Provider } from 'react-redux';

import store from './store/store';
import LoginPage from './pages/LoginPage/LoginPage';
import GameLobbyPage from './pages/GameLobbyPage/GameLobbyPage';
import GameIntroductionPage from './pages/GameIntroductionPage/GameIntroductionPage';
import './App.scss';

function App() {
  return (
    <Provider store={store}>
      <Router>
        <div className="App">
          <Switch>
            <Route exact path="/" component={LoginPage} />
            <Route path="/game-lobby" component={GameLobbyPage} />
            <Route path="/game-introduction/:id" component={GameIntroductionPage} />
          </Switch>
        </div>
      </Router>
    </Provider>
  );
}

export default App;