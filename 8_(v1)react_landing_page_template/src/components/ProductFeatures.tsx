import React from 'react';
import clsx from 'clsx';
import { featureData } from '../data/featureData';
import '../styles/product-features.scss';

interface FeatureProps {
  title: string;
  description: string;
  image: string;
}

const Feature: React.FC<FeatureProps> = ({ title, description, image }) => {
  return (
    <div className={clsx('feature')}>
      <img src={image} alt={title} className={clsx('feature-image')} />
      <h3 className={clsx('feature-title')}>{title}</h3>
      <p className={clsx('feature-description')}>{description}</p>
    </div>
  );
};

const ProductFeatures: React.FC = () => {
  return (
    <section id="features" className={clsx('product-features')}>
      <h2 className={clsx('section-title')}>產品特色和優勢介紹</h2>
      <div className={clsx('features-container')}>
        {featureData.map((feature, index) => (
          <Feature
            key={index}
            title={feature.title}
            description={feature.description}
            image={feature.image}
          />
        ))}
      </div>
    </section>
  );
};

export default ProductFeatures;