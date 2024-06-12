FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/WatchYourPorts/ && CGO_ENABLED=0 go build -o /WatchYourPorts .


FROM scratch

WORKDIR /data/WatchYourPorts
WORKDIR /app

COPY --from=builder /WatchYourPorts /app/

ENTRYPOINT ["./WatchYourPorts"]