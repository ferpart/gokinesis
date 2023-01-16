# GoKinesis

This repository is used to provide utils when running kinesis locally.

To determine the application flow, GoKinesis has the `operation-type` or `-o` flag. 
This flag defaults to `starter` but it supports any of the following values:
* `starter`: Used to run the [Starter](#starter) application.
* `consumer`: Used to run the [Consumer](#consumer) application

### Example

#### Starter application
```shell
./gokinesis -o starter
```

#### Consumer application
```shell
./gokinesis -o consumer
```

## Starter
When running the application with the `starter` param, a new stream will be created. 
The following flags are supported:


**Stream-name** `-s` or `--stream-name` flags are used for determining the name for
the new kinesis stream. The application defaults to the `default` stream name.

```shell
./gokinesis -o starter -s STREAM_NAME
```

**Hostname** `-n` or `--hostname` flags are used for determining the host of the 
kinesis stream. The application defaults to using `http://localhost:4568` as the 
default kinesis hostname.

```bash
./gokinesis -o starter -n HOSTNAME:PORT
```

The application will close with a message when the stream has been created.

## Consumer

When running the application with th `consumer` param, a kinesis stream will be
consumed. These contents will be saved, and parsed into a CSV stored in a new 
`tracked` directory with the `YYYY-MM-DDThh:mm:ss` format. The following flags
are supported for this application:

**Stream-name** `-s` or `--stream-name` flags are used for determining the name for
the new kinesis stream. The application defaults to the `default` stream name.

```shell
./gokinesis -o consumer -s STREAM_NAME
```

**Hostname** `-n` or `--hostname` flags are used for determining the host of the
kinesis stream. The application defaults to using `http://localhost:4568` as the
default kinesis hostname.

```bash
./gokinesis -o consumer -n HOSTNAME:PORT
```
