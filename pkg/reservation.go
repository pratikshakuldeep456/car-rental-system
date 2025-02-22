package pkg

import "time"

type Reservation struct {
	ID         int
	Car        Car
	Customer   Customer
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float64
}

func NewReservation() *Reservation {
	return &Reservation{}
}

func (r *Reservation) Price() {
	days := r.EndDate.Sub(r.StartDate).Hours() / 24
	r.TotalPrice = days * r.Car.RentalPrice
	// return r.TotalPrice

}
