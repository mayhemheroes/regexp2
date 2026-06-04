FROM golang:1.23 AS builder

ENV GOTOOLCHAIN=auto

WORKDIR /src
ADD . .

RUN go test -c ./syntax -o /syntax-fuzz.libfuzzer

FROM ubuntu:22.04
COPY --from=builder /syntax-fuzz.libfuzzer /
