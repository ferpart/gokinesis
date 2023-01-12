package domain

import "github.com/aws/aws-sdk-go/service/kinesis"

type IKinesis interface {
	GetListShardsOutput() (*kinesis.ListShardsOutput, error)
	GetShardIteratorOutput() (*kinesis.GetShardIteratorOutput, error)
	GetShardIterator() (*string, error)
	GetRecordsOutput() (*kinesis.GetRecordsOutput, error)
	NewStream() error
}
