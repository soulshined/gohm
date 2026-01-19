package identify

import (
	"gohm/test_utils"
	"strings"
	"testing"
)

// region Capacitance Tests
func TestCmdCapacitanceHandler(t *testing.T) {
	tests := []struct {
		name     string
		eiaValue string
		format   string
		contains []string
	}{
		{
			name:     "2 digit numeric",
			eiaValue: "47",
			format:   "abbr",
			contains: []string{"nominal=47pF"},
		},
		{
			name:     "2 digit EIA-198 A0",
			eiaValue: "A0",
			format:   "abbr",
			contains: []string{"nominal=1pF"},
		},
		{
			name:     "2 digit EIA-198 A1",
			eiaValue: "A1",
			format:   "abbr",
			contains: []string{"nominal=10pF"},
		},
		{
			name:     "2 digit EIA-198 J2",
			eiaValue: "J2",
			format:   "abbr",
			contains: []string{"nominal=", "pF"},
		},
		{
			name:     "2 digit EIA-198 S3",
			eiaValue: "S3",
			format:   "abbr",
			contains: []string{"nominal=", "nF"},
		},
		{
			name:     "3 digit 104",
			eiaValue: "104",
			format:   "abbr",
			contains: []string{"nominal=", "nF"},
		},
		{
			name:     "3 digit 473",
			eiaValue: "473",
			format:   "abbr",
			contains: []string{"nominal=", "nF"},
		},
		{
			name:     "3 digit 222",
			eiaValue: "222",
			format:   "abbr",
			contains: []string{"nominal=", "nF"},
		},
		{
			name:     "4 digit with tolerance J",
			eiaValue: "104J",
			format:   "abbr",
			contains: []string{"nominal=", "min=", "max="},
		},
		{
			name:     "4 digit with tolerance K",
			eiaValue: "473K",
			format:   "abbr",
			contains: []string{"nominal=", "nF"},
		},
		{
			name:     "4 digit with tolerance M",
			eiaValue: "222M",
			format:   "abbr",
			contains: []string{"nominal=", "nF"},
		},
		{
			name:     "raw format",
			eiaValue: "104",
			format:   "raw",
			contains: []string{"nominal="},
		},
		{
			name:     "json format",
			eiaValue: "104J",
			format:   "json",
			contains: []string{`"nominal":`, `"nominalAbbreviated":`, `"actualMin":`, `"actualMax":`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_capacitor_handler,
				map[string]string{
					"format": tt.format,
					"eiac":   tt.eiaValue,
				},
				nil,
				[]string{tt.eiaValue},
			)
			result := cmd_capacitor_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmdCapacitanceHandlerPanics(t *testing.T) {
	tests := []struct {
		name     string
		eiaValue string
		expected string
	}{
		{
			name:     "invalid 2 digit identifier",
			eiaValue: "I0",
			expected: "invalid or unsupported: EIA-198 identifier I",
		},
		{
			name:     "invalid 3 digit first char",
			eiaValue: "A04",
			expected: "invalid: capacitor value A04",
		},
		{
			name:     "invalid 3 digit third char",
			eiaValue: "10A",
			expected: "invalid: capacitor value 10A",
		},
		{
			name:     "invalid 3 digit middle identifier",
			eiaValue: "1A4",
			expected: "invalid or unsupported: capacitor decimal identifier A",
		},
		{
			name:     "invalid tolerance identifier",
			eiaValue: "104X",
			expected: "invalid or unsupported: capacitor tolerance identifier: X",
		},
		{
			name:     "too long value",
			eiaValue: "10000",
			expected: "unsupported: 5 digit codes",
		},
		{
			name:     "single character",
			eiaValue: "1",
			expected: "unsupported: 1 digit codes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_capacitor_handler,
				map[string]string{
					"format": "abbr",
					"eiac":   tt.eiaValue,
				},
				nil,
				[]string{},
			)
			test_utils.ExpectPanic(t, tt.expected, func() {
				cmd_capacitor_handler(cmd)
			})
		})
	}
}

func TestCmdCapacitanceHandlerNoFlag(t *testing.T) {
	cmd := test_utils.CreateTestCommand(
		cmd_capacitor_handler,
		map[string]string{"format": "abbr"},
		nil,
		[]string{"104"},
	)
	test_utils.ExpectPanic(t, "unsupported: identify capacitor flags", func() {
		cmd_capacitor_handler(cmd)
	})
}

func TestGetCapacitanceFromEIA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"A0", "A0", 1e-12},
		{"B1", "B1", 11e-12},
		{"C2", "C2", 120e-12},
		{"D3", "D3", 1300e-12},
		{"E4", "E4", 15000e-12},
		{"F5", "F5", 160000e-12},
		{"G6", "G6", 1800000e-12},
		{"H7", "H7", 20000000e-12},
		{"J8", "J8", 220000000e-12},
		{"K9", "K9", 2400000000e-12},
		{"100", "100", 10e-12},
		{"101", "101", 100e-12},
		{"102", "102", 1000e-12},
		{"103", "103", 10000e-12},
		{"104", "104", 100000e-12},
		{"105", "105", 1000000e-12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := get_capacitance_from_eia(tt.input, "raw")
			if !strings.Contains(result, "nominal=") {
				t.Errorf("expected result to contain nominal value for input %s", tt.input)
			}
		})
	}
}

