FROM golang:alpine AS build

ARG CGO=1
ENV CGO_ENABLED=${CGO}
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR /go/src/github.com/bhushan-mdn/words-api

COPY . /go/src/github.com/bhushan-mdn/words-api/

# gcc/g++ are required to build SASS libraries for extended version
# RUN apk update && \
#     apk add --no-cache gcc g++ musl-dev

RUN go build

# ---

FROM alpine:latest

COPY --from=build /go/src/github.com/bhushan-mdn/words-api/words-api /usr/bin/words-api

# libc6-compat & libstdc++ are required for extended SASS libraries
# ca-certificates are required to fetch outside resources (like Twitter oEmbeds)
# RUN apk update && \
#     apk add --no-cache ca-certificates libc6-compat libstdc++ git

# VOLUME /site
# WORKDIR /site

ENV SERVER_PORT=5000

# Expose port for live server
EXPOSE 5000-8000

ENTRYPOINT ["words-api"]
CMD ["serve"]