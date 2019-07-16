###############################
# Builder container
###############################

From golang:1.9 as builder
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

From alpine:latest
EXPOSE 8080
COPY --from=builder /go/src/app/app /app
CMD ["/app"]
