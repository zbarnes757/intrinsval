FROM alpine:latest

RUN mkdir web && mkdir web/dist
COPY web/dist/ web/dist/

COPY main /

CMD ["./main"]
