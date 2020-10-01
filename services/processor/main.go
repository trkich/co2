package main

import (
	"co2/common"
	"co2/common/database"
	"co2/common/database/entity"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	common.LoadEnv()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT")))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"measurement", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")


	mysqlConnector := database.MySQLConnector{}
	dbConn, dbErr := mysqlConnector.GetConnection()
	defer dbConn.Close()
	failOnError(dbErr, "Failed to connect to DB")
	dbConn.AutoMigrate(&entity.Measurement{})
	dbConn.AutoMigrate(&entity.Sensor{})

	dbConn.Exec(`DROP TRIGGER IF EXISTS checkExceeds`)

	dbConn.Exec(`CREATE TRIGGER checkExceeds
						  BEFORE UPDATE
						  ON sensors
						  FOR EACH ROW
						BEGIN
						  IF NEW.status_exceeded<0 THEN
							SET NEW.status_exceeded=0;
						  END IF;
						
						IF NEW.status_exceeded>3 THEN
							SET NEW.status_exceeded=3;
						  END IF;
						
						IF NEW.status_ok>0 THEN
							SET NEW.status_ok=0;
						  END IF;
						END `)


	forever := make(chan bool)

	go func() {
		for d := range msgs {

			newMeasurement := &entity.Measurement{}
			jsonErr := json.Unmarshal([]byte(d.Body), &newMeasurement)
			failOnError(jsonErr, "Unable to parse data")

			dbConn.Create(&newMeasurement)


			if(newMeasurement.Value > 2000){
				dbConn.Exec("INSERT INTO sensors (uuid, status_exceeded, status_ok) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE status_exceeded=status_exceeded+1, status_ok=-3", newMeasurement.Uuid, 1, -3)
			}else{
				dbConn.Exec("INSERT INTO sensors (uuid, status_exceeded, status_ok) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE status_exceeded=status_exceeded-1, status_ok=status_ok+1", newMeasurement.Uuid, 0, 0)
			}

			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}