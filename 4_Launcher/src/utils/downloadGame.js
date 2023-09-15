import axios from 'axios';
import { getGameVersion } from '../services/api';

export async function downloadGame(apiKey, gameId) {
  try {
    const gameVersion = await getGameVersion(apiKey, gameId);
    const downloadUrl = `https://s3.amazonaws.com/your-bucket-name/${gameId}/${gameVersion}/game.zip`;

    const response = await axios.get(downloadUrl, {
      responseType: 'blob',
      headers: {
        'Content-Type': 'application/zip',
      },
    });

    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', `game_${gameId}_${gameVersion}.zip`);
    document.body.appendChild(link);
    link.click();
    link.remove();
  } catch (error) {
    console.error('Error downloading game:', error);
  }
}