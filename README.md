# My Resume
[🚀🚀 Click here to see resume 🚀🚀](pdfs/naheed_rayhan_cv_1.pdf)

Note: This repo is also for testing purpose...



## File structure
```code 
resume/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── profile_pic
└── pdfs/
    └── naheed_rayhan_cv_1.pdf
```

<hr>

## Prerequisites

- Install ``golang``
- Clone the repo


## Running the server
```
go mod download
go run main.go
```
This will run on
```
http://localhost:8080
```



## Using dockerhub to pull image and run
```
docker pull naheed28/resume:1.0
docker run -d -p 8080:8080 --name pdf_server naheed28/resume:1.0
```
This will run on
```
http://localhost:8080
```
