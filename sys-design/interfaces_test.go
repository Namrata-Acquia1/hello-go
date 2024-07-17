package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestTV(t *testing.T) {
	tv := NewTV().(*TV)

	// initial state
	if tv.IsOn() {
		t.Error("Expected TV to be off initially, but it is on")
	}

	tv.TogglePower()
	if !tv.IsOn() {
		t.Error("Expected TV to be on after TogglePower, but it is off")
	}

	tv.NumberPad(5)
	if tv.channel != 5 {
		t.Errorf("Expected TV channel to be 5, got %d", tv.channel)
	}

	tv.TogglePower() // Turn TV off
	expectedErrorMsg := "TV is off. Cannot adjust channel."
	lastOutput := captureOutput(func() {
		tv.NumberPad(200)
	})
	if !strings.Contains(lastOutput, expectedErrorMsg) {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, lastOutput)
	}
}

func TestCeilingFan(t *testing.T) {
	fan := NewCeilingFan().(*CeilingFan)

	// initial state
	if fan.IsOn() {
		t.Error("Expected Ceiling Fan to be off initially, but it is on")
	}

	fan.TogglePower()
	if !fan.IsOn() {
		t.Error("Expected Ceiling Fan to be on after TogglePower, but it is off")
	}

	fan.NumberPad(3)
	if fan.speed != 3 {
		t.Errorf("Expected Ceiling Fan speed to be 3, got %d", fan.speed)
	}

	fan.TogglePower()
	expectedErrorMsg := "Fan is off. Cannot adjust speed."
	lastOutput := captureOutput(func() {
		fan.NumberPad(10)
	})
	if !strings.Contains(lastOutput, expectedErrorMsg) {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, lastOutput)
	}
}

func TestAirConditioner(t *testing.T) {
	ac := NewAirConditioner().(*AirConditioner)

	// initial state
	if ac.IsOn() {
		t.Error("Expected Air Conditioner to be off initially, but it is on")
	}

	ac.TogglePower()
	if !ac.IsOn() {
		t.Error("Expected Air Conditioner to be on after TogglePower, but it is off")
	}

	ac.NumberPad(25)
	if ac.temperature != 25 {
		t.Errorf("Expected Air Conditioner temperature to be 25, got %d", ac.temperature)
	}

	ac.TogglePower()
	expectedErrorMsg := "AC is off. Cannot adjust temperature."
	lastOutput := captureOutput(func() {
		ac.NumberPad(35)
	})
	if !strings.Contains(lastOutput, expectedErrorMsg) {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, lastOutput)
	}
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out := make([]byte, 1024)
	n, _ := r.Read(out)
	os.Stdout = rescueStdout

	return strings.TrimSpace(string(out[:n]))
}

func TestMain(m *testing.M) {
	fmt.Println("Running tests...")
	code := m.Run()
	os.Exit(code)
}
