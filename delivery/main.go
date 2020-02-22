package main

import (
	"flag"
	"log"
	"trivago/handler/csvhandler"
	"trivago/repository/dbrepo"
	"trivago/repository/jsonrepo"
	"trivago/utils/sqliteutils"
)

var (
	sqliteDBPath       = flag.String("sqliteDBPath", "hotel.db", "the db file for sqlite to store data")
	inputCSVFilePath   = flag.String("inputFilePath", "hotels.csv", "input file that we are going to process")
	jsonOutputFilePath = flag.String("jsonFilePath", "jsondump.json", "store the data in json file")
	workers            = flag.Int("workers", 5, "no of worker who write into sqlite")
)

func main() {
	flag.Parse()
	log.Println(*sqliteDBPath)
	sqliteAuth := sqliteutils.SqliteAuth{DBPath: *sqliteDBPath}
	db := sqliteutils.NewSqliteConnection(sqliteAuth)
	dbService := dbrepo.DBService{SqliteDBConnection: db}
	jsonService := &jsonrepo.JSONService{FilePath: *jsonOutputFilePath}
	workerPool := csvhandler.NewWorkerPool(*workers)
	CSVHandler := csvhandler.NewCSVHandler(*inputCSVFilePath, dbService, jsonService, workerPool)
	log.Println("Csv processing started")
	CSVHandler.ProcessCsvFile()
	log.Println("Csv processing end")

}
