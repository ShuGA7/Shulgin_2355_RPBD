package main

import (
	"fmt"

	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service/store"
)

func main() {
	conn := "postgres://shulgin:123123123@95.217.232.188:7777/shulgin"
	s := store.NewStore(conn)
	fmt.Println(s.GetPeopleByID("400"))
	fmt.Println(s.ListPeople())
}
