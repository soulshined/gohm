package identify

import (
	"gohm/cli"
	"gohm/utils"
)

var default_tolerance = utils.Tolerance{
	Value:    20,
	UnitType: utils.UNIT_TYPE_PERCENT,
}

func GetCommand() *cli.Command {
	cmd := &cli.Command{
		Name:        "identify",
		Aliases:     []string{"id"},
		Description: "Identify electrical components by visual indicators",
	}

	cmd.AddSubcommand(get_command_capacitor())
	cmd.AddSubcommand(get_command_resistor())

	return cmd
}

func get_command_capacitor() *cli.Command {
	cmd := &cli.Command{
		Name:        "capacitor",
		Description: "Identify capacitor value",
		Handler:     cmd_capacitor_handler,
		Examples: []cli.Example{
			{
				Command:     "gohm identify capacitor -eiac 21",
				Description: "2 digit code",
				Output:      "nominal=21pF min=nil max=nil",
			},
			{
				Command:     "gohm identify capacitor -eiac A8",
				Description: "2 digit eia-198 code",
				Output:      "nominal=100μF min=nil max=nil",
			},
			{
				Command:     "gohm identify capacitor -eiac 100",
				Description: "3 digit code - SMD or ceramic",
				Output:      "nominal=10pF min=nil max=nil",
			},
			{
				Command:     "gohm identify capacitor -eiac 9R4",
				Description: "3 digit decimal",
				Output:      "nominal=9.4pF min=nil max=nil",
			},
			{
				Command: "gohm identify capacitor -eiac 541",
				Output:  "nominal=540pF min=nil max=nil",
			},
			{
				Command:     "gohm identify capacitor -eiac 123B",
				Description: "4 digit code",
				Output:      "nominal=12nF min=11.988zF max=11.988zF",
			},
			{
				Command:     "gohm identify capacitor -eiac 6R7K",
				Description: "4 digit decimal",
				Output:      "nominal=6.7pF min=6.029999999999999yF max=6.029999999999999yF",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:        "eiac",
		Aliases:     []string{"code"},
		Description: "The eia code or identifier - supports 2-4 digit codes including SMD & EIA-198",
	})
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}

func get_command_resistor() *cli.Command {
	cmd := &cli.Command{
		Name:        "resistor",
		Description: "Identify resistor value from color bands - color bands are n args passed in",
		Handler:     cmd_resistor_handler,
		Examples: []cli.Example{
			{
				Command: "gohm identify resistor red red brown",
				Output:  "\u001b[91m▌▌\u001b[38;5;172m▌\033[0m nominal=220Ω min=176Ω max=264Ω temp_coefficient=nil",
			},
			{
				Command:     "gohm identify resistor rd rd bn rd",
				Description: "EIA Shorthand",
				Output:      "\u001b[91m▌▌\u001b[38;5;172m▌ \u001b[91m▌\033[0m nominal=220Ω min=215.6Ω max=224.4Ω temp_coefficient=nil",
			},
		},
	}
	cmd.AddFlag(&cli.Flag{
		Name:           "format",
		Description:    "Output format",
		Default:        "abbr",
		PossibleValues: []string{"abbr", "raw", "json"},
	})
	return cmd
}
