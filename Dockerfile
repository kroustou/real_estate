FROM golang:1.14.3-alpine AS build
WORKDIR /src
RUN apk update && apk add --no-cache git
COPY . .
RUN go get -d -v
RUN go build -o /out/house_market .

FROM alpine
COPY --from=build /out/house_market /bin/house_market

ENTRYPOINT /bin/house_market
