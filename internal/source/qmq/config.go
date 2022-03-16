package qmq

type configT struct {
	name  string
	sink  string
	topic string
}

func (this *configT) Name() string {
	return this.name
}

func (this *configT) GetString(key string) string {
	switch key {
	case "name":
		return this.name
	case "sink":
		return this.sink
	case "topic":
		return this.topic
	default:
		return ""
	}
}
