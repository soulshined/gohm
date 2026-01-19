package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
	"math"
)

func cmd_ohmslaw_handler(cmd *cli.Command) string {
	voltage := utils.ParseShorthand(cmd.GetFlagValue("voltage"), abbrvs.VOLTAGE)
	current := utils.ParseShorthand(cmd.GetFlagValue("current"), abbrvs.CURRENT)
	power := utils.ParseShorthand(cmd.GetFlagValue("power"), abbrvs.POWER)
	resistance := ResistanceInSeries(cmd.GetFlagValues("resistance"))

	flags := []struct {
		name string
		char string
	}{
		{"voltage", "V"},
		{"current", "I"},
		{"resistance", "R"},
		{"power", "P"},
	}

	key := ""
	for _, f := range flags {
		if cmd.IsFlagSet(f.name) {
			key += f.char
		}
	}

	if len(key) < 2 {
		panic("too few arguments: -resistance | -current | -voltage | -power")
	} else if len(key) > 2 {
		panic("too many arguments: -resistance | -current | -voltage | -power")
	}

	switch key {
	case "VI":
		resistance = voltage / current
		power = voltage * current
	case "VR":
		current = voltage / resistance
		power = voltage * current
	case "VP":
		current = power / voltage
		resistance = voltage / current
	case "IR":
		voltage = current * resistance
		power = voltage * current
	case "IP":
		voltage = power / current
		resistance = voltage / current
	case "RP":
		current = math.Sqrt(power / resistance)
		voltage = current * resistance
	}

	switch cmd.GetFlagValue("format") {
	case "json":
		return fmt.Sprintf(`{"voltage":%s,"voltageAbbreviated":"%sV","current":%s,"currentAbbreviated":"%sA","resistance":%s,"resistanceAbbreviated":"%sΩ","power":%s,"powerAbbreviated":"%sW"}`,
			utils.FormatFloat(voltage),
			utils.GetAbbreviatedValue(voltage),
			utils.FormatFloat(current),
			utils.GetAbbreviatedValue(current),
			utils.FormatFloat(resistance),
			utils.GetAbbreviatedValue(resistance),
			utils.FormatFloat(power),
			utils.GetAbbreviatedValue(power),
		)
	case "raw":
		return fmt.Sprintf("voltage=%sV current=%sA resistance=%sΩ power=%sW",
			utils.FormatFloat(voltage),
			utils.FormatFloat(current),
			utils.FormatFloat(resistance),
			utils.FormatFloat(power),
		)
	default:
		return fmt.Sprintf("voltage=%sV current=%sA resistance=%sΩ power=%sW",
			utils.GetAbbreviatedValue(voltage),
			utils.GetAbbreviatedValue(current),
			utils.GetAbbreviatedValue(resistance),
			utils.GetAbbreviatedValue(power),
		)
	}
}
