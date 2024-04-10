# BASE
###############
FROM golang:1.22.2-alpine3.18 AS base

WORKDIR /app
ADD . .

RUN apk update && apk add build-base git
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -buildvcs=false -o main .

# PRODUCTION
###############
FROM scratch AS prod

COPY --from=base /app/main ./main
CMD ["./main"]