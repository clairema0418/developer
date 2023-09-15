import AWS from 'aws-sdk';

const s3 = new AWS.S3({
  accessKeyId: process.env.AWS_ACCESS_KEY_ID,
  secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY,
  region: process.env.AWS_REGION,
});

export const downloadGameFromS3 = async (gameKey, gameVersion, destinationPath) => {
  const params = {
    Bucket: process.env.AWS_S3_BUCKET,
    Key: `${gameKey}/${gameVersion}`,
  };

  try {
    const data = await s3.getObject(params).promise();
    const fileStream = require('fs').createWriteStream(destinationPath);
    fileStream.write(data.Body);
    fileStream.end();
    return true;
  } catch (error) {
    console.error('Error downloading game from S3:', error);
    return false;
  }
};