FROM golang as build

WORKDIR /app

COPY . ./build/

RUN cd build && \
  go mod tidy && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=build app/build/main ./
COPY --from=build app/build/cfg/ ./cfg/

ENTRYPOINT ["./main"]
