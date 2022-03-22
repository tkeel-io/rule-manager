FROM alpine:3.13
RUN mkdir /iot
ADD bin/linux/rule-manager /iot
WORKDIR /iot
CMD ["/iot/rule-manager"]
