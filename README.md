# Filter bad words

[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

filter chat with go

This project uses a json file, data.json, to determine topics in a particular text and outputs topics of input text.


## Getting started

This project is written in golang and python 3

Build the go module filter.go by running:
```
go build -buildmode=c-shared -o filter.so filter.go
```
filter.so and filter.h files will be created.

Start the python server with python 3, and specify the port to run on as a command-line argument.

By default the server starts at port 8080.
```
python filter.py 8080
```
After the server has initialized, and starts listening to the port, on another terminal, run the request.go to send requests to the server.

Specify the port in the arguments.

```
go run request.go 8080
```
Requests will be sent to the server, and the response will be the result.

This will be json text with 'text' and 'topics'.




## Results

The model is fast, and runs at a speed of 200 queries per second.
