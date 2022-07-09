FROM golang as build

WORKDIR /app

RUN mkdir build

COPY . ./build/

RUN cd build && \
  go mod tidy && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

FROM alpine

WORKDIR /app

COPY --from=build app/build/main ./
COPY --from=build app/build/*.json ./
COPY --from=build app/build/*.env ./

ENTRYPOINT ["./main"]
