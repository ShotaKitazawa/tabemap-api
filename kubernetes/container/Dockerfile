FROM alpine:latest

EXPOSE 8080

COPY --from=builder tabemap-api /tabemap-api

ENTRYPOINT ["/tabemap-api"]
