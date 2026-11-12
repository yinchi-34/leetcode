package solution

import (
	"math"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected []float64
	}{
		{
			name:     "Example 1",
			celsius:  36.50,
			expected: []float64{309.65, 97.70},
		},
		{
			name:     "Example 2",
			celsius:  122.11,
			expected: []float64{395.26, 251.798},
		},
		{
			name:     "Zero",
			celsius:  0,
			expected: []float64{273.15, 32.00},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := convertTemperature(test.celsius)
			if len(result) != len(test.expected) {
				t.Fatalf("convertTemperature(%f) = %v; expected %v", test.celsius, result, test.expected)
			}
			// Floating-point results are compared within the 1e-5 tolerance
			// the problem allows, not with ==.
			for i := range result {
				if math.Abs(result[i]-test.expected[i]) > 1e-5 {
					t.Errorf("convertTemperature(%f) = %v; expected %v", test.celsius, result, test.expected)
				}
			}
		})
	}
}
