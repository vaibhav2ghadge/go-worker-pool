package jsonrepo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"trivago/domain"
)

// Store data in sqlite db and return stored id and error
func (jsons JSONService) Store(hotelInfo domain.JsonHotel) error {
	file, err := json.MarshalIndent(hotelInfo, "", "")
	if err != nil {
		return err
	}
	ioutil.WriteFile(jsons.FilePath, file, os.ModePerm)
	return nil
}
