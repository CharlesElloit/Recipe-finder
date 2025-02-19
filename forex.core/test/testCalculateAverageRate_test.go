package test

import (
	"testing"

	forexCore "api.forex.com/forex.core/lib"
)

func TestCalculateAverageRate(t *testing.T) {
	tests := []struct {
		selling_rate float64
		buying_rate  float64
		expected     float64
	}{
		{150.25, 149.75, 150.00},
		{200.00, 100.00, 150.00},
		{0.00, 150.00, 0},
		{100.00, 0.00, 0},
		{-100.00, 150.00, 0},
		{100.00, -150.00, 0},
	}

	for _, tt := range tests {
		results := forexCore.CalculateAveragegRate(tt.selling_rate, tt.buying_rate)
		if results != tt.expected {
			t.Errorf("CalculateAverageRate(%f, %f) = %f; expected %f",
				tt.selling_rate, tt.buying_rate, results, tt.expected,
			)
		}
	}
}
