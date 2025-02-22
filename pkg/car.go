package pkg

// enum
type Status string

const (
	Available Status = "Available"
	Booked    Status = "Booked"
)

type Car struct {
	Model        string
	ManYear      int
	LicensePlate string
	RentalPrice  float64
	Status       Status
}

func AddCar() *Car {
	return &Car{}
}
