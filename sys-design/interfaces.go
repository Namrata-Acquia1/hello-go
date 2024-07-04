package main

import (
	"fmt"
)

// Interface for all devices
type Device interface {
	PowerOn()
	PowerOff()
	Adjust(value int)
	IsOn() bool
}

// TV struct
type TV struct {
	isOn    bool
	channel int
}

func (tv *TV) PowerOn() {
	if !tv.isOn {
		fmt.Println("TV Power on")
		tv.isOn = true
	} else {
		fmt.Println("TV is already on")
	}
}

func (tv *TV) PowerOff() {
	if tv.isOn {
		fmt.Println("TV Power off")
		tv.isOn = false
	} else {
		fmt.Println("TV is already off")
	}
}

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

func (tv *TV) IsOn() bool {
	return tv.isOn
}

// CeilingFan struct
type CeilingFan struct {
	isOn  bool
	speed int
}

func (fan *CeilingFan) PowerOn() {
	if !fan.isOn {
		fmt.Println("Fan Power on")
		fan.isOn = true
	} else {
		fmt.Println("Fan is already on")
	}
}

func (fan *CeilingFan) PowerOff() {
	if fan.isOn {
		fmt.Println("Fan Power off")
		fan.isOn = false
	} else {
		fmt.Println("Fan is already off")
	}
}

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

func (fan *CeilingFan) IsOn() bool {
	return fan.isOn
}

// AirConditioner struct
type AirConditioner struct {
	isOn        bool
	temperature int
}

func (ac *AirConditioner) PowerOn() {
	if !ac.isOn {
		fmt.Println("AC Power on")
		ac.isOn = true
	} else {
		fmt.Println("AC is already on")
	}
}

func (ac *AirConditioner) PowerOff() {
	if ac.isOn {
		fmt.Println("AC Power off")
		ac.isOn = false
	} else {
		fmt.Println("AC is already off")
	}
}

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

func (ac *AirConditioner) IsOn() bool {
	return ac.isOn
}

type DeviceConfig struct {
	device Device
}

// controlDevice function
func controlDevice(dc DeviceConfig) {
	var action int

	switch dc.device.(type) {
	case *TV:
		fmt.Println("\nTV selected. What would you like to do?")
		fmt.Println("1. Power On/Off")
		fmt.Println("2. Adjust Channel")
	case *CeilingFan:
		fmt.Println("\nCeiling Fan selected. What would you like to do?")
		fmt.Println("1. Power On/Off")
		fmt.Println("2. Adjust Speed")
	case *AirConditioner:
		fmt.Println("\nAir Conditioner selected. What would you like to do?")
		fmt.Println("1. Power On/Off")
		fmt.Println("2. Adjust Temperature")
	}
	fmt.Println("3. Back to main menu")

	fmt.Print("Enter your action: ")

	_, err := fmt.Scanln(&action)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch action {
	case 1:
		if dc.device.IsOn() {
			dc.device.PowerOff()
		} else {
			dc.device.PowerOn()
		}
	case 2:
		if dc.device.IsOn() {
			var value int
			fmt.Print("Enter value: ")
			_, err := fmt.Scanln(&value)
			if err != nil {
				fmt.Println("Error reading value:", err)
				return
			}
			dc.device.Adjust(value)
		} else {
			fmt.Println("Device is off. Cannot adjust.")
		}
	case 3:
		fmt.Println("Returning to main menu...")
		return
	default:
		fmt.Println("Invalid action. Please enter a number between 1 and 3.")
	}
}

// RemoteControl struct to manage the appliances
type RemoteControl struct {
	devices []DeviceConfig
}

// main function
func main() {
	// Initialize appliances
	tv := &TV{}
	fan := &CeilingFan{}
	ac := &AirConditioner{}

	remote := &RemoteControl{
		devices: []DeviceConfig{
			{device: tv},
			{device: fan},
			{device: ac},
		},
	}

	fmt.Println("Welcome to the Universal Remote Control!")

	for {
		fmt.Println("\nSelect an appliance:")
		fmt.Println("1. TV")
		fmt.Println("2. Ceiling Fan")
		fmt.Println("3. Air Conditioner")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case 1, 2, 3:
			controlDevice(remote.devices[choice-1])
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}
