import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { useParams } from 'react-router-dom';
import { fetchGameIntroduction } from '../../store/gamesSlice';
import GameIntroduction from '../../components/GameIntroduction/GameIntroduction';
import './GameIntroductionPage.scss';

const GameIntroductionPage = () => {
  const dispatch = useDispatch();
  const { gameId } = useParams();
  const game = useSelector((state) => state.games.gameIntroduction);

  useEffect(() => {
    dispatch(fetchGameIntroduction(gameId));
  }, [dispatch, gameId]);

  return (
    <div className="game-introduction-page">
      {game && (
        <GameIntroduction
          imageUrl={game.imageUrl}
          description={game.description}
          onDownload={() => downloadGame(gameId)}
          onLaunch={() => launchGame(gameId)}
        />
      )}
    </div>
  );
};

export default GameIntroductionPage;