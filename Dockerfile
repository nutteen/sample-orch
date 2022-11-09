FROM golang:1.16-alpine

WORKDIR /app
ENV GONOSUMDB gitdev.inno.ktb/,gitdev.devops.krungthai.com/
ENV GOPROXY https://artifact.devops.krungthai.com/repository/goproxy,https://proxy.golang.or
COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD cmd ./cmd
ADD pkg ./pkg

WORKDIR /app/cmd
RUN go build -o /sample-orch

#Set default timezone
RUN apk add --no-cache tzdata
ENV TZ Asia/Bangkok
RUN ln -sf /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" > /etc/timezone

EXPOSE 3000

CMD [ "/sample-orch" ]