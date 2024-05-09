package pkg

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Println("Data from device: [%s] was already got" + d.Name)
	}
	if d.Next != nil {
		fmt.Println("Getting data from device: [%s]" + d.Name)
		d.Next.Execute(data)
		data.GetSource = true
	}
}

func (d *Device) SetNext(s Service) {
	d.Next = s
}
