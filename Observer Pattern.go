package main

import (
	_ "errors"
	"fmt"
)

type Observer interface {
	handleEvent([]string)
}

type Observable interface {
	subscribe(observer Observer) error
	unsubscribe(observer Observer) error
	sendAll()
}

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

type Person struct {
	name string
}

func (j *JobSite) addVacancies(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}

func (j *JobSite) removeVacancies(vacancy string) {
	isFound := false
	for i, value := range j.vacancies {
		if value == vacancy {
			j.vacancies = append(j.vacancies[:i], j.vacancies[i+1:]...)
			isFound = true
		}
	}
	if !isFound {
		fmt.Printf("vacancy %v not found", vacancy)
	}
	j.sendAll()

}

func (j *JobSite) subscribe(ob Observer) error {
	for _, observer := range j.subscribers {
		if ob == observer {
			fmt.Printf("subscriber %v is already exists \n", ob)
			return nil
		}
	}
	j.subscribers = append(j.subscribers, ob)
	return nil
}

func (j *JobSite) unsubscribe(ob Observer) error {
	for i, observer := range j.subscribers {
		//fmt.Println(ob, observer)
		if ob == observer {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
			return nil
		}
	}
	fmt.Printf("subscriber %v not found \n", ob)
	return nil
}

func (j *JobSite) sendAll() {
	for _, observer := range j.subscribers {
		observer.handleEvent(j.vacancies)
	}
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Println("Hello, ", p.name)
	fmt.Println("Changes on our site: ")
	for _, value := range vacancies {
		fmt.Println(value)
	}
}

func main() {
	hh := &JobSite{
		subscribers: nil,
		vacancies:   nil,
	}
	firstPerson := &Person{"Alfred"}
	secondPerson := &Person{"Anton"}

	hh.subscribe(firstPerson)
	hh.subscribe(secondPerson)
	hh.subscribe(secondPerson)

	hh.sendAll()
	fmt.Println("----------------------------------Update----------------------------------")
	hh.addVacancies("GameDeveloper (C#, Unity)")
	fmt.Println("----------------------------------Update----------------------------------")
	hh.unsubscribe(secondPerson)
	hh.unsubscribe(secondPerson)
	hh.addVacancies("UI/UX Designer")
	hh.addVacancies("Developer Java")
	hh.removeVacancies("UI/dd Designer")
}
