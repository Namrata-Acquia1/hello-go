package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Device interface
type Device interface {
	Name() string
	TogglePower()
	Adjust(value int)
	IsOn() bool
}

// TV struct
type TV struct {
	isOn    bool
	channel int
}

// TogglePower method for TV
func (tv *TV) TogglePower() {
	if tv.isOn {
		fmt.Println("TV Power off")
		tv.isOn = false
	} else {
		fmt.Println("TV Power on")
		tv.isOn = true
	}
}

// Adjust method for TV (adjust channel)
func (tv *TV) Adjust(value int) {
	if !tv.isOn {
		fmt.Println("TV is off. Cannot adjust channel.")
		return
	}

	if value < 1 || value > 150 {
		fmt.Println("Invalid channel value. Channel must be between 1 and 150.")
		return
	}

	tv.channel = value
	fmt.Printf("TV Channel set to %d\n", tv.channel)
}

// IsOn method for TV
func (tv *TV) IsOn() bool {
	return tv.isOn
}

// Name method for TV
func (tv *TV) Name() string {
	return "TV"
}

// Factory function for TV
func NewTV() Device {
	return &TV{}
}

// CeilingFan struct
type CeilingFan struct {
	isOn  bool
	speed int
}

// TogglePower method for CeilingFan
func (fan *CeilingFan) TogglePower() {
	if fan.isOn {
		fmt.Println("Fan Power off")
		fan.isOn = false
	} else {
		fmt.Println("Fan Power on")
		fan.isOn = true
	}
}

// Adjust method for CeilingFan (adjust speed)
func (fan *CeilingFan) Adjust(value int) {
	if !fan.isOn {
		fmt.Println("Fan is off. Cannot adjust speed.")
		return
	}

	if value < 1 || value > 5 {
		fmt.Println("Invalid speed value. Speed must be between 1 and 5.")
		return
	}

	fan.speed = value
	fmt.Printf("Fan Speed set to %d\n", fan.speed)
}

// IsOn method for CeilingFan
func (fan *CeilingFan) IsOn() bool {
	return fan.isOn
}

// Name method for CeilingFan
func (fan *CeilingFan) Name() string {
	return "Ceiling Fan"
}

// Factory function for Ceiling Fan
func NewCeilingFan() Device {
	return &CeilingFan{}
}

// AirConditioner struct
type AirConditioner struct {
	isOn        bool
	temperature int
}

// TogglePower method for AirConditioner
func (ac *AirConditioner) TogglePower() {
	if ac.isOn {
		fmt.Println("AC Power off")
		ac.isOn = false
	} else {
		fmt.Println("AC Power on")
		ac.isOn = true
	}
}

// Adjust method for AirConditioner (adjust temperature)
func (ac *AirConditioner) Adjust(value int) {
	if !ac.isOn {
		fmt.Println("AC is off. Cannot adjust temperature.")
		return
	}

	if value < 16 || value > 30 {
		fmt.Println("Invalid temperature value. Temperature must be between 16°C and 30°C.")
		return
	}

	ac.temperature = value
	fmt.Printf("AC Temperature set to %d°C\n", ac.temperature)
}

// IsOn method for AirConditioner
func (ac *AirConditioner) IsOn() bool {
	return ac.isOn
}

// Name method for AirConditioner
func (ac *AirConditioner) Name() string {
	return "Air Conditioner"
}

// Factory function for Air Conditioner
func NewAirConditioner() Device {
	return &AirConditioner{}
}

// controlDevice function to interact with a specific device
func controlDevice(device Device) {
	clearScreen()

	fmt.Printf("Device: %s\n", device.Name())
	fmt.Println("1. Power")
	fmt.Println("2. Number Pad")
	fmt.Println("3. Switch Device")

	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		device.TogglePower()
	case "2":
		var valueStr string
		fmt.Print("Enter value: ")
		fmt.Scanln(&valueStr)
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}
		device.Adjust(value)
	case "3":
		fmt.Println("Returning to main menu.")
		return
	default:
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}

	fmt.Println("Press Enter to continue.")
	fmt.Scanln()
	controlDevice(device) // Recursive call to keep interacting with the same device
}

// clearScreen clears the console screen
func clearScreen() {
	cmd := exec.Command("clear") // for Linux/OSX
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// main function
func main() {
	// Initialize appliances using factory functions
	tv := NewTV()
	fan := NewCeilingFan()
	ac := NewAirConditioner()

	for {
		clearScreen()

		fmt.Println("Select an appliance:")
		fmt.Println("1. TV")
		fmt.Println("2. Ceiling Fan")
		fmt.Println("3. Air Conditioner")
		fmt.Println("4. Exit")

		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		var device Device

		switch choice {
		case "1":
			device = tv
		case "2":
			device = fan
		case "3":
			device = ac
		case "4":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
			continue
		}

		controlDevice(device)
	}
}