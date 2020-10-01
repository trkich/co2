package handlers

import (
	"co2/common"
	"co2/common/apimessage"
	"co2/common/database/entity"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

func failOnError(w http.ResponseWriter, err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		common.CreateResponse(w, apimessage.DatabaseError, "rrr", http.StatusInternalServerError)
		return
	}
}

func HandleMeasurements(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	log.Println("/sensors/{"+vars["uuid"]+"}/measurements")

	var userRequest entity.Measurement
	jsonErr := json.NewDecoder(r.Body).Decode(&userRequest)

	if jsonErr != nil {
		common.CreateResponse(w, apimessage.JsonParseError, nil, http.StatusBadRequest)
		log.Println("Bad Request")
		return
	}

	if jsonErr != nil {
		log.Println("Bad Request")
		common.CreateResponse(w, apimessage.JsonParseError, nil, http.StatusBadRequest)
		return
	}

	v := validator.New()
	validateErr := v.Struct(userRequest)

	if validateErr != nil {
		log.Println("Bad Request 2")
		common.CreateResponse(w, apimessage.JsonParseError, nil, http.StatusBadRequest)
		return
	}

	userRequest.Uuid = vars["uuid"]

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(w, err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(w, err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"measurement", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(w, err, "Failed to declare a queue")

	body, _ := json.Marshal(userRequest)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(w, err, "Failed to publish a message")

	common.CreateOkResponse(w, userRequest)
}


