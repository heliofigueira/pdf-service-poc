PDF Service POC
=============================

The idea behind this project is prove the capabilities of using GO as a PDF service.

## Installation

To run this application, it's necessary to have `go` 1.18+ and `docker`.

```sh
go build
./pdf-service-poc
```


remove result folders
```sh
rm -rf examples/result

mkdir examples/result
mkdir examples/result/pdf
mkdir examples/result/split 
mkdir examples/result/split/example 
mkdir examples/result/split/generated 
mkdir examples/result/merged      
mkdir examples/result/jpg
```