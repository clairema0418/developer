the app is: 使用 next 開發一個做產品的形象官網
---

the files we have decided to generate are: 
- pages/index.tsx
- components/Logo.tsx
- components/Product.tsx
- components/Scene.tsx
- components/Feature.tsx
- components/Testimonial.tsx
- components/Contact.tsx
- components/Company.tsx
- components/News.tsx
- components/Download.tsx
- components/Footer.tsx
- styles/index.scss
- styles/components.scss
- public/static/images
- lib/seo.ts

Now that we have a list of files, we need to understand what dependencies they share.
Please name and briefly describe what is shared between the files we are generating, including exported variables, data schemas, id names of every DOM elements that javascript functions will use, message names, and function names.
Exclusively focus on the names of the shared dependencies, and do not add any other explanation.

Shared Dependencies:
- Exported Variables: logoData, productData, sceneData, featureData, testimonialData, contactData, companyData, newsData, downloadData
- Data Schemas: LogoSchema, ProductSchema, SceneSchema, FeatureSchema, TestimonialSchema, ContactSchema, CompanySchema, NewsSchema, DownloadSchema
- ID Names: logo, product, scene, feature, testimonial, contact, company, news, download
- Message Names: logoMessage, productMessage, sceneMessage, featureMessage, testimonialMessage, contactMessage, companyMessage, newsMessage, downloadMessage
- Function Names: renderLogo, renderProduct, renderScene, renderFeature, renderTestimonial, renderContact, renderCompany, renderNews, renderDownload