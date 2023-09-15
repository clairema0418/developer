import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { fetchGames } from '../../store/gamesSlice';
import GameList from '../../components/GameList/GameList';
import './GameLobbyPage.scss';

const GameLobbyPage = () => {
  const dispatch = useDispatch();
  const history = useHistory();
  const games = useSelector((state) => state.games.games);
  const userKey = useSelector((state) => state.auth.key);

  useEffect(() => {
    if (!userKey) {
      history.push('/');
    } else {
      dispatch(fetchGames(userKey));
    }
  }, [dispatch, userKey, history]);

  const handleGameClick = (gameId) => {
    history.push(`/game-introduction/${gameId}`);
  };

  return (
    <div className="game-lobby-page">
      <h1>Game Lobby</h1>
      <GameList games={games} onGameClick={handleGameClick} />
    </div>
  );
};

export default GameLobbyPage;