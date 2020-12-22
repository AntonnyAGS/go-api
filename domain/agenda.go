package domain

type Agenda struct {
	Id string `json:"id"`
	Date struct {
		Day string `json:"day"`
		Begin string `json:"begin"`
		End string `json:"end"`
		Period struct {
			Id string `json:"id"`
			Name string `json:"name"`
			Starts string `json:"starts"`
			Ends string `json:"ends"`
		}
	}
	DoctorId string `json:"doctorID"`
}

type Agendas []Agenda
