FROM golang:1.14.3-alpine AS build
WORKDIR /src
RUN apk update && apk add --no-cache git
COPY . .
RUN go get -d -v
RUN go build -o /out/real_estate .

FROM alpine
COPY --from=build /out/real_estate /bin/real_estate

ENTRYPOINT /bin/house_market
