package pkg

import "fmt"

type Subscriber struct {
	Name string
}

func (s *Subscriber) Update(pubName string) {
	fmt.Println("Subscriber", s.Name, "received", pubName)
}

func (s *Subscriber) GetName() string {
	return s.Name
}
