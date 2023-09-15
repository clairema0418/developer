import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { submitJoinUsForm } from '../../pages/api/apiSlice';
import { TextField, Button } from '@material-ui/core';
import clsx from 'clsx';
import { css } from '@emotion/react';
import './Join.scss';

const Join: React.FC = () => {
  const dispatch = useDispatch();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    dispatch(submitJoinUsForm({ name, email, message }));
    setName('');
    setEmail('');
    setMessage('');
  };

  return (
    <div className="join">
      <h2>Join Us</h2>
      <form onSubmit={handleSubmit}>
        <TextField
          label="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className={clsx('join__input', css`margin-bottom: 1rem;`)}
          required
        />
        <TextField
          label="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className={clsx('join__input', css`margin-bottom: 1rem;`)}
          required
        />
        <TextField
          label="Message"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className={clsx('join__input', css`margin-bottom: 1rem;`)}
          multiline
          rows={4}
          required
        />
        <Button type="submit" variant="contained" color="primary">
          Submit
        </Button>
      </form>
    </div>
  );
};

export default Join;