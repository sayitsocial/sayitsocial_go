# START container setup
FROM golang:1.15.5-alpine3.12 AS build

# END container setup

WORKDIR /build
ADD . ./
RUN go build main.go


FROM alpine

WORKDIR /dist

COPY --from=build /build/main /dist/
COPY --from=build /build/pkg/database/migrations /dist/pkg/database/migrations
COPY --from=build /build/swagger/ /dist/swagger
RUN chmod +x /dist/main

EXPOSE 8000

ENTRYPOINT ["/bin/sh", "-c", "/dist/main"]
