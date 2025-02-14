FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
COPY .env.example .env

RUN go mod tidy

RUN go build -o ./api -v main.go
RUN chmod +x ./api

EXPOSE 9005

CMD [ "./api" ]