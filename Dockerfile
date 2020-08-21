FROM golang:1.14-alpine3.12 AS builder

WORKDIR /go/src/app
COPY ./pkg ./pkg
COPY ./cmd ./cmd
COPY ./internal ./internal
#COPY . .

#RUN apt-get update && apt-get upgrade
RUN apk -U upgrade && \
    apk add sqlite && \
    apk add build-base && \
    apk add git

RUN go get -d -v ./...
RUN go install -v ./...

RUN ls -lha
# Populate the SQLite db
RUN popdb

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /go/src/app /app
COPY --from=builder /go/bin /usr/bin

# Put the name of the executable here
CMD ["app"]
