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

func main() {
	s := Student {
		Name: "Yamada",
		Number: 999,
		Grade: 5,
	}
	t := Teacher {
		Name: "Tsubomi",
	}

	ctxStu := sendEmailOfSudent(s)
	fmt.Println(ctxStu)
	ctxTea := sendEmailOfTeacher(t)
	fmt.Println(ctxTea)
}

func (s Student) getEmail() string {
	return s.Name + "@student.ed.jp"
}

func (t Teacher) getEmail() string {
	return t.Name + "@teacher.ed.jp"
}

func sendEmailOfSudent(s Student) (context string) {
	from := s.getEmail()
	context = `
from : ` + from + `
this is a test mail
regards.
`
	return context
}

func sendEmailOfTeacher(t Teacher) (context string) {
	from := t.getEmail()
	context = `
from : ` + from + `
this is a test mail
regards.
`
	return context
}
