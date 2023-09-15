the app is: Generate Questions API

the files we have decided to generate are: main.go, api.go, question_generator.py, database.go, message.go, Makefile, docker-compose.yml, go.mod

Shared dependencies:
1. Function names: generateQuestions, callPythonScript, saveQuestionsToDatabase, updateProgressInDB
2. Message names: questionGenerationCompleted
3. Exported variables: questionCount
4. Data schemas: Question (id, content, created_at, updated_at)
5. ID names of DOM elements: N/A (no frontend involved)