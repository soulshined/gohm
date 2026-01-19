package calculate

import (
	"gohm/cli"
	"gohm/test_utils"
	"testing"
)

//region 555 Timer Tests

func TestCmd555HandlerMonostable(t *testing.T) {
	tests := []struct {
		name        string
		resistance  string
		capacitance string
		format      string
		contains    []string
	}{
		{
			name:        "monostable basic",
			resistance:  "10k",
			capacitance: "100μ",
			format:      "abbr",
			contains:    []string{"time=1.0999999999999999s"},
		},
		{
			name:        "monostable raw format",
			resistance:  "1k",
			capacitance: "1μ",
			format:      "raw",
			contains:    []string{"time=0.0010999999999999998s"},
		},
		{
			name:        "monostable json format",
			resistance:  "1k",
			capacitance: "1μ",
			format:      "json",
			contains:    []string{`"time":0.0010999999999999998`, `"timeAbbreviated":"1.0999999999999999ms"`},
		},
		{
			name:        "monostable RKM resistance",
			resistance:  "4K7",
			capacitance: "10n",
			format:      "abbr",
			contains:    []string{"time=51.7μs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_555_handler,
				map[string]string{
					"format":      tt.format,
					"capacitance": tt.capacitance,
				},
				map[string][]string{
					"resistance": {tt.resistance},
				},
				nil,
			)
			result := cmd_555_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmd555HandlerAstable(t *testing.T) {
	tests := []struct {
		name        string
		r1          string
		r2          string
		capacitance string
		format      string
		contains    []string
	}{
		{
			name:        "astable basic",
			r1:          "10k",
			r2:          "10k",
			capacitance: "100μ",
			format:      "abbr",
			contains:    []string{"time_low=692.9999999999998ms", "time_high=1.3859999999999997s", "frequency=480mHz"},
		},
		{
			name:        "astable raw format",
			r1:          "1k",
			r2:          "2k",
			capacitance: "1μ",
			format:      "raw",
			contains:    []string{"time_low=0.0013859999999999999s", "time_high=0.002079s", "frequency=288Hz"},
		},
		{
			name:        "astable json format",
			r1:          "1k",
			r2:          "2k",
			capacitance: "1μ",
			format:      "json",
			contains:    []string{`"timeLow":0.0013859999999999999`, `"timeHigh":0.002079`, `"frequency":288`, `"frequencyAbbreviated":"288Hz"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_555_handler,
				map[string]string{
					"format":      tt.format,
					"capacitance": tt.capacitance,
				},
				map[string][]string{
					"resistance": {tt.r1, tt.r2},
				},
				nil,
			)
			result := cmd_555_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmd555HandlerPanics(t *testing.T) {
	t.Run("too many resistances", func(t *testing.T) {
		cmd := test_utils.CreateTestCommand(
			cmd_555_handler,
			map[string]string{
				"format":      "abbr",
				"capacitance": "1μ",
			},
			map[string][]string{
				"resistance": {"1k", "2k", "3k"},
			},
			nil,
		)
		test_utils.ExpectPanic(t, "too many arguments: -resistance", func() {
			cmd_555_handler(cmd)
		})
	})
}

//endregion 555 Timer Tests

//region Capacitance Tests

func TestCmdCapacitanceHandler(t *testing.T) {
	tests := []struct {
		name     string
		circuit  string
		args     []string
		format   string
		contains []string
	}{
		{
			name:     "series capacitance",
			circuit:  "series",
			args:     []string{"100μ", "100μ"},
			format:   "abbr",
			contains: []string{"capacitance=50.00000000000001μF"},
		},
		{
			name:     "parallel capacitance",
			circuit:  "parallel",
			args:     []string{"100μ", "100μ"},
			format:   "abbr",
			contains: []string{"capacitance=200μF"},
		},
		{
			name:     "series raw format",
			circuit:  "series",
			args:     []string{"10n", "10n"},
			format:   "raw",
			contains: []string{"capacitance=0.000000005F"},
		},
		{
			name:     "series json format",
			circuit:  "series",
			args:     []string{"10n", "10n"},
			format:   "json",
			contains: []string{`"capacitance":0.000000005`, `"capacitanceAbbreviated":"5nF"`},
		},
		{
			name:     "RKM notation",
			circuit:  "series",
			args:     []string{"4n7", "10n"},
			format:   "abbr",
			contains: []string{"capacitance=3.1972789115646263nF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_capacitance_handler,
				map[string]string{
					"format":  tt.format,
					"circuit": tt.circuit,
				},
				nil,
				tt.args,
			)
			result := cmd_capacitance_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

//endregion Capacitance Tests

//region Current Divider Tests

func TestCmdCurrentDividerHandlerResistive(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		circuit  string
		args     []string
		format   string
		contains []string
	}{
		{
			name:     "resistive two resistors",
			current:  "1A",
			circuit:  "resistive",
			args:     []string{"1k", "1k"},
			format:   "abbr",
			contains: []string{"r1=1kΩ current=500mA", "r2=1kΩ current=500mA"},
		},
		{
			name:     "resistive json format",
			current:  "1A",
			circuit:  "resistive",
			args:     []string{"1k", "1k"},
			format:   "json",
			contains: []string{`"current":0.5`, `"currentAbbreviated":"500mA"`},
		},
		{
			name:     "resistive raw format",
			current:  "1A",
			circuit:  "resistive",
			args:     []string{"1k", "2k"},
			format:   "raw",
			contains: []string{"r1=1000Ω current=0.6666666666666666A", "r2=2000Ω current=0.3333333333333333A"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_current_divider_handler,
				map[string]string{
					"format":  tt.format,
					"current": tt.current,
					"circuit": tt.circuit,
				},
				nil,
				tt.args,
			)
			result := cmd_current_divider_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmdCurrentDividerHandlerCapacitive(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		args     []string
		format   string
		contains []string
	}{
		{
			name:     "capacitive two capacitors",
			current:  "1A",
			args:     []string{"100μ", "100μ"},
			format:   "abbr",
			contains: []string{"c1=100μF current=500mA", "c2=100μF current=500mA"},
		},
		{
			name:     "capacitive json format",
			current:  "1A",
			args:     []string{"100μ", "100μ"},
			format:   "json",
			contains: []string{`"current":0.5`, `"currentAbbreviated":"500mA"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_current_divider_handler,
				map[string]string{
					"format":  tt.format,
					"current": tt.current,
					"circuit": "capacitive",
				},
				nil,
				tt.args,
			)
			result := cmd_current_divider_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmdCurrentDividerHandlerPanics(t *testing.T) {
	t.Run("less than 2 values", func(t *testing.T) {
		cmd := test_utils.CreateTestCommand(
			cmd_current_divider_handler,
			map[string]string{
				"format":  "abbr",
				"current": "1A",
				"circuit": "resistive",
			},
			nil,
			[]string{"1k"},
		)
		test_utils.ExpectPanic(t, "too few arguments: [args...] - requires at least 2", func() {
			cmd_current_divider_handler(cmd)
		})
	})
}

//endregion Current Divider Tests

//region Missing Resistance Tests

func TestCmdMissingResistanceHandler(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		args     []string
		format   string
		contains []string
	}{
		{
			name:     "find missing resistance",
			target:   "500",
			args:     []string{"1k"},
			format:   "abbr",
			contains: []string{"resistance=1kΩ"},
		},
		{
			name:     "raw format",
			target:   "500",
			args:     []string{"1k"},
			format:   "raw",
			contains: []string{"resistance=1000Ω"},
		},
		{
			name:     "json format",
			target:   "500",
			args:     []string{"1k"},
			format:   "json",
			contains: []string{`"resistance":1000`, `"resistanceAbbreviated":"1kΩ"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_missing_resistance_handler,
				map[string]string{
					"format": tt.format,
					"target": tt.target,
				},
				nil,
				tt.args,
			)
			result := cmd_missing_resistance_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

//endregion Missing Resistance Tests

//region Ohm's Law Tests

func TestCmdOhmslawHandler(t *testing.T) {
	tests := []struct {
		name       string
		voltage    string
		current    string
		resistance []string
		power      string
		format     string
		contains   []string
	}{
		{
			name:     "voltage and current (VI)",
			voltage:  "12V",
			current:  "2A",
			format:   "abbr",
			contains: []string{"voltage=12V", "current=2A", "resistance=6Ω", "power=24W"},
		},
		{
			name:       "voltage and resistance (VR)",
			voltage:    "12V",
			resistance: []string{"6"},
			format:     "abbr",
			contains:   []string{"voltage=12V", "current=2A", "resistance=6Ω", "power=24W"},
		},
		{
			name:     "current and power (IP)",
			current:  "2A",
			power:    "24W",
			format:   "abbr",
			contains: []string{"voltage=12V", "current=2A"},
		},
		{
			name:       "resistance and power (RP)",
			resistance: []string{"6"},
			power:      "24W",
			format:     "abbr",
			contains:   []string{"voltage=12V", "current=2A"},
		},
		{
			name:    "json format",
			voltage: "12V",
			current: "2A",
			format:  "json",
			contains: []string{
				`"voltage":12`,
				`"current":2`,
				`"resistance":6`,
				`"power":24`,
			},
		},
		{
			name:    "raw format",
			voltage: "12V",
			current: "2A",
			format:  "raw",
			contains: []string{
				"voltage=12V",
				"current=2A",
				"resistance=6Ω",
				"power=24W",
			},
		},
		{
			name:       "shorthand values",
			voltage:    "1kV",
			resistance: []string{"1k"},
			format:     "abbr",
			contains:   []string{"voltage=1kV", "current=1A", "resistance=1kΩ", "power=1kW"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := map[string]string{
				"format":  tt.format,
				"voltage": tt.voltage,
				"current": tt.current,
				"power":   tt.power,
			}

			multiFlags := map[string][]string{}
			if len(tt.resistance) > 0 {
				multiFlags["resistance"] = tt.resistance
			}

			cmd := test_utils.CreateTestCommand(cmd_ohmslaw_handler, flags, multiFlags, nil)
			result := cmd_ohmslaw_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

func TestCmdOhmslawHandlerPanics(t *testing.T) {
	tests := []struct {
		name       string
		voltage    string
		current    string
		resistance []string
		power      string
		expected   string
	}{
		{
			name:     "only one value",
			voltage:  "12V",
			expected: "too few arguments: -resistance | -current | -voltage | -power",
		},
		{
			name:       "three values",
			voltage:    "12V",
			current:    "2A",
			resistance: []string{"6"},
			expected:   "too many arguments: -resistance | -current | -voltage | -power",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := map[string]string{
				"format":  "abbr",
				"voltage": tt.voltage,
				"current": tt.current,
				"power":   tt.power,
			}

			multiFlags := map[string][]string{}
			if len(tt.resistance) > 0 {
				multiFlags["resistance"] = tt.resistance
			}

			cmd := test_utils.CreateTestCommand(cmd_ohmslaw_handler, flags, multiFlags, nil)
			test_utils.ExpectPanic(t, tt.expected, func() {
				cmd_ohmslaw_handler(cmd)
			})
		})
	}
}

//endregion Ohm's Law Tests

//region Resistance Tests

func TestCmdResistanceHandler(t *testing.T) {
	tests := []struct {
		name     string
		circuit  string
		args     []string
		format   string
		contains []string
	}{
		{
			name:     "series resistance",
			circuit:  "series",
			args:     []string{"1k", "2k", "3k"},
			format:   "abbr",
			contains: []string{"resistance=6kΩ"},
		},
		{
			name:     "parallel resistance",
			circuit:  "parallel",
			args:     []string{"1k", "1k"},
			format:   "abbr",
			contains: []string{"resistance=500Ω"},
		},
		{
			name:     "series raw format",
			circuit:  "series",
			args:     []string{"100", "200"},
			format:   "raw",
			contains: []string{"resistance=300Ω"},
		},
		{
			name:     "series json format",
			circuit:  "series",
			args:     []string{"1k", "1k"},
			format:   "json",
			contains: []string{`"resistance":2000`, `"resistanceAbbreviated":"2kΩ"`},
		},
		{
			name:     "RKM notation",
			circuit:  "series",
			args:     []string{"4K7", "10K"},
			format:   "abbr",
			contains: []string{"resistance=", "kΩ"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := test_utils.CreateTestCommand(
				cmd_resistance_handler,
				map[string]string{
					"format":  tt.format,
					"circuit": tt.circuit,
				},
				nil,
				tt.args,
			)
			result := cmd_resistance_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

//endregion Resistance Tests

//region Voltage Divider Tests

func TestCmdVoltageDividerHandler(t *testing.T) {
	tests := []struct {
		name       string
		voltage    string
		resistors  []string
		capacitors []string
		format     string
		contains   []string
	}{
		{
			name:      "resistive divider",
			voltage:   "12V",
			resistors: []string{"1k", "1k"},
			format:    "abbr",
			contains:  []string{"voltage=6V"},
		},
		{
			name:       "capacitive divider",
			voltage:    "12V",
			capacitors: []string{"100μ", "100μ"},
			format:     "abbr",
			contains:   []string{"voltage=6V"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiFlags := map[string][]string{}
			if len(tt.resistors) > 0 {
				multiFlags["resistance"] = tt.resistors
			}
			if len(tt.capacitors) > 0 {
				multiFlags["capacitance"] = tt.capacitors
			}

			cmd := test_utils.CreateTestCommand(
				cmd_voltage_divider_handler,
				map[string]string{
					"format": tt.format,
				},
				multiFlags,
				nil,
			)

			// Set voltage flag separately since it needs special handling
			voltageFlag := &cli.Flag{
				Name:  "voltage",
				Value: tt.voltage,
				IsSet: tt.voltage != "",
			}
			cmd.Flags = append(cmd.Flags, voltageFlag)

			result := cmd_voltage_divider_handler(cmd)
			test_utils.AssertContains(t, result, tt.contains...)
		})
	}
}

//region Voltage Divider Tests

//region Helper Function Tests

func TestResistanceInSeries(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected float64
	}{
		{"two equal", []string{"1k", "1k"}, 2000},
		{"three values", []string{"100", "200", "300"}, 600},
		{"RKM notation", []string{"4K7", "10K"}, 14700},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ResistanceInSeries(tt.values)
			test_utils.AssertEquals(t, result, tt.expected)
		})
	}
}

func TestResistanceInParallel(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected float64
	}{
		{"two equal", []string{"1k", "1k"}, 500},
		{"two different", []string{"1k", "2k"}, 666.6666666666666},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ResistanceInParallel(tt.values)
			test_utils.AssertEquals(t, result, tt.expected)
		})
	}
}

func TestCapacitanceInSeries(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected float64
	}{
		{"two equal", []string{"100μ", "100μ"}, 0.00005},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CapacitanceInSeries(tt.values)
			test_utils.AssertEquals(t, result, tt.expected)
		})
	}
}

func TestCapacitanceInParallel(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected float64
	}{
		{"two equal", []string{"100μ", "100μ"}, 0.00019999999999999998},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CapacitanceInParallel(tt.values)
			test_utils.AssertEquals(t, result, tt.expected)
		})
	}
}

//endregion Helper Function Tests
