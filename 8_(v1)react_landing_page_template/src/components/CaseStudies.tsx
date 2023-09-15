import React from 'react';
import { CaseStudy } from '../types';
import '../styles/case-studies.scss';

interface CaseStudiesProps {
  caseStudies: CaseStudy[];
}

const CaseStudies: React.FC<CaseStudiesProps> = ({ caseStudies }) => {
  return (
    <section id="case-studies" className="case-studies">
      <h2>客戶評價和案例分享</h2>
      <div className="case-studies-container">
        {caseStudies.map((caseStudy, index) => (
          <div key={index} className="case-study">
            <img src={caseStudy.image} alt={caseStudy.title} />
            <h3>{caseStudy.title}</h3>
            <p>{caseStudy.description}</p>
          </div>
        ))}
      </div>
    </section>
  );
};

export default CaseStudies;