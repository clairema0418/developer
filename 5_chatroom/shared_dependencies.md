the app is: 使用 go, 幫我使用 websocket 實作一個聊天室後端的功能。

the files we have decided to generate are: main.go, chatroom.go, user.go, message.go, websocket.go, router.go, middleware.go, config.go, database.go, pubsub.go

Shared dependencies:

1. Exported variables:
   - db *gorm.DB
   - router *gin.Engine
   - pub *watermill.PubSub

2. Data schemas:
   - ChatRoom struct
   - User struct
   - Message struct

3. ID names of DOM elements (not applicable, as this is a backend application)

4. Message names:
   - "chatroom_created"
   - "user_joined"
   - "user_left"
   - "message_sent"

5. Function names:
   - main()
   - setupDatabase()
   - setupRouter()
   - setupMiddleware()
   - setupPubSub()
   - createChatRoom()
   - joinChatRoom()
   - leaveChatRoom()
   - sendMessage()
   - handleMessage()
   - handleChatRoomCreated()
   - handleUserJoined()
   - handleUserLeft()
   - handleMessageSent()