import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Typography, Grid, Card, CardContent } from '@material-ui/core';

const useStyles = makeStyles(() => ({
  customerServiceContainer: {
    padding: '2rem 0',
  },
  customerServiceCard: {
    boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
    borderRadius: '8px',
    padding: '1rem',
  },
  title: {
    marginBottom: '1rem',
  },
}));

interface CustomerServiceProps {
  contactData: {
    title: string;
    description: string;
    phone: string;
    email: string;
  };
}

const CustomerService: React.FC<CustomerServiceProps> = ({ contactData }) => {
  const classes = useStyles();

  return (
    <Grid container className={classes.customerServiceContainer} id="contact">
      <Grid item xs={12} sm={6}>
        <Card className={classes.customerServiceCard}>
          <CardContent>
            <Typography variant="h5" className={classes.title}>
              {contactData.title}
            </Typography>
            <Typography variant="body1">{contactData.description}</Typography>
            <Typography variant="body1">{contactData.phone}</Typography>
            <Typography variant="body1">{contactData.email}</Typography>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default CustomerService;