package main

import "chain-of-responsibility/pkg"

func main() {
	device1 := pkg.Device{Name: "Device-1"}
	updateSvc := pkg.UpdateDataService{Name: "UpdateDataService-1"}
	saveSvc := pkg.ServiceDataSave{}

	device1.SetNext(&updateSvc)
	updateSvc.SetNext(&saveSvc)

	data := pkg.Data{}

	device1.Execute(&data)
	updateSvc.Execute(&data)
	saveSvc.Execute(&data)
}
