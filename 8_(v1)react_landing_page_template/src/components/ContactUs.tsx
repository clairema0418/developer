import React from 'react';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { Container, Typography, TextField, Button } from '@material-ui/core';

const useStyles = makeStyles(() => ({
  contactUs: {
    padding: '2rem 0',
  },
  title: {
    marginBottom: '1rem',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  textField: {
    marginBottom: '1rem',
  },
  submitButton: {
    alignSelf: 'flex-end',
  },
}));

const ContactUs: React.FC = () => {
  const classes = useStyles();

  return (
    <Container className={clsx(classes.contactUs)} id="contact">
      <Typography variant="h4" className={clsx(classes.title)}>
        聯繫我們
      </Typography>
      <form className={clsx(classes.form)}>
        <TextField
          label="姓名"
          variant="outlined"
          className={clsx(classes.textField)}
        />
        <TextField
          label="電子郵件"
          variant="outlined"
          className={clsx(classes.textField)}
        />
        <TextField
          label="留言"
          variant="outlined"
          multiline
          rows={4}
          className={clsx(classes.textField)}
        />
        <Button
          variant="contained"
          color="primary"
          className={clsx(classes.submitButton)}
        >
          提交
        </Button>
      </form>
    </Container>
  );
};

export default ContactUs;