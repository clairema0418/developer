the app is: Golang package that takes a video file as input, uses ffmpeg to extract the audio, and converts the audio to text using Azure TTS. Each segment of text should be divided at punctuation marks, and the corresponding time in the audio file should be recorded. This will allow me to find the exact seconds in the audio and video files when a specific segment of text is searched for by the user.

the files we have decided to generate are: go.mod, Makefile

Shared dependencies:
1. go.mod
   - Dependency names: ffmpeg, Azure TTS SDK
2. Makefile
   - Function names: build, clean, run