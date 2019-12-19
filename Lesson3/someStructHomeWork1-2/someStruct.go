package main

import "fmt"

type car struct {
	Model    string `json:"model,omitempty"`
	Year     string `json:"year,omitempty"`
	TrunkVol int    `json:"trunk_vol,omitempty"` //Объем багажника
	BodyVol  int    `json:"body_vol,omitempty"`  //Объем кузова
	Engine   bool   `json:"engine,omitempty"`    //Запущен ли двигатель
	Unmanned bool   `json:"unmanned,omitempty"`  //Есть ли автопилот
}

type truck struct {
	Model        string `json:"model,omitempty"`
	Year         string `json:"year,omitempty"`
	Engine       bool   `json:"engine,omitempty"`
	Unmanned     bool   `json:"unmanned,omitempty"`
	Tonnage      int    `json:"tonnage,omitempty"`       //Грузоподъемность
	TotalMileage int    `json:"total_mileage,omitempty"` //Сколько миль пройдет на одной заправке
}

func main() {
	car1 := car{
		Model:    "opel",
		Year:     "1990",
		TrunkVol: 35,
		BodyVol:  60,
		Engine:   true,
		Unmanned: false,
	}
	car1.Engine = false

	car2 := car{
		Model:    "tesla",
		Year:     "2020",
		TrunkVol: 57,
		BodyVol:  57,
		Engine:   false,
		Unmanned: true,
	}
	car2.Unmanned = false

	truck1 := truck{
		Model:        "kamas",
		Year:         "1995",
		Engine:       true,
		Unmanned:     false,
		Tonnage:      50,
		TotalMileage: 120,
	}
	truck1.Tonnage = 55

	truck2 := truck{
		Model:        "BMV",
		Year:         "2020",
		Engine:       false,
		Unmanned:     true,
		Tonnage:      87,
		TotalMileage: 200,
	}
	truck2.TotalMileage = 250

	var cars []car = []car{car1, car2}
	var trucks []truck = []truck{truck1, truck2}
	for _, v := range cars {
		fmt.Println("Car: ", v.Model, v.Year, v.Engine)
	}
	for _, v := range trucks {
		fmt.Println("Truck: ", v.Model, v.Year, v.Tonnage, v.Unmanned)
	}

}
