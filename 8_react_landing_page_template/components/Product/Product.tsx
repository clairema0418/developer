import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../store';
import ProductPhotos from '../ProductPhotos/ProductPhotos';
import ProductSpecifications from '../ProductSpecifications/ProductSpecifications';
import './Product.scss';

const Product: React.FC = () => {
  const productData = useSelector((state: RootState) => state.api.productData);

  return (
    <div className="product">
      {productData.map((product, index) => (
        <div key={index} className="product-item">
          <ProductPhotos photos={product.photos} />
          <ProductSpecifications specifications={product.specifications} />
        </div>
      ))}
    </div>
  );
};

export default Product;