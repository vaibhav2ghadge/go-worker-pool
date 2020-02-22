package jsonrepo

import (
	"trivago/domain"
)

type Writer interface {
	Store(hotelInfo domain.JsonHotel) error
}
