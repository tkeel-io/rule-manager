############################################################ 
# Dockerfile to build golang Installed Containers 

# Based on alpine

############################################################

FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM alpine:3.13

RUN mkdir /keel
COPY --from=builder /src/bin/linux/rule-manager /keel


EXPOSE 31234
WORKDIR /keel
CMD ["/keel/rule-manager"]
