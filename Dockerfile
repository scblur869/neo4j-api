FROM golang:alpine AS build-env
RUN apk --no-cache add build-base gcc git ca-certificates
ENV GOPROXY=direct
ADD . /src
RUN cd /src && go build -o neo4j-api

# final stage
FROM alpine
WORKDIR /app

#graph db
ENV GRAPHDB_URI=http://localhost:7474
ENV USERNAME=
ENV PASSWORD=
ENV AESKEY =
ENV PORT=4000
ENV GIN_MODE=debug
COPY --from=build-env /src/neo4j-api /app/
EXPOSE 4000
CMD ["./neo4j-api"]