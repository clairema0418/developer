import React from 'react';
import clsx from 'clsx';
import { DownloadSchema } from '../schemas/DownloadSchema';
import '../styles/download-center.scss';

interface DownloadCenterProps {
  downloadData: DownloadSchema[];
}

const DownloadCenter: React.FC<DownloadCenterProps> = ({ downloadData }) => {
  return (
    <section id="download" className="download-center">
      <h2 className="download-center__title">下載中心和資源分享</h2>
      <div className="download-center__content">
        {downloadData.map((downloadItem, index) => (
          <div key={index} className={clsx('download-center__item', `download-center__item--${index}`)}>
            <h3 className="download-center__item-title">{downloadItem.title}</h3>
            <p className="download-center__item-description">{downloadItem.description}</p>
            <a href={downloadItem.link} className="download-center__item-link" target="_blank" rel="noopener noreferrer">
              下載 {downloadItem.title}
            </a>
          </div>
        ))}
      </div>
    </section>
  );
};

export default DownloadCenter;