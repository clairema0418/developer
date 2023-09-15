the app is: YouTube Audio to Text Converter

the files we have decided to generate are: main.go, config.go, models.go, routes.go, controllers.go, services.go, middleware.go, message_handlers.go, message_producers.go

Shared dependencies:

1. Gin (Web framework)
2. Gorm (ORM for database operations)
3. PostgreSQL (Database)
4. Watermill (Message broker for DB read-write separation)

Exported variables:
- db (Database connection)
- router (Gin router)

Data schemas:
- YouTubeVideo (ID, URL, AudioFilePath, Transcription)

ID names of DOM elements:
- input_youtube_url (Input field for YouTube URL)
- btn_submit (Submit button)
- div_transcription (Div to display the transcription)

Message names:
- youtube_video_created (Message for new YouTube video entry)
- youtube_audio_extracted (Message for extracted audio)
- youtube_transcription_completed (Message for completed transcription)

Function names:
- main (Entry point)
- initConfig (Initialize configuration)
- initDB (Initialize database connection)
- initRouter (Initialize Gin router)
- initMessageHandlers (Initialize Watermill message handlers)
- initMessageProducers (Initialize Watermill message producers)
- createYouTubeVideo (Create YouTube video entry)
- extractYouTubeAudio (Extract audio from YouTube video)
- transcribeAudio (Transcribe audio using Whisper API)
- displayTranscription (Display transcription on the web page)