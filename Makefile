app_run: build run

build:
	docker build -t note_server -f Dockerfile .

run: 
	docker run -p 8080:8080 note_server