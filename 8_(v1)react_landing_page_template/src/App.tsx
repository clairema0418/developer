import React from 'react';
import Head from 'next/head';
import { Logo } from '../components/Logo';
import { Product } from '../components/Product';
import { Scene } from '../components/Scene';
import { Feature } from '../components/Feature';
import { Testimonial } from '../components/Testimonial';
import { Contact } from '../components/Contact';
import { Company } from '../components/Company';
import { News } from '../components/News';
import { Download } from '../components/Download';
import { Footer } from '../components/Footer';
import '../styles/main.scss';

const App: React.FC = () => {
  return (
    <>
      <Head>
        <title>產品形象官網</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="description" content="產品形象官網，包含公司品牌標誌、產品照片、產品使用場景等" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="container">
        <Logo />
        <Product />
        <Scene />
        <Feature />
        <Testimonial />
        <Contact />
        <Company />
        <News />
        <Download />
        <Footer />
      </div>
    </>
  );
};

export default App;