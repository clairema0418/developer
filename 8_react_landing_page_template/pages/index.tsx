import React from 'react';
import Head from 'next/head';
import { useSelector, useDispatch } from 'react-redux';
import { fetchProductData, fetchCustomerReviews, fetchCaseStudies, fetchNewsUpdates, fetchCompanyEvents } from '../store/apiSlice';
import BrandLogo from '../components/BrandLogo/BrandLogo';
import ProductPhotos from '../components/ProductPhotos/ProductPhotos';
import ProductSpecifications from '../components/ProductSpecifications/ProductSpecifications';
import UsageScenarios from '../components/UsageScenarios/UsageScenarios';
import ProductFeatures from '../components/ProductFeatures/ProductFeatures';
import CustomerReviews from '../components/CustomerReviews/CustomerReviews';
import CaseStudies from '../components/CaseStudies/CaseStudies';
import CustomerService from '../components/CustomerService/CustomerService';
import NewsUpdates from '../components/NewsUpdates/NewsUpdates';
import CompanyEvents from '../components/CompanyEvents/CompanyEvents';
import ContactUs from '../components/ContactUs/ContactUs';
import JoinUs from '../components/JoinUs/JoinUs';

const IndexPage: React.FC = () => {
  const dispatch = useDispatch();

  React.useEffect(() => {
    dispatch(fetchProductData());
    dispatch(fetchCustomerReviews());
    dispatch(fetchCaseStudies());
    dispatch(fetchNewsUpdates());
    dispatch(fetchCompanyEvents());
  }, [dispatch]);

  return (
    <>
      <Head>
        <title>Product Image Website</title>
        <meta name="description" content="A product image website showcasing various products and their features." />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <main>
        <BrandLogo />
        <ProductPhotos />
        <ProductSpecifications />
        <UsageScenarios />
        <ProductFeatures />
        <CustomerReviews />
        <CaseStudies />
        <CustomerService />
        <NewsUpdates />
        <CompanyEvents />
        <ContactUs />
        <JoinUs />
      </main>
    </>
  );
};

export default IndexPage;