# CO2 measurements collector
Service is intended to be used as a CO2 collector for collecting measurement from sensors via simple API and showcase of microservices implementation with Go and RabbitMQ.
There is a lot of space for improvements (Check TODO) :) 

Service have 3 microservices:
- API - Exposing API endpoints
- Processors - Processing measurements and storing them to MySQL
- Aggregator - Running in background and calculating Max and Avg for each sensor

In this way Service can be scaled both horizontally and vertically as show on following diagram:

```bash
                             +----------------------------------------------------+
                             |                                                    |
                             |                                                    |
                         +---v---+                          +-----------+         |        +------------+
                         |       |                          |           |         |        |            |
                     +--->  API  +--+                   +---> Processor +-+       |      +-> Aggregator |
+---------------+    |   |       |  |  +-------------+  |   |           | |  +----v----+ | |            |
|               |    |   +-------+  +-->             |  |   +-----------+ +-->         <-+ +------------+
| API Gateway   <----+                 |  RabbitMQ   +-->                    |  MySQL  |
| Load Balancer <----+   +-------+  +-->             |  |   +-----------+ +-->         <-+ +------------+
|               |    |   |       |  |  +-------------+  |   |           | |  +----^----+ | |            |
+---------------+    +--->  API  +--+                   +---> Processor +-+       |      +-> Aggregator |
                         |       |                          |           |         |        |            |
                         +---^---+                          +-----------+         |        +------------+
                             |                                                    |
                             |                                                    |
                             +----------------------------------------------------+
```

#### Service have following APIs:

1. Push sensor measurements:
```bash
POST /api/v1/sensors/{uuid}/measurements
{
"co2" : 2000,
"time" : "2019-02-01T18:55:47+00:00" 
}
```

2. Get sensor max and average measurements:
```bash
GET /api/v1/sensors/{uuid}/metrics
Response: {
    "maxLast30Days" : 1200,
    "avgLast30Days" : 900
}
```

2. Obtain sensor status:
```bash
POST /api/v1/sensors/{uuid}/measurements
Response: {
    "status" : "OK" // Possible status OK,WARN,ALERT
}
```

## Build and run in local environment

### GoLang

To build project you need to have GoLang installed:
https://golang.org/doc/install

### MySQL and RabbitMq
Before starting project in local environment make sure that you have MySQL server and RabbitMQ services up and running.

If you do not have it installed, you can follow official instructions to install it or to run it in Docker:

https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/
https://www.rabbitmq.com/#getstarted

Make sure to change configuration in .env file to match your local MySQL and RabbitMQ settings:
```bash
API_PORT=8080
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root
DB_NAME=co2db
DB_PORT=3306
RABBITMQ_HOST=localhost
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_PORT=5672
AGGREGATION_AVG_INTERVAL_TIME=2
AGGREGATION_MAX_INTERVAL_TIME=1
```


### Build
To build project run:
```bash
make clean build
```

### Start

To start API service run:
```bash
make run-api
```

To start Processor service run:
```bash
make run-processor
```

To start Aggregator service run:
```bash
make run-aggregator
```


## Testing
To run all test execute following command:
```bash
make test
```

## TODO
- Improve error handling 
  - MySQL connection
  - RabbitMQ connection
  - General validation
- Increase test coverage
- Improve logging
- Dockerization
- Add support for Minikube
- Create new microservice for managing status
- Create new microservice for handling MySQL queries (Move communication over RabbitMQ)
