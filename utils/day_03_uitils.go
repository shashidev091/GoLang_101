package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	floatText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(floatText), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	defaultValue := 10000.00
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to read the file.")
	}
	float64Text := string(data)
	resultVal, err := strconv.ParseFloat(float64Text, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse the float value.")
	}

	return resultVal, nil
}
