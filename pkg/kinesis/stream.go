package kinesis

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/ferpart/gokinesis/pkg/wait"
)

func (k *Kinesis) NewStream() error {
	csi := &kinesis.CreateStreamInput{
		StreamName: k.streamName,
		ShardCount: aws.Int64(1),
	}

	err := wait.For(func() bool {
		_, err := k.client.CreateStream(csi)
		return err == nil
	})

	if err != nil {
		return err
	}

	dsi := &kinesis.DescribeStreamInput{
		StreamName: k.streamName,
	}

	err = wait.For(func() bool {
		describe, _ := k.client.DescribeStream(dsi)
		return describe.StreamDescription.StreamStatus != nil &&
			*describe.StreamDescription.StreamStatus == "ACTIVE"
	})
	if err != nil {
		return err
	}

	return nil
}
