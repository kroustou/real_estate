FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/house_market .

FROM alpine
COPY --from=build /out/house_market /bin/house_market

ENTRYPOINT /bin/house_market
