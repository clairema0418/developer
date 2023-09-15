import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../store';
import { CustomerReviewData } from '../../store/apiSlice';
import './Reviews.scss';

const Reviews: React.FC = () => {
  const customerReviews = useSelector((state: RootState) => state.api.customerReviewData);

  return (
    <div className="reviews" id="customerReviews">
      <h2>Customer Reviews</h2>
      <div className="reviews-container">
        {customerReviews.map((review: CustomerReviewData, index: number) => (
          <div key={index} className="review">
            <div className="review-header">
              <img src={review.avatar} alt={`${review.name}'s avatar`} />
              <h3>{review.name}</h3>
            </div>
            <p>{review.review}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Reviews;