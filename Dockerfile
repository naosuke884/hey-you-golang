FROM golang:1.24.0-bullseye AS build
COPY . /work
WORKDIR /work
RUN go mod download
RUN CGO_ENABLE=0 go build -trimpath -ldflags='-s -w' -o main ./cmd/serve/main.go
RUN chmod +x main

FROM ubuntu:24.04 AS production
WORKDIR /work/
COPY --from=build /work/main ./main
USER ubuntu
CMD ["./main"]