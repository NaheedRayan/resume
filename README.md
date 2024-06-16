# My Resume
[My resume](Naheed_Rayhan_CV_1.pdf)



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




Using docker to pull image and run
```
docker pull naheed28/resume:1.0
docker run -d -p 8080:8080 --name pdf_server naheed28/resume:1.0
```