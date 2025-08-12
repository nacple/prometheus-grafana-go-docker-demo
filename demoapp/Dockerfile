FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o demoapp

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/demoapp /app/demoapp
EXPOSE 2112
CMD ["/app/demoapp"]

