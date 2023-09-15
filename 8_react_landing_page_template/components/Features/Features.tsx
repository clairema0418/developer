import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../store';
import './Features.scss';

const Features: React.FC = () => {
  const features = useSelector((state: RootState) => state.apiSlice.productData.features);

  return (
    <div className="features" id="productFeatures">
      <h2>Product Features and Advantages</h2>
      <ul>
        {features.map((feature, index) => (
          <li key={index}>
            <h3>{feature.title}</h3>
            <p>{feature.description}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Features;