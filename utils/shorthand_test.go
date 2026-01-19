package utils

import (
	"gohm/test_utils"
	"math"
	"testing"
)

func TestParseShorthand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		targets  []string
		expected float64
	}{
		// Plain numbers - no identifiers
		{"integer only", "10", []string{"V"}, 10},
		{"float only", "10.5", []string{"V"}, 10.5},
		{"zero only", "0", []string{"V"}, 0},
		{"large integer only", "12345", []string{"V"}, 12345},
		{"small float only", "0.001", []string{"V"}, 0.001},

		// Kilo prefix - with and without decimal, with and without target
		{"kilo int", "10k", []string{"V"}, 10000},
		{"kilo int with target", "10kV", []string{"V"}, 10000},
		{"kilo float", "1.5k", []string{"V"}, 1500},
		{"kilo float with target", "1.5kV", []string{"V"}, 1500},

		// Mega prefix - with and without decimal, with and without target
		{"mega int", "1M", []string{"V"}, 1000000},
		{"mega int with target", "1MV", []string{"V"}, 1000000},
		{"mega float", "2.5M", []string{"V"}, 2500000},
		{"mega float with target", "2.5MV", []string{"V"}, 2500000},

		// Giga prefix - with and without decimal, with and without target
		{"giga int", "1G", []string{"V"}, 1000000000},
		{"giga int with target", "1GV", []string{"V"}, 1000000000},
		{"giga float", "1.2G", []string{"V"}, 1200000000},
		{"giga float with target", "1.2GV", []string{"V"}, 1200000000},

		// Tera prefix - with and without decimal, with and without target
		{"tera int", "1T", []string{"V"}, 1000000000000},
		{"tera int with target", "1TV", []string{"V"}, 1000000000000},
		{"tera float", "1.5T", []string{"V"}, 1500000000000},
		{"tera float with target", "1.5TV", []string{"V"}, 1500000000000},

		// Milli prefix - with and without decimal, with and without target
		{"milli int", "10m", []string{"V"}, 0.01},
		{"milli int with target", "10mV", []string{"V"}, 0.01},
		{"milli float", "2.5m", []string{"A"}, 0.0025},
		{"milli float with target", "2.5mA", []string{"A"}, 0.0025},

		// Micro prefix - with and without decimal, with and without target
		{"micro int", "1μ", []string{"V"}, 1e-6},
		{"micro int with target", "1μV", []string{"V"}, 1e-6},
		{"micro int 2", "2μ", []string{"V"}, 2e-6},

		// Nano prefix - with and without decimal, with and without target
		{"nano int", "10n", []string{"V"}, 0.00000001},
		{"nano int with target", "10nV", []string{"V"}, 0.00000001},
		{"nano float", "5n", []string{"V"}, 0.000000005},
		{"nano float with target", "5nV", []string{"V"}, 0.000000005},

		// Pico prefix - with and without decimal, with and without target
		{"pico int", "10p", []string{"V"}, 0.00000000001},
		{"pico int with target", "10pV", []string{"V"}, 0.00000000001},
		{"pico float", "1.5p", []string{"V"}, 0.0000000000015},
		{"pico float with target", "1.5pV", []string{"V"}, 0.0000000000015},

		// Different unit targets
		{"kilo resistance", "1kR", []string{"r", "R"}, 1000},
		{"kilo resistance with ronna and target - same identifiers", "1RR", []string{"r", "R"}, math.Pow10(27)},
		{"milli amps", "2mA", []string{"a", "A"}, 0.002},
		{"mega watts", "5MW", []string{"w", "W"}, 5000000},

		// Frequency (Hz) - multi-character target
		{"hertz only", "100Hz", []string{"Hz"}, 100},
		{"hertz float", "50.5Hz", []string{"Hz"}, 50.5},
		{"kilo hertz", "1kHz", []string{"Hz"}, 1000},
		{"kilo hertz float", "2.4kHz", []string{"Hz"}, 2400},
		{"mega hertz", "100MHz", []string{"Hz"}, 100000000},
		{"mega hertz float", "2.4MHz", []string{"Hz"}, 2400000},
		{"giga hertz", "1GHz", []string{"Hz"}, 1000000000},
		{"giga hertz float", "2.4GHz", []string{"Hz"}, 2400000000},
		{"milli hertz", "10mHz", []string{"Hz"}, 0.01},
		{"micro hertz", "1μHz", []string{"Hz"}, 1e-6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseShorthand(tt.input, tt.targets)
			test_utils.AssertEquals(t, result, tt.expected)
		})
	}
}

func TestParseRKMCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   rune
		expected float64
	}{
		// Resistance (R target) - using R as decimal point
		{"R no leading digits", "R47", 'R', 0.47},
		{"R single leading digit", "4R7", 'R', 4.7},
		{"R two leading digits", "47R", 'R', 47},
		{"R three leading digits", "470R", 'R', 470},
		{"R two leading one trailing", "47R5", 'R', 47.5},

		// Resistance with K (kilo) prefix
		{"K no leading digits", "K47", 'R', 470},
		{"K single leading digit", "4K7", 'R', 4700},
		{"K two leading digits", "47K", 'R', 47000},
		{"K two leading one trailing", "47K3", 'R', 47300},
		{"K three leading digits", "470K", 'R', 470000},

		// Resistance with M (mega) prefix
		{"M no leading digits", "M47", 'R', 470000},
		{"M single leading digit", "4M7", 'R', 4700000},
		{"M two leading digits", "47M", 'R', 47000000},
		{"M two leading one trailing", "47M3", 'R', 47300000},

		// Resistance with G (giga) prefix
		{"G single leading digit", "4G7", 'R', 4700000000},
		{"G two leading digits", "47G", 'R', 47000000000},

		// Resistance with T (tera) prefix
		{"T single leading digit", "1T0", 'R', 1000000000000},
		{"T two leading digits", "10T", 'R', 10000000000000},

		// Resistance with L (milli) prefix
		{"L no leading digits", "L47", 'R', 0.00047},
		{"L single leading digit", "4L7", 'R', 0.0047},
		{"L two leading digits", "47L", 'R', 0.047},

		// Capacitance (F target) - pico (using integer values to avoid float precision)
		{"pico single leading digit", "4p7", 'F', 4.7e-12},
		{"pico two leading digits", "47p", 'F', 47e-12},
		{"pico integer", "10p", 'F', 10e-12},

		// Capacitance (F target) - nano (using integer values to avoid float precision)
		{"nano two leading digits no trailing", "10n", 'F', 10e-9},
		{"nano single leading digit", "1n0", 'F', 1e-9},

		// Capacitance (F target) - micro (μ is multi-byte, skipping for now)
		{"micro single leading digit", "4μ7", 'F', 4.7e-6},
		{"micro two leading digits", "47μ", 'F', 47e-6},

		// Edge cases - minimum length (2 chars)
		{"min length R prefix", "1R", 'R', 1},
		{"min length K prefix", "1K", 'R', 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseRKMCode(tt.input, tt.target)
			if err != nil {
				t.Error(err)
			}
			test_utils.AssertEquals(t, result, tt.expected)
		})
	}
}

func TestParseRKMCodeErrors(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		target      rune
		expectedErr string
	}{
		// Invalid length
		{"too short - single char", "R", 'R', "invalid or unsupported: RKM notation R"},
		{"too long - 6 chars", "123K56", 'R', "invalid or unsupported: RKM notation 123K56"},

		// Unimplemented target
		{"unimplemented target V", "4K7", 'V', "invalid or unsupported: RKM target V"},
		{"unimplemented target L", "4K7", 'L', "invalid or unsupported: RKM target L"},

		// Invalid prefix for target
		{"invalid prefix for resistance", "4p7", 'R', "invalid or unsupported: prefix p for target R"},
		{"invalid prefix for resistance - lowercase k", "4k7", 'R', "invalid or unsupported: prefix k for target R"},
		{"invalid prefix X", "4X7", 'R', "invalid or unsupported: prefix X for target R"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseRKMCode(tt.input, tt.target)
			if err == nil {
				t.Errorf("ParseRKMCode(%q, %q) expected error, got nil", tt.input, string(tt.target))
				return
			}
			test_utils.AssertContains(t, err.Error(), tt.expectedErr)
		})
	}
}

