FROM golang:1.16-alpine AS dev

RUN apk add --no-cache make git gcc libc-dev
RUN go get -u github.com/cosmtrek/air
RUN go get -v github.com/rubenv/sql-migrate/...
RUN mkdir -p /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod dowloand 

COPY . .

EXPOSE 3000

CMD ["make", "dev"]