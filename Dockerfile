FROM golang:alpine AS builder

WORKDIR /usr/src/todo
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go mod tidy && go build -v -o main cmd/main.go && chmod +x ./main

FROM scratch
COPY --from=builder /usr/src/todo/main .
ENTRYPOINT [ "./main" ]