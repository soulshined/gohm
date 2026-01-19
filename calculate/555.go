package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
)

func cmd_555_handler(cmd *cli.Command) string {
	format := cmd.GetFlagValue("format")

	in_resistances := cmd.GetFlagValues("resistance")
	len_in_resistances := len(in_resistances)
	capacitance := utils.GetValueForRKMElseShorthand(cmd.GetFlagValue("capacitance"), abbrvs.RKM_FARAD, abbrvs.FARAD)

	if len_in_resistances == 2 {
		return cmd_555_handler_astable(format, in_resistances, capacitance)
	} else if len_in_resistances > 2 {
		panic("too many arguments: -resistance")
	}

	return cmd_555_handler_monostable(format, in_resistances[0], capacitance)
}

func cmd_555_handler_monostable(format string, in_resistance string, capacitance float64) string {
	resistance := utils.GetValueForRKMElseShorthand(in_resistance, abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
	time := 1.1 * resistance * capacitance

	switch format {
	case "json":
		return fmt.Sprintf(`{"time":%s,"timeAbbreviated":"%ss"}`, utils.FormatFloat(time), utils.GetAbbreviatedValue(time))
	case "raw":
		return fmt.Sprintf("time=%ss", utils.FormatFloat(time))
	default:
		return fmt.Sprintf("time=%ss", utils.GetAbbreviatedValue(time))
	}
}

func cmd_555_handler_astable(format string, resistances []string, capacitance float64) string {
	r1 := utils.GetValueForRKMElseShorthand(resistances[0], abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
	r2 := utils.GetValueForRKMElseShorthand(resistances[1], abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)

	th := 0.693 * (r1 + r2) * capacitance
	tl := 0.693 * r2 * capacitance
	f := 1.44 / ((r1 + 2*r2) * capacitance)

	switch format {
	case "json":
		return fmt.Sprintf(`{"timeLow":%s,"timeLowAbbreviated":"%ss","timeHigh":%s,"timeHighAbbreviated":"%ss","frequency":%s,"frequencyAbbreviated":"%sHz"}`,
			utils.FormatFloat(tl),
			utils.GetAbbreviatedValue(tl),
			utils.FormatFloat(th),
			utils.GetAbbreviatedValue(th),
			utils.FormatFloat(f),
			utils.GetAbbreviatedValue(f),
		)
	case "raw":
		return fmt.Sprintf("time_low=%ss time_high=%ss frequency=%sHz", utils.FormatFloat(tl), utils.FormatFloat(th), utils.FormatFloat(f))
	default:
		return fmt.Sprintf("time_low=%ss time_high=%ss frequency=%sHz", utils.GetAbbreviatedValue(tl), utils.GetAbbreviatedValue(th), utils.GetAbbreviatedValue(f))
	}
}
