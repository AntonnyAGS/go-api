package functions

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

import . "github.com/ahmetb/go-linq"
import . "github.com/nleeper/goment"

import "github.com/AntonnyAgs/go-api/domain"


func CheckConflicts(agendas domain.Agendas, agendasToPublish domain.Agendas) domain.Agendas {
	_agendas := From(agendas).GroupBy(
		func(i interface{}) interface{} { return i.(domain.Agenda).DoctorId },
		func(i interface{}) interface{} { return i }).Results()

	result := make(map[string][]interface{})

	From(_agendas).SelectT(
		func(i Group) KeyValue { return KeyValue{ Key: i.Key, Value: i.Group } },
	).ToMap(&result)

	conflicts := make(domain.Agendas, 0) 

	for _, agenda := range agendasToPublish {
		allocations := result[agenda.DoctorId]

		for _, allocation := range allocations {
			isSameDate := CompairDates(agenda, allocation.(domain.Agenda))

			isSameAgenda := allocation.(domain.Agenda).Id == agenda.Id

			isSameDoctor := allocation.(domain.Agenda).DoctorId == agenda.DoctorId

			if(!isSameAgenda && isSameDate && isSameDoctor) {
				conflicts = append(conflicts, allocation.(domain.Agenda))
			}
		} 
	}

	return conflicts
}

func CompairDates(publishAllocation domain.Agenda, agendaAllocation domain.Agenda) bool {
	a := FormatDate(agendaAllocation.Date.Begin).IsBefore(publishAllocation.Date.End)
	b := FormatDate(publishAllocation.Date.Begin).IsBefore(agendaAllocation.Date.End)

	return a && b
}

func FormatDate(date string) *Goment {
	result, error := New(date)

	if (error != nil) {
		fmt.Print("error", error)
	}

	return result
}

func ReadJson(fileUrl string) domain.Agendas {
	data, err := ioutil.ReadFile(fileUrl)

	if err != nil {
		fmt.Print("error", err)
	}
	
	var obj domain.Agendas

	err = json.Unmarshal(data, &obj)

	if err != nil {
		fmt.Println("error", err)
	}

	return obj
}