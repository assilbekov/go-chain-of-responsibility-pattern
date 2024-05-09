package pkg

import "fmt"

type ServiceDataSave struct {
	Next Service
}

func (d *ServiceDataSave) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Data was not updated yet")
	}
	if !data.GetSource {
		fmt.Println("Data was not got yet")
	}

	fmt.Println("Data is saved!")
}

func (d *ServiceDataSave) SetNext(s Service) {
	d.Next = s
}
