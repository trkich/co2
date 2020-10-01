package main

import (
	"co2/common"
	"co2/common/database"
	"co2/common/database/entity"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

func runEvery(n time.Duration, f func(db *gorm.DB), db *gorm.DB) {
	for _ = range time.Tick(n * time.Second) {
		f(db)
	}
}

func aggregateAvg(db *gorm.DB) {
	db.Exec(`update sensors s join
       (select uuid, MAX(value) as maxco2
        from measurements m  where m.time > (NOW() - INTERVAL 1 MONTH)
        group by uuid 
       ) r
       on s.uuid = r.uuid
    set s.max_last30_days = maxco2;`)

	log.Println("Averages updated!")
}

func aggregateMax(db *gorm.DB) {
	db.Exec(`update sensors s join
       (select uuid, avg(value) as avg
        from measurements m  where m.time > (NOW() - INTERVAL 1 MONTH)
        group by uuid 
       ) r
       on s.uuid = r.uuid
    set s.avg_last30_days = r.avg;`)

	log.Println("Max updated!")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	common.LoadEnv()

	log.Println("Starting aggregation service")

	mysqlConnector := database.MySQLConnector{}
	dbConn, dbErr := mysqlConnector.GetConnection()
	defer dbConn.Close()
	failOnError(dbErr, "Failed to connect to DB")
	dbConn.AutoMigrate(&entity.Measurement{})
	dbConn.AutoMigrate(&entity.Sensor{})

	avgAggTime, err := strconv.Atoi(os.Getenv("AGGREGATION_AVG_INTERVAL_TIME"))
	failOnError(err, "Can't cast AGGREGATION_AVG_INTERVAL_TIME to int")

	maxAggTime, err := strconv.Atoi(os.Getenv("AGGREGATION_MAX_INTERVAL_TIME"))
	failOnError(err, "Can't cast AGGREGATION_MAX_INTERVAL_TIME to int")


	log.Printf("Average aggregation time: %d seconds", avgAggTime)
	log.Printf("Max aggregation time: %d seconds", maxAggTime)

	go runEvery(time.Duration(int32(maxAggTime)), aggregateAvg, dbConn)
	go runEvery(time.Duration(int32(maxAggTime)), aggregateMax, dbConn)

	select {}
}