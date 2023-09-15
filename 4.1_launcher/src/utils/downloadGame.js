import { fetchGameVersion } from './fetchGameVersion';
import awsS3 from '../services/awsS3';

async function downloadGame(gameId, userKey) {
  try {
    const gameVersion = await fetchGameVersion(gameId, userKey);
    const downloadUrl = awsS3.getSignedUrl('getObject', {
      Bucket: 'your-bucket-name',
      Key: `games/${gameId}/${gameVersion}.zip`,
      Expires: 60 * 5, // 5 minutes
    });

    const response = await fetch(downloadUrl);
    const blob = await response.blob();
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', `${gameId}-${gameVersion}.zip`);
    document.body.appendChild(link);
    link.click();
    link.remove();
  } catch (error) {
    console.error('Error downloading game:', error);
  }
}

export default downloadGame;