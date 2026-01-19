package calculate

import (
	"gohm/cli"
)

func GetCommand() *cli.Command {
	cmd := &cli.Command{
		Name:        "calculate",
		Aliases:     []string{"calc"},
		Description: "Perform electrical calculations",
	}

	cmd.AddSubcommand(get_command_555())
	cmd.AddSubcommand(get_command_capacitance())
	cmd.AddSubcommand(get_command_current_divider())
	cmd.AddSubcommand(get_command_missing_resistance())
	cmd.AddSubcommand(get_command_ohmslaw())
	cmd.AddSubcommand(get_command_resistance())
	cmd.AddSubcommand(get_command_voltage_divider())

	return cmd
}

func get_command_555() *cli.Command {
	cmd := &cli.Command{
		Name:        "555",
		Description: "Calculate 555 timer",
		Handler:     cmd_555_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate 555 -capacitance 10μF -resistance 100k",
				Output:  "time=1.1s",
			},
			{
				Command:     "gohm calculate 555 -capacitance 1μF -resistance 1k -resistance 1k",
				Description: "astable circuit example - providing 2 resistors",
				Output:      "time_low=693μs time_high=1.386ms frequency=480Hz",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:        "capacitance",
		Aliases:     []string{"c"},
		Description: "Capacitance value (F) - supports RKM & shorthand",
		Required:    true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "resistance",
		Aliases:     []string{"r"},
		Description: "Resistance value (R) - when specified 2 times - circuit is assumed astable - supports RKM & shorthand",
		IsMulti:     true,
		Required:    true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_capacitance() *cli.Command {
	cmd := &cli.Command{
		Name:        "capacitance",
		Description: "Calculate total capacitance of capacitors - values are n args passed in - args supports RKM & shorthand",
		Handler:     cmd_capacitance_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate capacitance 10μF 22μF",
				Output:  "capacitance=6.875μF",
			},
			{
				Command: "gohm calculate capacitance 10μF 22μF 18μF -circuit parallel",
				Output:  "50μF",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:           "circuit",
		Description:    "Type of circuit",
		Default:        "series",
		PossibleValues: []string{"series", "parallel"},
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_current_divider() *cli.Command {
	cmd := &cli.Command{
		Name:        "current-divider",
		Aliases:     []string{"cdiv"},
		Description: "Calculate current for parallel components - values are n args passed in",
		Handler:     cmd_current_divider_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate current-divider -current 6A 10 20 22",
				Output: `r1=10Ω current=3.0697674418604644A
      r2=20Ω current=1.5348837209302322A
      r3=22Ω current=1.3953488372093021A`,
			},
			{
				Command: "gohm calculate current-divider -current 6A 10pF 22pF 18pF",
				Output: `c1=10pF current=1.2000000000000002A
      c2=22pF current=2.64A
      c3=18pF current=2.16A`,
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:           "circuit",
		Description:    "Circuit divider type - resistive & capacitive args support RKM & shorthand",
		Default:        "resistive",
		PossibleValues: []string{"capacitive", "resistive"},
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "current",
		Aliases:     []string{"c", "i"},
		Description: "Input current - RKM & shorthand supported",
		Required:    true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_missing_resistance() *cli.Command {
	cmd := &cli.Command{
		Name:        "missing-resistance",
		Description: "Calculate a resistance value needed to complete a parallel circuit - resistors are n args passed in - RKM & shorthand supported",
		Handler:     cmd_missing_resistance_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate missing-resistance -target 150 250",
				Output:  "resistance=374.99999999999994Ω",
			},
			{
				Command: "gohm calculate missing-resistance -target 1k 2.5k 2k",
				Output:  "resistance=9.999999999999996kΩ",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:        "target",
		Aliases:     []string{"t"},
		Description: "Desired total/target resistance - RKM & shorthand supported",
		Required:    true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_ohmslaw() *cli.Command {
	cmd := &cli.Command{
		Name:        "ohmslaw",
		Aliases:     []string{"ohms"},
		Description: "Calculate Ohm's Law values (V=IR, P=IV, etc) based on 2 input values",
		Handler:     cmd_ohmslaw_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate ohmslaw -voltage 5V -current 5A",
				Output:  "voltage=5V current=5A resistance=1Ω power=25W",
			},
			{
				Command: "gohm calculate ohmslaw -voltage 5V -resistance 20",
				Output:  "voltage=5V current=250mA resistance=20Ω power=1.25W",
			},
			{
				Command:     "gohm calculate ohmslaw -voltage 5V -resistance 20 -resistance 10 -resistance 10",
				Description: "series total resistance 40Ω",
				Output:      "voltage=5V current=125mA resistance=40Ω power=625mW",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:        "current",
		Aliases:     []string{"c", "i"},
		Description: "Current value (I) - shorthand supported",
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "power",
		Aliases:     []string{"p"},
		Description: "Power value (W) - shorthand supported",
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "resistance",
		Aliases:     []string{"r"},
		Description: "Resistance value (R) - can be specified multiple times for series - RKM & shorthand supported",
		IsMulti:     true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "voltage",
		Aliases:     []string{"v"},
		Description: "Voltage value (V) - shorthand supported",
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_resistance() *cli.Command {
	cmd := &cli.Command{
		Name:        "resistance",
		Description: "Calculate total resistance of resistors - values are n args passed in - args supports RKM & shorthand",
		Handler:     cmd_resistance_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate resistance 5k 2.5k 5k",
				Output:  "resistance=12.5kΩ",
			},
			{
				Command: "gohm calculate resistance 5k 2.5k 5k -circuit parallel",
				Output:  "resistance=1.25kΩ",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:           "circuit",
		Description:    "Type of circuit",
		Default:        "series",
		PossibleValues: []string{"series", "parallel"},
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_voltage_divider() *cli.Command {
	cmd := &cli.Command{
		Name:        "voltage-divider",
		Aliases:     []string{"vdiv"},
		Description: "Calculate output voltage for series components",
		Handler:     cmd_voltage_divider_handler,
		Examples: []cli.Example{
			{
				Command: "gohm calculate voltage-divider -voltage 9v -resistance 3k -resistance 3k",
				Output:  "voltage=4.5V",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:        "capacitance",
		Aliases:     []string{"c"},
		Description: "RKM & shorthand supported",
		IsMulti:     true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "frequency",
		Aliases:     []string{"f"},
		Description: "used only with a resistor <-> capacitor divider type - shorthand supported",
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "resistance",
		Aliases:     []string{"r"},
		Description: "RKM & shorthand supported",
		IsMulti:     true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:        "voltage",
		Aliases:     []string{"v"},
		Description: "Input voltage - shorthand supported",
		Required:    true,
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}
