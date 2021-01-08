FROM golang:1.10
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/shri7756/batman
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["batman"]
