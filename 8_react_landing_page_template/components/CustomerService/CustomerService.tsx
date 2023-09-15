import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Typography, Grid, Paper } from '@material-ui/core';
import './CustomerService.scss';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
}));

const CustomerService: React.FC = () => {
  const classes = useStyles();

  return (
    <div className={classes.root} id="customerService">
      <Typography variant="h4" component="h2" gutterBottom>
        Customer Service
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6} md={4}>
          <Paper className={classes.paper}>
            <Typography variant="h6" component="h3">
              Phone
            </Typography>
            <Typography>+1 (123) 456-7890</Typography>
          </Paper>
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <Paper className={classes.paper}>
            <Typography variant="h6" component="h3">
              Email
            </Typography>
            <Typography>support@example.com</Typography>
          </Paper>
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <Paper className={classes.paper}>
            <Typography variant="h6" component="h3">
              Address
            </Typography>
            <Typography>123 Main St, Anytown, USA</Typography>
          </Paper>
        </Grid>
      </Grid>
    </div>
  );
};

export default CustomerService;