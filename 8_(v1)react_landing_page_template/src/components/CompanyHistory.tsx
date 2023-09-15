import React from 'react';
import { CompanySchema } from '../shared/types';
import '../styles/company-history.scss';

interface CompanyHistoryProps {
  companyData: CompanySchema;
}

const CompanyHistory: React.FC<CompanyHistoryProps> = ({ companyData }) => {
  return (
    <section id="company-history" className="company-history">
      <h2 className="company-history__title">公司歷史沿革</h2>
      <div className="company-history__content">
        {companyData.history.map((item, index) => (
          <div key={index} className="company-history__item">
            <h3 className="company-history__item-title">{item.title}</h3>
            <p className="company-history__item-description">{item.description}</p>
          </div>
        ))}
      </div>
    </section>
  );
};

export default CompanyHistory;