# Message Sending System

This project automatically sends 2 messages which have not yet been sent in the database every 2 minutes, which can be configured. It consists of two parts: API and a scheduler.

Upon deployment scheduler starts running. The API is designed to manipulate the scheduler and get processed messages.

## Configuration

Configuration files are put in `.configs` folder. Please add your configs according to your environment or modify the existing one for development usage.

| Field    | Description |
| -------- | ------- |
| scheduler.period_secs     |   This field is used to set frequency of the scheduler in seconds. |
| couchbase.host            |   This field is used to set host name of your Couchbase instance.  |
| couchbase.username        |   This field is used to set username of your Couchbase instance.   |
| couchbase.password        |   This field is used to set password of your Couchbase instance.   |
| couchbase.wait_until_ready_secs |  This field is used to set seconds to wait until your Couchbase instance is ready.   |
| webhook.url               |  This field is used to set webhook.site url. |
| webhook.api_key           |  This field is used to set key in x-ins-auth-key header. |
| redis.host           |  This field is used to set host name of your Redis instance. |
| redis.password           |  This field is used to set password of your Redis instance. |
| redis.db           |  This field is used to set db of your Redis instance. |
| redis.ttl_secs           |  This field is used to ttl seconds in Redis instance. |

## How to run

1. Assuming you have Docker installed in your system run the following command in the terminal:

```
docker-compose up --build
```

`docker-compose.yml` file is used to set up the whole environment.

2. Go to <http://localhost:8091> to set up your Couchbase cluster.
3. Add a bucket with name `messages` to your cluster.
4. Add documents to the _default collection to test.

NOTE: An example json document can be found in `/samples` folder.

5. If you get a connection timeout, run `docker start go-app` in another terminal to start the application.

## How to test

```
make tests
```

## To generate swagger docs

Swagger docs are located in main path of your application. Go to <http://localhost:8080>.

```
swag init -g cmd/main.go -o docs --parseDependency --parseInternal
```
