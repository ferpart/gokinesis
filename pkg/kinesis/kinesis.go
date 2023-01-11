package kinesis

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type Kinesis struct {
	client        *kinesis.Kinesis
	streamName    *string
	shardIterator *string
}

func New(
	endpoint string,
	region string,
	streamName string,
) *Kinesis {
	cfg := &aws.Config{
		Endpoint:    &endpoint,
		Region:      &region,
		Credentials: credentials.NewStaticCredentials("x", "y", ""),
	}

	ss, _ := session.NewSession(cfg)

	return &Kinesis{
		client:     kinesis.New(ss, cfg),
		streamName: &streamName,
	}
}
