import React from 'react';
import clsx from 'clsx';
import { contactData } from '../data/contactData';
import '../styles/contact-info.scss';

interface ContactInfoProps {
  className?: string;
}

const ContactInfo: React.FC<ContactInfoProps> = ({ className }) => {
  return (
    <div className={clsx('contact-info', className)}>
      <h2 className="contact-info__title">客戶服務和聯繫方式</h2>
      <div className="contact-info__content">
        {contactData.map((item, index) => (
          <div key={index} className="contact-info__item">
            <h3 className="contact-info__item-title">{item.title}</h3>
            <p className="contact-info__item-description">{item.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ContactInfo;