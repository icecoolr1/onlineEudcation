FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY Student/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/studentapi Student/api/.\studentapi.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/studentapi /app/studentapi
COPY --from=builder /app/etc /app/etc

CMD ["./studentapi", "-f", "etc/studentapi.yaml"]
