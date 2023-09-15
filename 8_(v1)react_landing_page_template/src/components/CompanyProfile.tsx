import React from 'react';
import clsx from 'clsx';
import { CompanySchema } from '../lib/schemas';
import '../styles/company-profile.scss';

interface CompanyProfileProps {
  companyData: CompanySchema;
}

const CompanyProfile: React.FC<CompanyProfileProps> = ({ companyData }) => {
  const { profile, history } = companyData;

  return (
    <div className={clsx('company-profile')}>
      <div className={clsx('company-profile__section', 'company-profile__profile')}>
        <h2>公司簡介</h2>
        <p>{profile}</p>
      </div>
      <div className={clsx('company-profile__section', 'company-profile__history')}>
        <h2>歷史沿革</h2>
        <ul>
          {history.map((item, index) => (
            <li key={index}>
              <span>{item.year}</span>
              <span>{item.event}</span>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default CompanyProfile;