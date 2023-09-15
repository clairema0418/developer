請幫我寫一個 go 小程式, 裡面有一隻 api,api 帶入生成的題目數量, 就可以 call 我的 python 程式, 讓 python 在背景執行題目生成, 直到所有題目生成完畢,python 程式會通知我的 go 程式, 將所有題目存到資料庫中, 並在 db 更新生成進度

技術架構:

1. 使用 Gin
2. 使用 Gorm
3. 使用 PostgresQL DB
4. 請使用 watermill 幫我設計 db 讀寫分離的架構, 讓我未來可以擴展
5. 請寫 makefile 跟 docker-compose file
6. 請寫 go.mod
7. 請寫 README.md 檔案說明專案架構
