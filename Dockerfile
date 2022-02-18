FROM golang:latest

LABEL container="malli1998 <mallikarjunaa49@gmail.com>"
RUN mkdir /RoverApp
ADD . /RoverApp
WORKDIR /RoverApp

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8001
RUN go build

CMD ["./awesomeProject1"]