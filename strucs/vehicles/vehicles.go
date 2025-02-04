package vehicles

import "fmt"


type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle = "TRUCK"
)


func New (v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type %s", v)
	}
}

type Car struct {
	Time int 
}


func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

// MotorCycle represents a motorcycle vehicle
type Motorcycle struct {
	Time int
}

// Distance returns the distance traveled by the motorcycle in kilometers
func (m *Motorcycle) Distance() float64 {
	return 50 * (float64(m.Time) / 12)
}

// Truck represents a truck vehicle
type Truck struct {
	Time int
}

// Distance returns the distance traveled by the truck in kilometers
func (t *Truck) Distance() float64 {
	return 30 * (float64(t.Time) / 1000)
}
