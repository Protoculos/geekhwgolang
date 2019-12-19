package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PhoneNumbers struct {
	Person []Person `json:"person,omitempty"`
}

type Person struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type NameValidation struct {
	Name     bool
	LastName bool
}

func LoadPhons(file string) PhoneNumbers {
	var phones PhoneNumbers
	personFile, err := os.Open(file)
	defer personFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(personFile)
	jsonParser.Decode(&phones)
	return phones
}
func AddNewPhones (file string, phones string) {
	personFile, err := os.Create(file)
	defer personFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	personFile.WriteString(phones)
}

func main() {
	var ps PhoneNumbers
	var name, lname, phone string
	var age int
	var exist NameValidation
	var newperson Person
	file := "phone.json"
	ps = LoadPhons(file)
	fmt.Println(ps)
	fmt.Println(ps.Person[0].Name)

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите фамилию: ")
	fmt.Scan(&lname)
	for _, v := range ps.Person {
		if name == v.Name {
			exist.Name = true
		}
		if lname == v.LastName {
			exist.LastName = true
		}
	}
	fmt.Println(exist)
	if exist.Name == true && exist.LastName == true {
		fmt.Println("Данный контакт уже внесен в телефонный справочник.")
	} else {
		fmt.Print("Введите возраст цифрами: ")
		fmt.Scan(&age)
		fmt.Print("Введите телефонный номер: ")
		fmt.Scan(&phone)
		newperson.Name = name
		newperson.LastName = lname
		newperson.Age = age
		newperson.Phone = phone

		ps.Person = append(ps.Person, newperson)
		fmt.Println(ps)

		jsMarPhone, _ := json.Marshal(ps)
		fmt.Println(string(jsMarPhone))
		strPhone := fmt.Sprintln(string(jsMarPhone))
		fmt.Println(strPhone)
		AddNewPhones(file, strPhone)
	}

}
