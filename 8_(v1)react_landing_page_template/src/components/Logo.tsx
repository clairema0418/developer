import React from 'react';
import { logoData } from '../data/logoData';
import '../styles/logo.scss';

interface LogoProps {
  logo: string;
  slogan: string;
}

const Logo: React.FC<LogoProps> = ({ logo, slogan }) => {
  return (
    <div className="logo-container">
      <img src={logo} alt="公司品牌標誌" className="logo-image" />
      <h1 className="slogan">{slogan}</h1>
    </div>
  );
};

export const renderLogo = () => {
  return <Logo logo={logoData.logo} slogan={logoData.slogan} />;
};

export default Logo;