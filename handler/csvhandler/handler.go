package csvhandler

import (
	"encoding/csv"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"trivago/domain"
	"trivago/repository/dbrepo"
	"trivago/repository/jsonrepo"
	"unicode/utf8"
)

//CSVHandler it will handle all csv operation
type CSVHandler struct {
	CSVFilePath        string
	SqliteStoreService dbrepo.DBService
	JSONStoreService   *jsonrepo.JSONService
	WorkerPool         *WorkerPool
}
type WorkerChannel chan chan domain.Hotel
type WorkerPool struct {
	WorkerCount int
	WorkerChannel
}

func NewWorkerPool(noOfWorkers int) *WorkerPool {
	return &WorkerPool{
		WorkerCount:   noOfWorkers,
		WorkerChannel: make(chan chan domain.Hotel)}
}

type Worker struct {
	ID      int
	Channel chan domain.Hotel
	WorkerChannel
}

func (c *CSVHandler) Start(w *Worker) {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case job := <-w.Channel: // worker has received job
				c.SqliteStoreService.Store(job)
			}
		}
	}()
}

func (c *CSVHandler) StartWorkers(workerCount int, wc WorkerChannel) {
	//x := make(chan domain.Hotel)
	for i := 0; i < workerCount; i++ {
		worker := &Worker{Channel: make(chan domain.Hotel), WorkerChannel: wc}
		c.Start(worker)
	}
}

//NewCSVHandler ...
func NewCSVHandler(filePath string, sqliteStoreService dbrepo.DBService, jsonStoreServce *jsonrepo.JSONService, workerPool *WorkerPool) *CSVHandler {
	return &CSVHandler{CSVFilePath: filePath,
		SqliteStoreService: sqliteStoreService,
		JSONStoreService:   jsonStoreServce,
		WorkerPool:         workerPool,
	}
}

//ProcessCsvFile it will read from csv file and send it to controler for stroing data
func (c *CSVHandler) ProcessCsvFile() {
	var Hotels domain.JsonHotel
	csvfile, err := os.Open(c.CSVFilePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	c.StartWorkers(c.WorkerPool.WorkerCount, c.WorkerPool.WorkerChannel)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		rating, err := strconv.Atoi(record[2])
		//if rating is not number
		if err != nil {
			continue
		}
		hotel := domain.Hotel{record[0], record[1], rating, record[3], record[4], record[5]}
		if validateHotelName([]byte(hotel.Name)) {
			if validateHotelURL(hotel.URL) {
				if hotel.Rating >= 0 && hotel.Rating <= 5 {
					//store into sqlite
					// err := c.SqliteStoreService.Store(hotel)
					// if err != nil {
					// 	log.Println(err)
					// }
					//store into json file

					Hotels.Hotels = append(Hotels.Hotels, hotel)
					//c.JSONStoreService.JSONService.Store(hotel)
					worker := <-c.WorkerPool.WorkerChannel // wait for available channel
					worker <- hotel
				}
			}
		}
	}
	c.JSONStoreService.Store(Hotels)

}

// validate url https://golang.org/pkg/net/url/#ParseRequestURI
// invalidate url "http//google.com","google.com","/foo/bar"
func validateHotelURL(urla string) bool {
	_, err := url.ParseRequestURI(urla)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// check given name is valid utf8
func validateHotelName(name []byte) bool {
	return utf8.Valid(name)
}
