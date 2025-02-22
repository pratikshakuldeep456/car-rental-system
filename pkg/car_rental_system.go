package pkg

// car should have details such as make, model,
//  year, license plate number, and rental price per day.

//The system should handle payment processing for reservations.

type CRS struct {
	Cars         map[string]*Car
	Reservations map[int]*Reservation
	Payment      PaymentProcessor
}

var CRSInstance *CRS

func GetCRSInstance() *CRS {

	if CRSInstance == nil {
		CRSInstance = &CRS{
			Cars:         make(map[string]*Car),
			Reservations: make(map[int]*Reservation),
		}
	}
	return CRSInstance
}

func (c *CRS) AddCar(car *Car) {
	c.Cars[car.LicensePlate] = car
}

func (c *CRS) RemoveCar(LicensePlate string) {

	delete(c.Cars, LicensePlate)

}

func (c *CRS) AvailableCar() {

}
