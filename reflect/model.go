package reflect

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func (p Person) GetName() string {
	return fmt.Sprintf("person: %s", p.Name)
}

type Student struct {
	Person
	Class string
	Score string
}

func (s Student) GetName() string {
	return fmt.Sprintf("student: %s", s.Name)
}

type Teacher struct {
	Person
	Subject string
}

func (t Teacher) GetName() string {
	return fmt.Sprintf("teacher: %s", t.Name)
}
