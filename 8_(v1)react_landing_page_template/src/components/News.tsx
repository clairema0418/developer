import React from 'react';
import { NewsSchema } from '../schemas/NewsSchema';
import '../styles/news.scss';

interface NewsProps {
  newsData: NewsSchema[];
}

const News: React.FC<NewsProps> = ({ newsData }) => {
  return (
    <section id="news" className="news-section">
      <h2 className="news-title">新聞動態和公司活動</h2>
      <div className="news-container">
        {newsData.map((newsItem, index) => (
          <div key={index} className="news-item">
            <img src={newsItem.image} alt={newsItem.title} className="news-image" />
            <div className="news-content">
              <h3 className="news-item-title">{newsItem.title}</h3>
              <p className="news-item-date">{newsItem.date}</p>
              <p className="news-item-description">{newsItem.description}</p>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default News;