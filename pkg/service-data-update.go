package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (d *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Println("Data from device: [%s] was already got" + d.Name)
	}
	if d.Next != nil {
		fmt.Println("Getting data from device: [%s]" + d.Name)
		d.Next.Execute(data)
		data.UpdateSource = true
	}
}

func (d *UpdateDataService) SetNext(s Service) {
	d.Next = s
}
