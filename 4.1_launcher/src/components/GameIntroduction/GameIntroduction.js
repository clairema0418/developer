import React from 'react';
import { useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { DownloadButton } from '../DownloadButton/DownloadButton';
import { LaunchButton } from '../LaunchButton/LaunchButton';
import './GameIntroduction.scss';

export const GameIntroduction = () => {
  const { gameId } = useParams();
  const game = useSelector((state) => state.games.games.find((g) => g.id === gameId));

  if (!game) {
    return <div className="game-introduction">Game not found</div>;
  }

  return (
    <div className="game-introduction">
      <img src={game.imageUrl} alt={game.name} className="game-introduction__image" />
      <h2 className="game-introduction__title">{game.name}</h2>
      <p className="game-introduction__description">{game.description}</p>
      <div className="game-introduction__actions">
        <DownloadButton gameId={game.id} />
        <LaunchButton gameId={game.id} />
      </div>
    </div>
  );
};