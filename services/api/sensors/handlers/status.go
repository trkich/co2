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

type ResponseBody struct {
	Status	string	`json:"status"`
}

func HandleStatus(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	log.Println("/sensors/{"+vars["uuid"]+"}")

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

	response := ResponseBody{}

	log.Println(sensor.Uuid);

	if sensor.Uuid == "" {
		common.CreateResponse(w, apimessage.NotFound, nil, http.StatusNotFound)
		return;
	}

	if(sensor.StatusExceeded > 2 && sensor.StatusOk < 0){
		response.Status = "ALERT"
	}else if(sensor.StatusOk < 0){
		response.Status = "WARN"
	}else{
		response.Status = "OK"
	}

	common.PlainResponse(w, apimessage.Ok, response)
}


