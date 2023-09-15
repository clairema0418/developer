import { NextApiRequest, NextApiResponse } from 'next';

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method === 'POST') {
    const { name, email, message } = req.body;

    if (!name || !email || !message) {
      res.status(400).json({ error: 'All fields are required.' });
      return;
    }

    // Here, you can integrate your email service to send the contact form data
    // For example, using SendGrid, Nodemailer, or any other email service

    res.status(200).json({ message: 'Contact form submitted successfully.' });
  } else {
    res.status(405).json({ error: 'Method not allowed.' });
  }
};

export default handler;