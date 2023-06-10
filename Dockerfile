FROM golang

EXPOSE 8000

WORKDIR /app
COPY go.mod /app/
RUN go mod download
RUN go install github.com/rubenv/sql-migrate/...@latest
COPY . /app/

RUN apt-get update && apt-get -y install sudo
RUN sudo apt-get -y install build-essential

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

CMD ["/app/main"]