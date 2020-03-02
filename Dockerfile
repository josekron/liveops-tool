FROM golang:alpine
RUN mkdir /go/src/liveops-tool
ADD . /go/src/liveops-tool/
WORKDIR /go/src/liveops-tool
COPY . /go/src/liveops-tool
EXPOSE 8080
CMD ["go","mod", "tidy"]
CMD ["go", "run", "/go/src/liveops-tool/main.go"]