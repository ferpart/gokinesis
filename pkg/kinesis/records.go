package kinesis

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func (k *Kinesis) GetListShardsOutput() (*kinesis.ListShardsOutput, error) {
	lsi := &kinesis.ListShardsInput{
		StreamName: k.streamName,
	}
	return k.client.ListShards(lsi)
}

func (k *Kinesis) GetShardIteratorOutput() (*kinesis.GetShardIteratorOutput, error) {
	lso, err := k.GetListShardsOutput()
	if err != nil {
		return nil, err
	}

	if len(lso.Shards) != 1 {
		return nil, errors.New("should have exactly 1 shard")
	}

	sii := &kinesis.GetShardIteratorInput{
		StreamName:        k.streamName,
		ShardId:           lso.Shards[0].ShardId,
		ShardIteratorType: aws.String("AT_TIMESTAMP"),
		Timestamp:         &time.Time{},
	}

	return k.client.GetShardIterator(sii)
}

func (k *Kinesis) GetShardIterator() (*string, error) {
	if k.shardIterator != nil && *k.shardIterator != "" {
		return k.shardIterator, nil
	}
	sho, err := k.GetShardIteratorOutput()
	if err != nil {
		return nil, err
	}

	k.shardIterator = sho.ShardIterator
	return sho.ShardIterator, nil
}

func (k *Kinesis) GetRecordsOutput() (*kinesis.GetRecordsOutput, error) {
	si, err := k.GetShardIterator()
	if err != nil {
		return nil, err
	}
	gri := &kinesis.GetRecordsInput{
		ShardIterator: si,
		Limit:         aws.Int64(100),
	}

	gro, _ := k.client.GetRecords(gri)

	k.shardIterator = gro.NextShardIterator

	return gro, nil
}
