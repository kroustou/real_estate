FROM golang:alpine AS builder

ENV GOPATH=/
RUN apk update && apk add --no-cache git

COPY go.mod .
COPY go.sum .

RUN env
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
# mac
RUN go build -o /go/bin/real_estate-mac
RUN chmod +x /go/bin/real_estate-mac
# arm
RUN GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-w -s" -o /go/bin/real_estate-arm
RUN chmod +x /go/bin/real_estate-arm

FROM scratch as mac
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/real_estate-mac /go/bin/real_estate

FROM mac as arm
COPY --from=builder /go/bin/real_estate-arm /go/bin/real_estate

ENTRYPOINT ["/go/bin/real_estate"]
