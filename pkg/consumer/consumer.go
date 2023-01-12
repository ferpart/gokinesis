package consumer

import (
	"encoding/json"
	"errors"

	"github.com/mjneil/kinesis-producer/deaggregation"

	"github.com/ferpart/gokinesis/pkg/kinesis"
)

type Consumer struct {
	recordMap map[string][]map[string]interface{}
	recordKey string
}

func New(recordKey string) *Consumer {
	return &Consumer{
		recordMap: make(map[string][]map[string]interface{}, 0),
		recordKey: recordKey,
	}
}

func (c *Consumer) Consume(k *kinesis.Kinesis, stopConsume <-chan bool, storeChan chan *Consumer) error {
	for {
		select {
		case <-stopConsume:
			storeChan <- c
			return nil
		default:
			gro, _ := k.GetRecordsOutput()
			for _, record := range gro.Records {
				datas, _ := deaggregation.ExtractRecordDatas(record.Data)
				for _, data := range datas {
					if _, err := c.Store(data); err != nil {
						return err
					}
				}
			}
		}
	}
}

func (c *Consumer) Store(data []byte) (*Consumer, error) {
	record := toRecord(data)
	if recordKey, ok := (*record)[c.recordKey].(string); ok {
		if val, ok := (c.recordMap)[recordKey]; ok {
			c.recordMap[recordKey] = append(val, *record)
		} else {
			c.recordMap[recordKey] = []map[string]interface{}{*record}
		}
		return c, nil
	}
	return c, errors.New("provided recordKey returns non string value")
}

func toRecord(data []byte) *map[string]interface{} {
	record := make(map[string]interface{})
	_ = json.Unmarshal(data, &record)
	return &record
}
