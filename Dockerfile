FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETARCH

WORKDIR /src
RUN apk --no-cache add git ca-certificates
COPY . .

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN GOOS=linux GOARCH=$TARGETARCH go build -o /src/mozhi

FROM alpine:3.16 as bin

WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=build /src/mozhi .

EXPOSE 3000

CMD ["/app/mozhi", "serve"]
