import React from 'react';
import clsx from 'clsx';
import { Button, Container, Grid, Typography } from '@mui/material';
import { makeStyles } from '@mui/styles';

const useStyles = makeStyles({
  joinUsSection: {
    padding: '4rem 0',
  },
  joinUsTitle: {
    marginBottom: '2rem',
  },
  joinUsButton: {
    marginTop: '1rem',
  },
});

const JoinUs: React.FC = () => {
  const classes = useStyles();

  return (
    <section id="join-us" className={clsx(classes.joinUsSection)}>
      <Container>
        <Typography variant="h4" align="center" className={clsx(classes.joinUsTitle)}>
          加入我們
        </Typography>
        <Grid container justifyContent="center">
          <Grid item xs={12} sm={8} md={6}>
            <Typography variant="body1" align="center">
              我們正在尋找有才華的人才加入我們的團隊。如果您對我們的產品感興趣，請點擊下面的按鈕提交您的簡歷。
            </Typography>
            <Grid container justifyContent="center">
              <Button
                variant="contained"
                color="primary"
                className={clsx(classes.joinUsButton)}
                href="/contact"
              >
                提交簡歷
              </Button>
            </Grid>
          </Grid>
        </Grid>
      </Container>
    </section>
  );
};

export default JoinUs;