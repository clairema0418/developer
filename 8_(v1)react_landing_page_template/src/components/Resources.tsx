import React from 'react';
import clsx from 'clsx';
import { makeStyles } from '@mui/styles';

const useStyles = makeStyles({
  resourcesContainer: {
    // Add your scss styles here
  },
});

interface Resource {
  title: string;
  description: string;
  link: string;
}

interface ResourcesProps {
  resourcesData: Resource[];
}

const Resources: React.FC<ResourcesProps> = ({ resourcesData }) => {
  const classes = useStyles();

  return (
    <div id="resources" className={clsx(classes.resourcesContainer)}>
      <h2>資源分享</h2>
      <ul>
        {resourcesData.map((resource, index) => (
          <li key={index}>
            <h3>{resource.title}</h3>
            <p>{resource.description}</p>
            <a href={resource.link} target="_blank" rel="noopener noreferrer">
              下載
            </a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Resources;