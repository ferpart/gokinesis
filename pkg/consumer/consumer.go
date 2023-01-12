package consumer

import (
	"encoding/json"
	"github.com/ferpart/gokinesis/domain"
	"github.com/ferpart/gokinesis/pkg/save2csv"
	"github.com/mjneil/kinesis-producer/deaggregation"
)

type Consumer struct {
	kinesis domain.IKinesis
}

func New(kinesis domain.IKinesis) *Consumer {
	return &Consumer{
		kinesis: kinesis,
	}
}

func (c *Consumer) Consume(stopConsume <-chan bool) error {
	for {
		select {
		case <-stopConsume:
			return nil
		default:
			consumed, err := c.consume()
			if err != nil {
				return err
			}
			if err = save2csv.Save2CSV(consumed); err != nil {
				return err
			}
		}
	}
}

func (c *Consumer) consume() ([]map[string]interface{}, error) {
	gro, err := c.kinesis.GetRecordsOutput()
	if err != nil {
		return nil, err
	}

	var parsedRecords []map[string]interface{}
	for _, record := range gro.Records {
		dData, _ := deaggregation.ExtractRecordDatas(record.Data)
		for _, data := range dData {
			parsedRecords = append(parsedRecords, dataParser(data))
		}
	}

	return parsedRecords, nil
}

func dataParser(data []byte) map[string]interface{} {
	record := make(map[string]interface{})
	_ = json.Unmarshal(data, &record)
	return record
}
