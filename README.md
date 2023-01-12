# GOKinesis

This repository is used to provide utils when running kinesis locally. This repo
provides the following features:

## Starter
When running the [starter.go](https://github.com/ferpart/gokinesis/blob/091fb8c59bfb821dac43540180c5f564a0e28c55/cmd/starter/starter.go)
application, a new stream will be created with the provided **stream name**.

The `-s` or `--stream-name` flags are required.

```bash
go run starter.go -s STREAM_NAME
```

The application defaults to using `http://localhost:4568` as the default kinesis
hostname. This can be changed with the `-h` or `--hostname` flags.

```bash
go run starter.go -s stream_a -h HOSTNAME:PORT
```

The application will close with a message when the stream has been created.

## Consumer

When running the [consumer.go](https://github.com/ferpart/gokinesis/blob/091fb8c59bfb821dac43540180c5f564a0e28c55/cmd/consumer/consumer.go)
application, the stream with the provided **stream name** will be read, and its
contents will be saved, and parsed into a CSV stored in a new `tracked/` directory 
with the `YYYY-MM-DDThh:mm:ss` format.

The `-s` or `--stream-name` flags are required.

```bash
go run consumer.go -s STREAM_NAME
```

The application defaults to using `http://localhost:4568` as the default kinesis
hostname. This can be changed with the `-h` or `--hostname` flags.

```bash
go run starter.go -s stream_a -h HOSTNAME:PORT
```
