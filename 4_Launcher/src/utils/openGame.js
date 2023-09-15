import { ipcRenderer } from 'electron';

export const openGame = (gameId, apiKey) => {
  ipcRenderer.send('open-game', { gameId, apiKey });
};