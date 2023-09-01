package main

// Реализовать паттерн «адаптер» на любом примере.

// итерфейс из другого пакета, под который нужно адаптироваться
type TemperatureSensor interface {
	SendStatuses()
}

// структура с какой-то логикой, её нужно адаптировать под нужный интерфейс.
type Computer struct {
}

func (computer *Computer) TemperatureGPU() int {
	return 54
}

func (computer *Computer) TemperatureCPU() int {
	return 56
}

// реализация адаптера
type ComputerAdapter struct {
	computer *Computer
}

func (adapter *ComputerAdapter) SendStatuses() {
	_ = adapter.computer.TemperatureCPU()
	_ = adapter.computer.TemperatureGPU()
	// ... логика отправки статусов
}
