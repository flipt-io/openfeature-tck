FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
COPY cmd ./cmd
RUN CGO_ENABLED=0 go build -trimpath -buildvcs=false -mod=mod -ldflags="-s -w" -o /out/control ./cmd/control

FROM ghcr.io/flipt-io/flipt:v2
COPY --from=build /out/control /control
EXPOSE 8080 8888
USER flipt
ENTRYPOINT ["/control"]
CMD []
