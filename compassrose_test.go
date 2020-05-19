// Package compassrose provides data structures and minimal functions to use a 32 point compass
package compassrose

import (
	"testing"
)

func TestDegreeToHeading(t *testing.T) {

	type args struct {
		inDegrees float32
		level     int
		standard  bool
		valid     string
	}

	tests := []struct {
		name     string
		testargs args
	}{
		{name: "Positive",
			testargs: args{
				inDegrees: 22.5,
				level:     3,
				standard:  true,
				valid:     "NNE"}},
		{name: "Negative",
			testargs: args{
				inDegrees: -22.5,
				level:     3,
				standard:  true,
				valid:     "NNW"}},
		{name: "Zero",
			testargs: args{
				inDegrees: -0.0,
				level:     3,
				standard:  true,
				valid:     "N"}},
		{name: "Huge",
			testargs: args{
				inDegrees: 3530.0,
				level:     3,
				standard:  true,
				valid:     "WNW"}},
		{name: "Backwards",
			testargs: args{
				inDegrees: -3530.0,
				level:     3,
				standard:  true,
				valid:     "ENE"}},
		{name: "VerySmall",
			testargs: args{
				inDegrees: -0.00025527,
				level:     3,
				standard:  true,
				valid:     "N"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tResult, _ := DegreeToHeading(tt.testargs.inDegrees, tt.testargs.level, tt.testargs.standard)
			if tResult != tt.testargs.valid {
				t.Errorf("Expected %v, got %v for %v\n", tt.testargs.valid, tResult, tt.testargs.inDegrees)
			}
		})
	}
}

func TestDegreeToHeadingSouthfacing(t *testing.T) {

	type args struct {
		inDegrees float32
		level     int
		standard  bool
		valid     string
	}

	tests := []struct {
		name     string
		testargs args
	}{
		{name: "Positive",
			testargs: args{
				inDegrees: 22.5,
				level:     3,
				standard:  false,
				valid:     "SSW"}},
		{name: "Negative",
			testargs: args{
				inDegrees: -22.5,
				level:     3,
				standard:  false,
				valid:     "SSE"}},
		{name: "Zero",
			testargs: args{
				inDegrees: 0.0,
				level:     3,
				standard:  false,
				valid:     "S"}},
		{name: "Huge",
			testargs: args{
				inDegrees: 3530.0,
				level:     3,
				standard:  true,
				valid:     "ESE"}},
		{name: "Backwards",
			testargs: args{
				inDegrees: -3530.0,
				level:     3,
				standard:  true,
				valid:     "WSW"}},
		{name: "VerySmall",
			testargs: args{
				inDegrees: -0.00025527,
				level:     3,
				standard:  true,
				valid:     "S"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tResult, _ := DegreeToHeadingSouthfacing(tt.testargs.inDegrees, tt.testargs.level, tt.testargs.standard)
			if tResult != tt.testargs.valid {
				t.Errorf("Expected %v, got %v for %v\n", tt.testargs.valid, tResult, tt.testargs.inDegrees)
			}
		})
	}
}

func TestInvalidLevel(t *testing.T) {

	type args struct {
		inDegrees float32
		level     int
		standard  bool
		valid     string
	}

	tests := []struct {
		name     string
		testargs args
	}{
		{name: "TooLarge",
			testargs: args{
				inDegrees: 22.5,
				level:     7,
				standard:  true,
				valid:     "ERROR"},
		},
		{name: "Negative",
			testargs: args{
				inDegrees: 22.5,
				level:     -2,
				standard:  false,
				valid:     "ERROR"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tResult, _ := DegreeToHeading(tt.testargs.inDegrees, tt.testargs.level, tt.testargs.standard)
			if tResult != tt.testargs.valid {
				t.Errorf("Expected %v, got %v for %v\n", tt.testargs.valid, tResult, tt.testargs.inDegrees)
			}
		})
	}
}
