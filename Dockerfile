############################################################ 
# Dockerfile to build golang Installed Containers 

# Based on alpine

############################################################

FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM alpine:3.13

RUN mkdir /keel
COPY --from=builder /src/bin/linux/rule-manager /keel
COPY --from=builder /src/config.yml /keel


EXPOSE 31234
WORKDIR /keel
CMD ["/keel/rule-manager"]
