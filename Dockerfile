FROM golang as build

WORKDIR /app

RUN mkdir build

COPY go.mod ./build
COPY go.sum ./build
COPY *.go ./build

RUN cd build && \
  go mod download && \
  go build main.go

FROM alpine

WORKDIR /app

COPY --from=build app/build/main ./
CMD ["./main"]
