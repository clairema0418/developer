import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { fetchNewsUpdates } from '../../store/apiSlice';
import { RootState } from '../../store/index';
import { NewsUpdateData } from '../../config/testData.json';
import './News.scss';

interface NewsProps {}

const News: React.FC<NewsProps> = () => {
  const dispatch = useDispatch();
  const newsUpdates = useSelector((state: RootState) => state.api.newsUpdates);

  useEffect(() => {
    dispatch(fetchNewsUpdates());
  }, [dispatch]);

  return (
    <div className="news">
      <h2>News Updates</h2>
      <ul>
        {newsUpdates.map((newsUpdate: NewsUpdateData, index: number) => (
          <li key={index}>
            <h3>{newsUpdate.title}</h3>
            <p>{newsUpdate.description}</p>
            <p>{newsUpdate.date}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default News;