FROM golang:1.18.2-bullseye

RUN apt-get update

RUN apt-get -y install ghostscript

RUN apt-get -y install libgs-dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]