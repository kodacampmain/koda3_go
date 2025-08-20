package minitask1

import (
	"errors"
)

var temperatureConstant = map[string]int{
	"C": 5,
	"R": 4,
	"K": 5,
	"F": 9,
}

func TemperatureConversion(temperature float32, fromUnit, targetUnit string) (float32, error) {
	result := temperature
	// asumsikan fromUnit dan targetUnit selalu huruf besar
	if fromUnit == "K" {
		result -= 273
	} else if fromUnit == "F" {
		result -= 32
	} else if fromUnit != "C" && fromUnit != "R" {
		return 0, errors.New("invalid from unit")
	}

	result = float32(temperatureConstant[targetUnit]) / float32(temperatureConstant[fromUnit]) * result

	if targetUnit == "K" {
		result += 273
	} else if targetUnit == "F" {
		result += 32
	} else if targetUnit != "C" && targetUnit != "R" {
		return 0, errors.New("invalid target unit")
	}

	return result, nil
}
