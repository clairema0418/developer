import { NextApiRequest, NextApiResponse } from 'next';
import fs from 'fs';
import path from 'path';
import { joinUsData } from '../../config/testData.json';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const { method } = req;

  switch (method) {
    case 'POST':
      try {
        const newJoinUsEntry = JSON.parse(req.body);
        joinUsData.push(newJoinUsEntry);
        fs.writeFileSync(
          path.join(process.cwd(), 'config', 'testData.json'),
          JSON.stringify({ ...joinUsData }, null, 2)
        );
        res.status(201).json({ message: 'Join us entry added successfully.' });
      } catch (error) {
        res.status(500).json({ message: 'Error adding join us entry.' });
      }
      break;
    default:
      res.setHeader('Allow', ['POST']);
      res.status(405).end(`Method ${method} Not Allowed`);
  }
}