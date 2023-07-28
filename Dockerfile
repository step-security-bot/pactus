FROM alpine:3.18@sha256:82d1e9d7ed48a7523bdebc18cf6290bdb97b82302a8a9c27d4fe885949ea94d1 as builder

RUN apk add --no-cache git go gmp-dev build-base g++ openssl-dev
ADD . /pactus

# Building pactus-daemon
RUN cd /pactus && \
    go build -ldflags "-s -w" -o ./build/pactus-daemon ./cmd/daemon


## Copy binary files from builder into second container
FROM alpine:3.15@sha256:3362f865019db5f14ac5154cb0db2c3741ad1cce0416045be422ad4de441b081

COPY --from=builder /pactus/build/pactus-daemon /usr/bin

ENV WORKING_DIR "/pactus"

VOLUME $WORKING_DIR
WORKDIR $WORKING_DIR

ENTRYPOINT ["pactus-daemon"]
