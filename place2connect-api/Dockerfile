FROM golang:1.20-bullseye as builder

RUN go install golang.org/dl/go1.20@latest \
  && go1.20 download

WORKDIR /build

COPY go.mod .
COPY go.sum .
# COPY .env .
COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build -o ./app -tags timetzdata -trimpath cmd/http/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build -o ./app cmd/http/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=builder /build/app /app
# COPY --from=builder /build/.env /
EXPOSE 8000 8000

ENTRYPOINT ["/app"]