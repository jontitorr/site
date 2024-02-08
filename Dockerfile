FROM golang:1.21 AS build
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint cmd/site/main.go

FROM gcr.io/distroless/static-debian11 AS final
WORKDIR /
COPY --from=build /entrypoint /entrypoint
COPY --from=build /app/public/static /public/static
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]