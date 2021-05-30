package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	UNKNOWN status = iota // 0
	TODO                  // 1
	DONE                  // 2
)

type status int

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
}

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// 직렬화
// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}
func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		1,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

// 역직렬화
// Buy Milk
// 2
// 2015-08-16 15:43:00 +0000 UTC
func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk", "Status":2, "Deadline":"2015-08-16T15:43:00Z"}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
}

func main() {
	fmt.Println("<<< marshall >>>")
	Example_marshalJSON()
	fmt.Println("<<<unmarshall>>>")
	Example_unmarshalJSON()
}
