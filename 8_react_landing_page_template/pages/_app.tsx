import React from 'react';
import { Provider } from 'react-redux';
import { store } from '../store';
import { AppProps } from 'next/app';
import Head from 'next/head';
import '../styles/global.scss';

const MyApp: React.FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <>
      <Head>
        <title>Product Image Website</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta name="description" content="A product image website showcasing various products and their features." />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    </>
  );
};

export default MyApp;