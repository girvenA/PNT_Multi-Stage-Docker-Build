FROM golang:1.23.2-alpine3.20 as base

FROM base as builder

WORKDIR /app
COPY . .

RUN go get
RUN go build -o ./bin/app ./main.go

FROM base as runner
COPY --from=builder /app/bin/app /app/bin/app

EXPOSE 3000
ENTRYPOINT [ "/app/bin/app" ]