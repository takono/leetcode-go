FROM golang:1.22

RUN apt-get update && apt-get install -y \
    git \
    curl \
    vim \
    && rm -rf /var/lib/apt/lists/*

ENV GO111MODULE=on