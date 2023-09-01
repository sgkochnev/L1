package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

var temperatures = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.0, 7, 0, -1, -2, 39.8, -12.5}

func main() {
	temperatureGroups := make(map[int][]float64)

	for _, t := range temperatures {
		k := int(t/10) * 10 // зануляем разряд единиц
		temperatureGroups[k] = append(temperatureGroups[k], t)
	}

	for k, v := range temperatureGroups {
		fmt.Println(k, v)
	}
}

// func main() {
// 	temperatureGroups := make(map[string][]float64)

// 	for _, t := range temperatures {
// 		sign := ""
// 		if -10 < t && t < 0 {
// 			sign = "-"
// 		}
// 		if t >= 0 {
// 			sign = "+"
// 		}
// 		k := fmt.Sprintf("%s%d", sign, int(t/10)*10)
// 		temperatureGroups[k] = append(temperatureGroups[k], t)
// 	}

// 	for k, v := range temperatureGroups {
// 		fmt.Println(k, v)
// 	}
// }
