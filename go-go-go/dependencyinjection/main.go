package main

import "log"

type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}

type MyDataBase struct{}

func (db MyDataBase) GetUser() string {
	return "everyone"
}

func (db MyDataBase) GetAllUsers() []string {
	return []string{}
	// starts an empty string array
}

type FamousDatabase struct{}

func (db FamousDatabase) GetUser() string {
	return "Will Smith"
}

func (db FamousDatabase) GetAllUsers() []string {
	return []string{}

}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}
type MyGreeter struct{}

func (mine MyGreeter) Greet(userName string) {
	log.Printf("Hello %s!! Nice to meet you :)", userName)
}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s!! Nice to meet you", userName)
	// ng = variabila numita ng de tip greeter
}

type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	db := MyDataBase{}
	greeter := MyGreeter{}

	p := Program{
		Db:      db,
		Greeter: greeter,
	}

	p.Execute()
}
