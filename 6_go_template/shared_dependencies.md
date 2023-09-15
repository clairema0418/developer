the app is: Golang 應用程式範本 with account CRUD API, RestAPI, Gin, Gorm, PostgresQL DB, and watermill for db 讀寫分離

the files we have decided to generate are: main.go, config.go, models.go, handlers.go, middleware.go, router.go, repository.go, messages.go, subscriber.go, publisher.go

Shared dependencies:
1. Gin: Web framework used in main.go, handlers.go, middleware.go, and router.go
2. Gorm: ORM library used in models.go and repository.go
3. PostgresQL: Database used in config.go and repository.go
4. Watermill: Messaging library used in messages.go, subscriber.go, and publisher.go
5. Account: Data schema used in models.go, handlers.go, and repository.go
6. AccountID: ID name used in models.go, handlers.go, and repository.go
7. CreateAccount, UpdateAccount, DeleteAccount, GetAccount, ListAccounts: Function names used in handlers.go and repository.go
8. AccountCreated, AccountUpdated, AccountDeleted: Message names used in messages.go, subscriber.go, and publisher.go
9. AccountRepository: Interface used in repository.go and handlers.go
10. NewAccountRepository: Function name used in repository.go and main.go
11. SetupRouter: Function name used in router.go and main.go
12. SetupMiddleware: Function name used in middleware.go and main.go