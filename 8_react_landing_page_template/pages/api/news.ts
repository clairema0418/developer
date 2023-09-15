import { NextApiRequest, NextApiResponse } from 'next';
import testData from '../../config/testData.json';

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
  try {
    const newsData = testData.newsUpdates;
    res.status(200).json(newsData);
  } catch (error) {
    res.status(500).json({ message: 'Error fetching news data' });
  }
};

export default handler;