FROM golang:1.21.5 AS builder

ENV ROOT=/go/src

WORKDIR ${ROOT}

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd


FROM alpine:3.19.0 AS deploy

ENV ROOT=/go/src

WORKDIR ${ROOT}

COPY --from=builder ${ROOT}/main .

CMD [ "/go/src/main" ]
