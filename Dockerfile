FROM golang:latest

WORKDIR /app

COPY . ./
RUN go mod download

EXPOSE 8010

CMD ["bin/gothy", "serve"]