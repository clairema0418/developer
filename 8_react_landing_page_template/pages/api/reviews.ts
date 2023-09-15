import { NextApiRequest, NextApiResponse } from 'next';
import fs from 'fs';
import path from 'path';

const reviewsHandler = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method === 'GET') {
    try {
      const dataPath = path.join(process.cwd(), 'config', 'testData.json');
      const fileContent = fs.readFileSync(dataPath, 'utf8');
      const testData = JSON.parse(fileContent);
      const reviews = testData.customerReviewData;

      res.status(200).json(reviews);
    } catch (error) {
      res.status(500).json({ message: 'Error fetching reviews data' });
    }
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
};

export default reviewsHandler;