func TestParseShorthandPanics(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		targets     []string
		expectedErr string
	}{
		// Unknown SI prefix errors
		{"unknown prefix", "10x", []string{"V"}, "invalid or unsupported: si prefix x"},
		{"unknown prefix with target", "10xV", []string{"V"}, "invalid or unsupported: si prefix x"},
		{"unknown prefix after number", "10.1b", []string{"V"}, "invalid or unsupported: si prefix b"},
		{"unknown prefix after number with target", "10.1bV", []string{"V"}, "invalid or unsupported: si prefix b"},

		// Target not in valid targets (error 2)
		{"duplicate target that has no valid prefix like ronna and resistance", "10VV", []string{"V"}, "invalid or unsupported: si prefix V"},
		{"wrong target", "10kR", []string{"V"}, "invalid: type identifier R for target"},
		{"wrong target after prefix", "10mW", []string{"V"}, "invalid: type identifier W for target"},

		// Extra characters after valid pattern (error 3)
		{"extra chars after target", "10mVolt", []string{"V"}, "invalid or unsupported: type identifier olt"},
		{"extra chars after prefix and target", "10mVDD", []string{"V"}, "invalid or unsupported: type identifier DD"},
		{"extra chars after target no prefix", "10VSS", []string{"V"}, "invalid or unsupported: si prefix V"},
		{"extra number after target", "10kV5", []string{"V"}, "invalid or unsupported: type identifier 5"},
		{"extra prefix after target", "10kVk", []string{"V"}, "invalid or unsupported: type identifier k"},

		// Frequency (Hz) panic cases
		{"Hz unknown prefix", "10xHz", []string{"Hz"}, "invalid or unsupported: si prefix x"},
		{"Hz wrong target", "10kV", []string{"Hz"}, "invalid: type identifier V for target"},
		{"Hz extra chars after target", "10kHzz", []string{"Hz"}, "invalid or unsupported: type identifier z"},
		{"Hz extra chars after target no prefix", "10Hzz", []string{"Hz"}, "invalid or unsupported: si prefix H"},
		{"Hz partial match", "10H", []string{"Hz"}, "invalid or unsupported: si prefix H"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test_utils.ExpectPanic(t, tt.expectedErr, func() {
				ParseShorthand(tt.input, tt.targets)
			})
		})
	}
}

func TestGetLeadingDigits(t *testing.T) {
	tests := []struct {
		name                    string
		input                   string
		input_supports_decimals bool
		expectedVal             float64
		expectedIndex           int
		expectError             bool
	}{
		{"integer", "123", true, 123, 3, false},
		{"float", "12.5", true, 12.5, 4, false},
		{"integer with suffix", "10k", true, 10, 2, false},
		{"float with suffix", "4.7M", true, 4.7, 3, false},
		{"no digits", "abc", true, 0, 0, false},
		{"empty string", "", true, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, index, err := get_leading_digits(tt.input, tt.input_supports_decimals)

			if tt.expectError && err == nil {
				t.Errorf("get_leading_digits(%q) expected error, got nil", tt.input)
			}
			if !tt.expectError && err != nil {
				t.Errorf("get_leading_digits(%q) unexpected error: %v", tt.input, err)
			}
			test_utils.AssertEquals(t, val, tt.expectedVal)
			test_utils.AssertEquals(t, index, tt.expectedIndex)
		})
	}
}
