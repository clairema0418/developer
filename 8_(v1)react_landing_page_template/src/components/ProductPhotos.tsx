import React from 'react';
import { ProductSchema } from '../schemas/ProductSchema';
import '../styles/product-photos.scss';

interface ProductPhotosProps {
  productData: ProductSchema[];
}

const ProductPhotos: React.FC<ProductPhotosProps> = ({ productData }) => {
  return (
    <div className="product-photos" id="product">
      <h2>產品照片和規格</h2>
      <div className="product-photos-container">
        {productData.map((product, index) => (
          <div key={index} className="product-photo">
            <img src={product.image} alt={product.name} />
            <div className="product-specifications">
              <h3>{product.name}</h3>
              <ul>
                {product.specifications.map((spec, idx) => (
                  <li key={idx}>{spec}</li>
                ))}
              </ul>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProductPhotos;