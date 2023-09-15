請幫我使用 next 開發一個做產品的形象官網, 裡面要包涵

- 公司品牌標誌和標語
- 產品照片和規格
- 產品使用場景和示例
- 產品特色和優勢介紹
- 客戶評價和案例分享
- 客戶服務和聯繫方式
- 公司簡介和歷史沿革
- 新聞動態和公司活動
- 下載中心和資源分享
- 聯繫我們和加入我們的選項

其他需求:

1. 使用 scss 開發
2. 使用 typescript 開發
3. 使用 @reduxjs/toolkit 串接 API
4. 使用 clsx, emotion
5. 使用 mui, 但是 layout 請使用 scss
6. 請製作手機/平板/網頁 3 個版本
7. 幫我處理好 SEO
8. 請幫我寫好 apiSlice, 並且都先幫我使用 json 產生好測試資料, 放在 config 中
9. 所有的 page 頁面幫我放在/src/page 底下,scss 跟 tsx 放在同一個資料夾

## File Structure

```
.
├── public
│   ├── index.html
│   ├── robots.txt
│   └── manifest.tsxon
├── src
│   ├── react-app-env.d.ts
│   ├── App.tsx
│   ├── App.scss
│   ├── index.tsx
│   ├── assets
│   ├── components
│   ├── features
│   ├── hooks
│   ├── lib
│   ├── pages
│   ├── locales
│   ├── config
│   ├── types
│   │   ├── interfaces
│   │       └── index.ts
│   ├── utils
│   └── services
│       ├── api
│       │   └── apiSlice.ts
│       ├── middlewares
│       └── store.ts
├── .gitignore
├── .env.local
├── env.example
├── .tool-versions
├── tsconfig.json
├── package.json
└── README.md
```
