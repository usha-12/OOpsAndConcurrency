package composition

type engine struct {
	hp int
}

type wheel struct {
	wheelDimension int
}

type Car struct { //has a relationship
	Engine  engine
	Wheel   wheel
	CarName string
}

func NewCar(carName string, hp, wd int) Car {
	return Car{
		CarName: carName,
		Engine: engine{
			hp: hp,
		},
		Wheel: wheel{
			wheelDimension: wd,
		},
	}
}
