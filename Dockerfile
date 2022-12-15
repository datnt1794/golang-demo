FROM golang:1.18-alpine
LABEL maintainer="Me"


# Setting up Dev environment

WORKDIR /app
COPY .  /app
RUN go mod download

EXPOSE 8000
CMD [ "go", "run", "cmd/main.go"]
