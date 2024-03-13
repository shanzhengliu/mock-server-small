FROM golang:1.20.6-alpine3.18 AS BuildStage

WORKDIR /app

RUN apk add --no-cache upx

COPY . .

RUN go mod download

RUN go build -ldflags "-s -w" -o /app/main .

RUN upx --brute /app/main


FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=BuildStage app/ app/

EXPOSE 8080

ENTRYPOINT [ "app/main" ]
