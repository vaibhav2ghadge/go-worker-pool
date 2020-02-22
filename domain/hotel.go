package domain

type Hotel struct {
	Name          string
	Address       string
	Rating        int
	ContactPerson string
	PhoneNumber   string
	URL           string
}
type JsonHotel struct {
	Hotels []Hotel
}
