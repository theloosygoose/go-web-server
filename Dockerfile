# syntax=docker/dockerfile:1

FROM golang
COPY . ./bin/main
RUN make build
CMD
