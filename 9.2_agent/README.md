# Audio Processing App

This is a Go program that allows you to segment an audio file, convert each segment into text, convert the text into vectors, and store them in a database. It also uses WebSocket to send the user the current processing progress.

## Project Structure

- `main.go`: The main entry point of the application.
- `audio_segmentation.go`: Contains the logic for segmenting audio files.
- `audio_to_text.go`: Contains the logic for converting audio segments to text.
- `text_to_vector.go`: Contains the logic for converting text segments to vectors.
- `database.go`: Handles database connections and operations using Gorm ORM and PostgreSQL.
- `websocket.go`: Manages WebSocket connections and sending progress updates to the user.
- `pubsub.go`: Implements the pub/sub pattern using Watermill for communication between different parts of the application.
- `Makefile`: Contains build and run commands for the application.
- `docker-compose.yml`: Docker Compose configuration for running the application and its dependencies.
- `go.mod`: Go module file containing the project's dependencies.
- `README.md`: This file, which provides an overview of the project and its structure.

## Technical Requirements

- Gin framework: Used for handling HTTP requests and routing.
- Gorm ORM: Used for database operations with PostgreSQL.
- PostgreSQL: The database used for storing audio segments, text segments, and vector segments.
- Watermill: A Go library for working with message streams and event-driven architectures, used for implementing the pub/sub pattern.
- Makefile: Provides build and run commands for the application.
- Docker Compose: Used for running the application and its dependencies in containers.

## How to Run

1. Install Docker and Docker Compose on your machine.
2. Clone this repository and navigate to the project directory.
3. Run `make build` to build the application.
4. Run `docker-compose up` to start the application and its dependencies.
5. Open your browser and navigate to `http://localhost:8080` to access the application.

## Usage

1. Upload an audio file using the provided interface.
2. The application will segment the audio file, convert each segment into text, convert the text into vectors, and store them in the database.
3. The progress of the processing will be sent to the user via WebSocket and displayed on the page.