package main

import "fmt"

type Student struct {
	Name string
	Number int
	Grade int
}

type Teacher struct {
	Name string
}

type Person interface {
	getEmail() string
}

func main() {
	var s, t Person
	s = Student {
		Name: "Yamada",
		Number: 999,
		Grade: 5,
	}
	t = Teacher {
		Name: "Tsubomi",
	}

	ctxStu := sendEmail(s)
	fmt.Println(ctxStu)
	ctxTea := sendEmail(t)
	fmt.Println(ctxTea)
}

func (s Student) getEmail() string {
	return s.Name + "@student.ed.jp"
}

func (t Teacher) getEmail() string {
	return t.Name + "@teacher.ed.jp"
}

func sendEmail(p Person) (context string) {
	from := p.getEmail()
	context = `
from : ` + from + `
this is a test mail
regards.
`
	return context
}
