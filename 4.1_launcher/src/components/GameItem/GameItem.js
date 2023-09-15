import React from 'react';
import PropTypes from 'prop-types';
import './GameItem.scss';

const GameItem = ({ game, onClick }) => {
  return (
    <div className="game-item" onClick={() => onClick(game.id)}>
      <img src={game.imageUrl} alt={game.name} className="game-item__image" />
      <div className="game-item__info">
        <h3 className="game-item__name">{game.name}</h3>
      </div>
    </div>
  );
};

GameItem.propTypes = {
  game: PropTypes.shape({
    id: PropTypes.string.isRequired,
    name: PropTypes.string.isRequired,
    imageUrl: PropTypes.string.isRequired,
  }).isRequired,
  onClick: PropTypes.func.isRequired,
};

export default GameItem;