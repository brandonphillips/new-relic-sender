FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build

ENTRYPOINT ["/app/main"]