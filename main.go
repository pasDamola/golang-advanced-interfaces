package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// Always rely on abstractions and not concrete implementations, this is why you use interfaces

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *NormalTruck) LoadCargo() error {
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	return nil
}

func (t *ElectricTruck) LoadCargo() error {
	return nil
}

func (t *ElectricTruck) UnloadCargo() error {
	return nil
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error truck did not load correctly %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error truck did not unload correctly %w", err)
	}
	return nil
}

func main() {

	if err := processTruck(&NormalTruck{id: "1"}); err != nil {
		log.Fatalf("Error processing truck %s\n", err)
	}

	if err := processTruck(&ElectricTruck{id: "1"}); err != nil {
		log.Fatalf("Error processing truck %s\n", err)
	}
}
