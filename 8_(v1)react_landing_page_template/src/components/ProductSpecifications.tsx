import React from 'react';
import { ProductSchema } from '../dataSchemas/ProductSchema';
import '../styles/product-specifications.scss';

interface ProductSpecificationsProps {
  productData: ProductSchema[];
}

const ProductSpecifications: React.FC<ProductSpecificationsProps> = ({ productData }) => {
  return (
    <div className="product-specifications" id="product">
      <h2>產品照片和規格</h2>
      <div className="product-specifications-container">
        {productData.map((product, index) => (
          <div key={index} className="product-specification">
            <img src={product.image} alt={product.name} />
            <h3>{product.name}</h3>
            <ul>
              {product.specifications.map((spec, idx) => (
                <li key={idx}>{spec}</li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProductSpecifications;