###############################
# Builder container
###############################

FROM golang:1.12.7 AS builder
RUN apt-get update
WORKDIR /go/src/app
COPY . .

# Install dependencies...
RUN go-wrapper download
RUN go-wrapper install

# Compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


###############################
# Exec container
###############################

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /go/src/app/app /app
CMD ["/app"]
