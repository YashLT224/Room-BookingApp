package repository

import (
	"time"

	"github.com/vyash5075/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAlRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)
}
