FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD cmd ./cmd
ADD pkg ./pkg

WORKDIR /app/cmd
RUN go build -o /sample-orch

EXPOSE 3000

CMD [ "/sample-orch" ]