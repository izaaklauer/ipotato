#--------------------------------------------------------------------
# builder
#--------------------------------------------------------------------

FROM golang:1.19.3-alpine3.17 AS builder

WORKDIR /app-src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /tmp/ipotato cmd/ipotato-api/main.go

#--------------------------------------------------------------------
# final image
#--------------------------------------------------------------------

FROM alpine:3.17

# Config file will be delivered here at runtime by waypoint
RUN mkdir /opt/config

COPY --from=builder /tmp/ipotato /ipotato
EXPOSE 8080
CMD [ "/ipotato" ]