FROM golang:1.18.2-bullseye

RUN apt-get update

RUN apt-get -y install ghostscript

RUN apt-get -y install libgs-dev

RUN apt -y install wget

RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb
RUN apt -y install ./wkhtmltox_0.12.6-1.buster_amd64.deb

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]