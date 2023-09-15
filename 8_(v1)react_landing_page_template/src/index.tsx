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
import '../styles/index.scss';

const IndexPage: React.FC = () => {
  return (
    <>
      <Head>
        <title>產品形象官網</title>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta name="description" content="產品形象官網，包含公司品牌標誌、產品照片、規格、使用場景、特色、優勢介紹、客戶評價、案例分享、客戶服務、聯繫方式、公司簡介、歷史沿革、新聞動態、公司活動、下載中心、資源分享、聯繫我們和加入我們的選項。" />
      </Head>
      <main>
        <Logo />
        <Product />
        <Scene />
        <Feature />
        <Testimonial />
        <Contact />
        <Company />
        <News />
        <Download />
      </main>
      <Footer />
    </>
  );
};

export default IndexPage;