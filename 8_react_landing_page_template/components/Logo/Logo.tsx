import React from 'react';
import Image from 'next/image';
import styles from './Logo.module.scss';

interface LogoProps {
  slogan: string;
}

const Logo: React.FC<LogoProps> = ({ slogan }) => {
  return (
    <div className={styles.logoContainer}>
      <Image src="/logo.svg" alt="Company Logo" width={150} height={50} />
      <h1 className={styles.slogan}>{slogan}</h1>
    </div>
  );
};

export default Logo;