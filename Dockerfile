FROM golang:1.23.4 AS build
WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct
ENV GOTOOLCHAIN=auto

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -trimpath -buildvcs=false -o /server .

FROM gcr.io/distroless/base-debian12
ENV PORT=8080

WORKDIR /srv
COPY --from=build /server /srv/server
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/srv/server"]
