import React from 'react';
import './Slogan.scss';

interface SloganProps {
  slogan: string;
}

const Slogan: React.FC<SloganProps> = ({ slogan }) => {
  return (
    <div className="slogan-container">
      <h2 className="slogan-text">{slogan}</h2>
    </div>
  );
};

export default Slogan;