//endregion Capacitance Tests

//region Resistance Tests

func TestCmdResistorHandler(t *testing.T) {
	tests := []struct {
		name     string
		bands    []string
		format   string
		contains []string
	}{
		{
			name:     "4 band brown black red gold",
			bands:    []string{"brown", "black", "red", "gold"},
			format:   "abbr",
			contains: []string{"nominal=1kΩ", "min=950Ω", "max=1.05kΩ"},
		},
		{
			name:     "4 band red red brown gold",
			bands:    []string{"red", "red", "brown", "gold"},
			format:   "abbr",
			contains: []string{"nominal=220Ω", "min=209Ω", "max=231Ω"},
		},
		{
			name:     "4 band yellow violet orange silver",
			bands:    []string{"yellow", "violet", "orange", "silver"},
			format:   "abbr",
			contains: []string{"nominal=47kΩ"},
		},
		{
			name:     "5 band brown black black red brown",
			bands:    []string{"brown", "black", "black", "red", "brown"},
			format:   "abbr",
			contains: []string{"nominal=10kΩ"},
		},
		{
			name:     "5 band red red black black brown",
			bands:    []string{"red", "red", "black", "black", "brown"},
			format:   "abbr",
			contains: []string{"nominal=220Ω"},
		},
		{
			name:     "6 band with temp coefficient",
			bands:    []string{"brown", "black", "black", "brown", "brown", "red"},
			format:   "abbr",
			contains: []string{"nominal=1kΩ", "temp_coefficient=50 ppm/K"},
		},
		{
			name:     "raw format",
			bands:    []string{"brown", "black", "red", "gold"},
			format:   "raw",
			contains: []string{"nominal=1000Ω", "min=950Ω", "max=1050Ω"},
		},
		{
			name:     "json format",
			bands:    []string{"brown", "black", "red", "gold"},
			format:   "json",
			contains: []string{`"nominal":1000`, `"nominalAbbreviated":"1kΩ"`, `"actualMin":950`, `"actualMax":1050`},
		},
		{
			name:     "abbreviated color names",
			bands:    []string{"bn", "bk", "rd", "gd"},
			format:   "abbr",
			contains: []string{"nominal=1kΩ"},
		},
		{
			name:     "3 band no tolerance",
			bands:    []string{"brown", "black", "red"},
			format:   "abbr",
			contains: []string{"nominal=1kΩ", "min=800Ω", "max=1.2kΩ"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_resistor_handler,
				map[string]string{"format": tt.format},
				nil,
				tt.bands,
			)
			result := cmd_resistor_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmdResistorHandlerPanics(t *testing.T) {
	tests := []struct {
		name     string
		bands    []string
		expected string
	}{
		{
			name:     "gold as significant digit",
			bands:    []string{"gold", "black", "red", "gold"},
			expected: "invalid: significant digit band color can not be gold",
		},
		{
			name:     "silver as significant digit",
			bands:    []string{"silver", "black", "red", "gold"},
			expected: "invalid: significant digit band color can not be silver",
		},
		{
			name:     "pink as significant digit",
			bands:    []string{"pink", "black", "red", "gold"},
			expected: "invalid: significant digit band color can not be pink",
		},
		{
			name:     "invalid color name",
			bands:    []string{"purple", "black", "red", "gold"},
			expected: "invalid: significant digit band color purple",
		},
		{
			name:     "invalid multiplier color",
			bands:    []string{"brown", "black", "invalid", "gold"},
			expected: "invalid: multiplier band color invalid",
		},
		{
			name:     "invalid tolerance color",
			bands:    []string{"brown", "black", "red", "white"},
			expected: "invalid: tolerance band color white",
		},
		{
			name:     "invalid temp coefficient color",
			bands:    []string{"brown", "black", "black", "brown", "brown", "silver"},
			expected: "invalid: temperature coefficient band color silver",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_resistor_handler,
				map[string]string{"format": "abbr"},
				nil,
				tt.bands,
			)
			test_utils.ExpectPanic(t, tt.expected, func() {
				cmd_resistor_handler(cmd)
			})
		})
	}
}

func TestResistorBandMapping(t *testing.T) {
	expectedColors := []string{
		"black", "brown", "red", "orange", "yellow",
		"green", "blue", "violet", "grey", "white",
		"gold", "silver", "pink",
	}

	for _, color := range expectedColors {
		if _, ok := resistor_band_mapping[color]; !ok {
			t.Errorf("expected resistor_band_mapping to contain %q", color)
		}
	}

	expectedAbbrevs := []string{
		"bk", "bn", "rd", "og", "ye",
		"gn", "bu", "vt", "gy", "wh",
		"gd", "sr", "pk",
	}

	for _, abbrev := range expectedAbbrevs {
		if _, ok := resistor_band_mapping[abbrev]; !ok {
			t.Errorf("expected resistor_band_mapping to contain abbreviation %q", abbrev)
		}
	}
}

func TestResistorToleranceValues(t *testing.T) {
	tests := []struct {
		color     string
		tolerance float64
	}{
		{"brown", 1},
		{"red", 2},
		{"green", .5},
		{"blue", .25},
		{"violet", .1},
		{"grey", .01},
		{"gold", 5},
		{"silver", 10},
	}

	for _, tt := range tests {
		t.Run(tt.color, func(t *testing.T) {
			band := resistor_band_mapping[tt.color]
			if band.tolerance == nil {
				t.Errorf("expected %s to have tolerance", tt.color)
				return
			}
			if band.tolerance.Value != tt.tolerance {
				t.Errorf("expected %s tolerance to be %v, got %v", tt.color, tt.tolerance, band.tolerance.Value)
			}
		})
	}
}

func TestResistorTempCoefficientValues(t *testing.T) {
	tests := []struct {
		color  string
		tempCE int
	}{
		{"black", 250},
		{"brown", 100},
		{"red", 50},
		{"orange", 15},
		{"yellow", 25},
		{"green", 20},
		{"blue", 10},
		{"violet", 5},
		{"grey", 1},
	}

	for _, tt := range tests {
		t.Run(tt.color, func(t *testing.T) {
			band := resistor_band_mapping[tt.color]
			if !band.is_valid_temp_ce_band {
				t.Errorf("expected %s to be valid temp coefficient band", tt.color)
			}
			if band.temp_ce != tt.tempCE {
				t.Errorf("expected %s temp_ce to be %v, got %v", tt.color, tt.tempCE, band.temp_ce)
			}
		})
	}
}

//endregion Resistance Tests
