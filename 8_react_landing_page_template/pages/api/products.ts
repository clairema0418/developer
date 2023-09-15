import { NextApiRequest, NextApiResponse } from 'next';
import fs from 'fs';
import path from 'path';
import { productData } from '../../config/testData.json';

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method === 'GET') {
    res.status(200).json(productData);
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
};

export default handler;