import { NextSeo } from 'next-seo';

interface SEOProps {
  title: string;
  description: string;
  canonical?: string;
  openGraph?: {
    url: string;
    title: string;
    description: string;
    images: Array<{
      url: string;
      width: number;
      height: number;
      alt: string;
    }>;
  };
}

export const SEO = ({ title, description, canonical, openGraph }: SEOProps) => {
  return (
    <NextSeo
      title={title}
      description={description}
      canonical={canonical}
      openGraph={openGraph}
    />
  );
};