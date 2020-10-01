package handlers

import (
	"co2/common"
	"co2/common/apimessage"
	"co2/common/database"
	"co2/common/database/entity"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type MetricsBody struct {
	MaxLast30Days	int	`json:"maxLast30Days"`
	AvgLast30Days	int	`json:"avgLast30Days"`
}

func HandleMetrics(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	log.Println("/sensors/{"+vars["uuid"]+"}/metrics")

	mysqlConnector := database.MySQLConnector{}
	dbConn, dbErr := mysqlConnector.GetConnection()
	defer dbConn.Close()
	if dbErr != nil {
		fmt.Print(dbErr)
		common.CreateResponse(w, apimessage.DatabaseError, nil, http.StatusInternalServerError);
	}
	dbConn.AutoMigrate(&entity.Sensor{})

	var sensor entity.Sensor
	dbConn.First(&sensor, "uuid = ?", vars["uuid"])

	if sensor.Uuid == "" {
		common.CreateResponse(w, apimessage.NotFound, nil, http.StatusNotFound)
		return;
	}

	common.PlainResponse(w, apimessage.Ok, MetricsBody{MaxLast30Days:sensor.MaxLast30Days, AvgLast30Days:sensor.AvgLast30Days})
}


