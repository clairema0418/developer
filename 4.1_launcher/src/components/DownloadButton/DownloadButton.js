import React from 'react';
import PropTypes from 'prop-types';
import { useDispatch } from 'react-redux';
import { downloadGame } from '../../utils/downloadGame';
import './DownloadButton.scss';

const DownloadButton = ({ gameId, userKey }) => {
  const dispatch = useDispatch();

  const handleDownload = async () => {
    try {
      await downloadGame(gameId, userKey);
      dispatch({ type: 'DOWNLOAD_GAME_SUCCESS' });
    } catch (error) {
      dispatch({ type: 'DOWNLOAD_GAME_FAILURE', payload: error.message });
    }
  };

  return (
    <button className="download-button" onClick={handleDownload}>
      Download
    </button>
  );
};

DownloadButton.propTypes = {
  gameId: PropTypes.string.isRequired,
  userKey: PropTypes.string.isRequired,
};

export default DownloadButton;