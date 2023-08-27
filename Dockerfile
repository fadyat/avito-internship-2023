FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download && go mod verify

ARG VERSION
RUN cd cmd/segment && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-X 'main.Version=${VERSION}' -s -w" \
    -o /main .

FROM gcr.io/distroless/static-debian11
COPY --from=build /main /app/${MIGRATIONS_PATH} ./
CMD ["./main"]