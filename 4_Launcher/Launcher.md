請用 Electron+React 幫我寫一個 Launcher App, 這個 Laucher App 有兩個頁面

1. 登入頁面(包涵一個 Logo, 以及一個輸入框, 可以輸入金鑰), 按下確定按鈕後, 就會驗證金鑰是否合法, 如果金鑰合法, 就會跳轉到遊戲大廳
2. 遊戲大廳頁面, 可以透過金鑰, 拿取這個使用者的所有遊戲, 點選遊戲後, 會跳出介紹頁面
3. 遊戲介紹頁面, 可以放遊戲圖片, 以及遊戲說明, 有按鈕點選下載遊戲或是開啟遊戲
4. 點選下載遊戲按鈕後, 透過金鑰抓取目前使用者的遊戲版本, 到 aws s3 上面下載遊戲

其他需求:

1. 請用 scss 撰寫
2. 要可以支援手機, 平板, 以及桌面等多種裝置
3. 使用 rtk query 串接 api

請用幫我產生一個登入頁面, 登入頁面包涵 Logo, 可輸入金鑰, , 按下確定按鈕後, 就會驗證金鑰是否合法, 如果金鑰合法, 就會跳轉到遊戲大廳
其他需求:

1. 請用 scss 撰寫
2. 要可以支援手機, 平板, 以及桌面等多種裝置
3. 使用 rtk query 串接 api
4. 使用 typescript

Please help me create a react app, with login page that includes a logo and an input field for entering a key. Upon clicking the "Confirm" button, it should validate the key's authenticity. If the key is valid, it should redirect to the game lobby.

Additional requirements:

Write styles using SCSS.
Support multiple devices such as mobile, tablet, and desktop.
Use RTK Query to connect with APIs.
Use TypeScript.

---

Please help me create a Launcher App using Electron and React. The Launcher App should have two pages:

Login page: It should include a logo and an input field where users can enter a key. Upon clicking the "Confirm" button, the app should validate the key's authenticity. If the key is valid, it should redirect to the game lobby.

Game lobby page: Users can retrieve a list of all their games using the key. They can click on a game to view its introduction.

Game introduction page: It should display the game's image and description, along with buttons to download or launch the game.

After clicking the download button, the app should use the key to fetch the user's game version and download the game from AWS S3.

Additional requirements:

Write styles using SCSS.
Support multiple devices such as mobile, tablet, and desktop.
Use RTK Query to connect with APIs.
