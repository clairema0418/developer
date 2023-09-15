the app is: Go program to segment audio, convert to text, convert text to vectors, store in database, and send progress via WebSocket

the files we have decided to generate are: main.go, audio_segmenter.go, text_converter.go, vector_converter.go, database.go, websocket.go, pubsub.go, Makefile, docker-compose.yml, go.mod, README.md

Shared dependencies:

1. Exported variables:
   - audioSegmenter
   - textConverter
   - vectorConverter
   - dbConnection
   - websocketConnection
   - pubsubConnection

2. Data schemas:
   - AudioSegment
   - TextSegment
   - VectorSegment

3. ID names of DOM elements:
   - progressContainer
   - progressBar

4. Message names:
   - audioSegmented
   - textConverted
   - vectorConverted
   - progressUpdate

5. Function names:
   - segmentAudio
   - convertText
   - convertVector
   - storeInDatabase
   - sendProgress
   - setupWebSocket
   - setupPubSub
   - main