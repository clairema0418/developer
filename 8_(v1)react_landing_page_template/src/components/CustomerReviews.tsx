import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import { Card, CardContent, Typography } from '@material-ui/core';

const useStyles = makeStyles(() => ({
  customerReviewsContainer: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  reviewCard: {
    maxWidth: 345,
    margin: '1rem',
  },
}));

interface Testimonial {
  name: string;
  title: string;
  review: string;
}

const testimonialData: Testimonial[] = [
  // Add your testimonial data here
];

const CustomerReviews: React.FC = () => {
  const classes = useStyles();

  return (
    <div className={classes.customerReviewsContainer}>
      <Typography variant="h4">客戶評價</Typography>
      {testimonialData.map((testimonial, index) => (
        <Card key={index} className={clsx(classes.reviewCard)}>
          <CardContent>
            <Typography variant="h6">{testimonial.name}</Typography>
            <Typography variant="subtitle1">{testimonial.title}</Typography>
            <Typography variant="body1">{testimonial.review}</Typography>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};

export default CustomerReviews;