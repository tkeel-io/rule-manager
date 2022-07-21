package influxdb

import "testing"

func TestPingInfluxdb(t *testing.T) {
	type args struct {
		url    string
		token  string
		org    string
		bucket string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{url: "http://127.0.0.1:8086", token: "2eOYQbPf5Rw0J1U31cS_1W8jUUzZTcXy7_K6udM7tpZyQjz-rG4QR3ECQXpXdEx-yOlQsvSCW4l2YhEiYHdK2g==", org: "org", bucket: "bucket"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PingInfluxdb(tt.args.url, tt.args.token, tt.args.org, tt.args.bucket); got != tt.want {
				t.Errorf("PingInfluxdb() = %v, want %v", got, tt.want)
			}
		})
	}
}
