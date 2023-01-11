package consume

import (
	"encoding/json"

	"github.com/fpartidabc/gokinesis/pkg/kinesis"
	"github.com/mjneil/kinesis-producer/deaggregation"
)

func Consume(k *kinesis.Kinesis, stopConsume <-chan bool, storeChan chan *CommonMap) {
	commonStore := make(CommonMap, 0)
	for {
		select {
		case <-stopConsume:
			storeChan <- &commonStore
			break
		default:
			gro, _ := k.GetRecordsOutput()
			for _, record := range gro.Records {
				datas, _ := deaggregation.ExtractRecordDatas(record.Data)
				for _, data := range datas {
					commonStore.Store(toCommon(data))
				}
			}
		}
	}
}

type CommonMap map[string][]map[string]interface{}

func toCommon(data []byte) *map[string]interface{} {
	commonLog := make(map[string]interface{})
	_ = json.Unmarshal(data, &commonLog)
	return &commonLog
}

func (c *CommonMap) Store(common *map[string]interface{}) *CommonMap {
	sessionID := (*common)["session_id"].(string)
	if val, ok := (*c)[sessionID]; ok {
		(*c)[sessionID] = append(val, *common)
	} else {
		(*c)[sessionID] = []map[string]interface{}{*common}
	}

	return c
}
