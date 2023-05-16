FROM golang:1.20-alpine3.18 AS builder


WORKDIR /app

# COPY /place2connect-api/wait-for.sh .
COPY wait-for.sh .
COPY app.env .
COPY ./place2connect-api/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build  -v -a -installsuffix cgo -o ./main -tags timetzdata -trimpath cmd/http/main.go

FROM alpine:3.18  

WORKDIR /app
COPY --from=builder /app/main .

COPY   app.env .
COPY   start.sh .
COPY   wait-for.sh .
# COPY   /place2connect-api/start.sh .
# COPY   /place2connect-api/wait-for.sh .
RUN mkdir -p ./images
# COPY   /place2connect-api/images ./images
# COPY   /place2connect-api/images ./srv/images

# EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT ["/app/start.sh" ]