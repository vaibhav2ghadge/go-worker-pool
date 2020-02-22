package dbrepo

import (
	"trivago/domain"
)

type Writer interface {
	Store(hotelInfo domain.Hotel) error
}
