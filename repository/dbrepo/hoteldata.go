package dbrepo

import (
	"log"
	"trivago/domain"
)

// Store data in sqlite db and return stored id and error
func (db DBService) Store(hotelInfo domain.Hotel) error {
	stmt, err := db.SqliteDBConnection.Prepare("INSERT INTO hotel(name, address, stars, contact, phone, url) values(?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(hotelInfo.Name, hotelInfo.Address, hotelInfo.Rating, hotelInfo.ContactPerson, hotelInfo.PhoneNumber, hotelInfo.URL)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
