import React from 'react';
import PropTypes from 'prop-types';
import './LaunchButton.scss';
import { launchGame } from '../../utils/launchGame';

const LaunchButton = ({ gameId }) => {
  const handleLaunch = () => {
    launchGame(gameId);
  };

  return (
    <button className="launch-button" onClick={handleLaunch}>
      Launch Game
    </button>
  );
};

LaunchButton.propTypes = {
  gameId: PropTypes.string.isRequired,
};

export default LaunchButton;