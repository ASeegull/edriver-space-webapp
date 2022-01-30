package sorts

import "github.com/ASeegull/edriver-space-webapp/model"

const (
	CarMake    = "Mazda"
	CarRegDate = "31/01/2022"
	CarYear    = "2015"
)

func SearchFinesByNumberPlate(fines model.Fines, numberplate string) []model.CarsFine {
	var tempFines []model.CarsFine

	for _, fine := range fines.CarsFines {
		if fine.VehicleRegistrationNumber == numberplate {
			tempFines = append(tempFines, fine)
		}
	}
	return tempFines
}

func GetCarListFromFines(fines model.Fines) []model.Car {
	var cars []model.Car
	carsmap := make(map[string]model.Car)

	for _, fine := range fines.CarsFines {
		var tempcar model.Car
		if _, excist := carsmap[fine.VehicleRegistrationNumber]; excist {
			buffercar := carsmap[fine.VehicleRegistrationNumber]
			buffercar.FinesNum++
			carsmap[fine.VehicleRegistrationNumber] = buffercar
			continue

		} else {
			tempcar.FinesNum = 1
			tempcar.RegistrationNum = fine.VehicleRegistrationNumber
			tempcar.Make = CarMake
			tempcar.RegistrationDate = CarRegDate
			tempcar.Year = CarYear
			carsmap[fine.VehicleRegistrationNumber] = tempcar
		}
	}

	for _, val := range carsmap {
		cars = append(cars, val)
	}

	return cars
}
