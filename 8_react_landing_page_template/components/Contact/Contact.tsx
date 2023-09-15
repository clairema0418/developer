import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { TextField, Button, Typography } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import { submitContactForm } from '../../pages/api/contact';
import clsx from 'clsx';
import './Contact.scss';

const useStyles = makeStyles({
  contactForm: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  textField: {
    marginBottom: '1rem',
  },
  submitButton: {
    marginTop: '1rem',
  },
});

const Contact: React.FC = () => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const result = await dispatch(submitContactForm({ name, email, message }));
    if (result) {
      setName('');
      setEmail('');
      setMessage('');
    }
  };

  return (
    <div id="contactUs" className="contact-container">
      <Typography variant="h4">Contact Us</Typography>
      <form onSubmit={handleSubmit} className={clsx(classes.contactForm, 'contact-form')}>
        <TextField
          label="Name"
          variant="outlined"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className={classes.textField}
          required
        />
        <TextField
          label="Email"
          variant="outlined"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className={classes.textField}
          required
        />
        <TextField
          label="Message"
          variant="outlined"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className={classes.textField}
          multiline
          rows={4}
          required
        />
        <Button type="submit" variant="contained" color="primary" className={classes.submitButton}>
          Submit
        </Button>
      </form>
    </div>
  );
};

export default Contact;