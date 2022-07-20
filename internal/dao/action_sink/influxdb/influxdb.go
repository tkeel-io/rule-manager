package influxdb

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log "github.com/tkeel-io/rule-util/pkg/commonlog"
)

const InfluxdbSinkId = "InfluxdbSinkId"

func PingInfluxdb(url, token, org, bucket string) bool {
	url = "http://" + url
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	err := writeAPI.WriteRecord(context.Background(), "ping,test=test test=1")
	if err != nil {
		log.Error(err)
	}
	return err == nil
}
