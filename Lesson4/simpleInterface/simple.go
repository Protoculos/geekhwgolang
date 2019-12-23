//1. Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.
package main

import (
	"fmt"
	"time"
)

//Person struct
type Person struct {
	Name  string `json:"name,omitempty"`
	Lname string `json:"lname,omitempty"`
	Age   int    `json:"age,omitempty"`
	Call
}

//Company struct
type Company struct {
	Name     string  `json:"name,omitempty"`
	Contact  *Person `json:"contact,omitempty"`
	Location string  `json:"location,omitempty"`
	Call
}

//Call struct
type Call struct {
	Rings      map[time.Time]bool `json:"rings,omitempty"`
	CountTrue  int                `json:"count_true,omitempty"`
	CountFalse int                `json:"count_false,omitempty"`
}

//AddRings method
func (p *Person) AddRings(t time.Time, b bool) error {
	if b == true {
		p.Rings[t] = true
		p.CountTrue++
	} else {
		p.Rings[t] = false
		p.CountFalse++
	}
	return nil
}

//AddRings method
func (c *Company) AddRings(t time.Time, b bool) error {
	if b == true {
		c.Rings[t] = true
		c.CountTrue++
	} else {
		c.Rings[t] = false
		c.CountFalse++
	}
	return nil
}

//ClientsRings interface
type ClientsRings interface {
	AddRings(t time.Time, b bool) error
}

//CallsClients function
func CallsClients(c ClientsRings, s string) {
	var b bool
	now := time.Now()
	if s == "да" {
		b = true
	} else if s == "нет" {
		b = false
	} else {
		fmt.Println("Введите да или нет о результате звонка.")
	}
	err := c.AddRings(now, b)
	if err != nil {
		fmt.Println("Так вы дозвонились до клиента или нет?")
	}
}

func main() {
	m := make(map[time.Time]bool)
	pers1 := &Person{
		Name:  "Alex",
		Lname: "Petrov",
		Age:   27,
		Call: Call{
			Rings: m,
		},
	}
	pers2 := &Person{
		Name:  "Simon",
		Lname: "Jonson",
		Age:   54,
		Call: Call{
			Rings: m,
		},
	}
	comp1 := &Company{
		Name:     "Company Co",
		Location: "Russia",
		Contact:  pers2,
		Call: Call{
			Rings: m,
		},
	}
	CallsClients(pers1, "да")
	fmt.Println(pers2)
	CallsClients(comp1, "нет")
	fmt.Println(comp1)
}
