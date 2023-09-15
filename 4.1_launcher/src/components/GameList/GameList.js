import React from 'react';
import { useSelector } from 'react-redux';
import { selectGames } from '../../store/gamesSlice';
import GameItem from '../GameItem/GameItem';
import './GameList.scss';

const GameList = () => {
  const games = useSelector(selectGames);

  return (
    <div className="game-list">
      {games.map((game) => (
        <GameItem key={game.id} game={game} />
      ))}
    </div>
  );
};

export default GameList;