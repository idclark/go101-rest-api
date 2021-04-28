from golang:1.16 as builder

run mkdir /app
add . /app
workdir /app

run cgo_enabled=0 goos=linux go build -o app cmd/server/main.go

from alpine:latest as production
copy --from=builder /app .
cmd ["./app"]