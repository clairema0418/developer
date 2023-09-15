import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(() => ({
  eventsContainer: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  event: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginBottom: '1rem',
  },
  eventTitle: {
    fontWeight: 'bold',
  },
  eventDate: {
    fontStyle: 'italic',
  },
  eventDescription: {
    textAlign: 'center',
  },
}));

interface Event {
  title: string;
  date: string;
  description: string;
}

interface EventsProps {
  eventsData: Event[];
}

const Events: React.FC<EventsProps> = ({ eventsData }) => {
  const classes = useStyles();

  return (
    <div className={classes.eventsContainer}>
      {eventsData.map((event, index) => (
        <div key={index} className={classes.event}>
          <h3 className={classes.eventTitle}>{event.title}</h3>
          <p className={classes.eventDate}>{event.date}</p>
          <p className={classes.eventDescription}>{event.description}</p>
        </div>
      ))}
    </div>
  );
};

export default Events;