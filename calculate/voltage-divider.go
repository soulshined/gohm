package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
	"math"
)

func cmd_voltage_divider_handler(cmd *cli.Command) string {
	supply_voltage := utils.ParseShorthand(cmd.GetFlagValue("voltage"), abbrvs.VOLTAGE)

	resistors := cmd.GetFlagValues("resistance")
	len_resistors := len(resistors)
	capacitors := cmd.GetFlagValues("capacitance")
	len_capacitors := len(capacitors)

	vout := 0.
	if len_resistors == 2 {
		r1 := utils.GetValueForRKMElseShorthand(resistors[0], abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
		r2 := utils.GetValueForRKMElseShorthand(resistors[1], abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
		vout = (r2 / (r1 + r2)) * supply_voltage
	} else if len_capacitors == 2 {
		c1 := utils.GetValueForRKMElseShorthand(capacitors[0], abbrvs.RKM_FARAD, abbrvs.FARAD)
		c2 := utils.GetValueForRKMElseShorthand(capacitors[1], abbrvs.RKM_FARAD, abbrvs.FARAD)
		vout = (c1 / (c1 + c2)) * supply_voltage
	} else if len_resistors == 1 && len_capacitors == 1 {
		r1 := utils.GetValueForRKMElseShorthand(resistors[0], abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
		c1 := utils.GetValueForRKMElseShorthand(capacitors[0], abbrvs.RKM_FARAD, abbrvs.FARAD)

		f := 1 / (2 * math.Pi * r1 * c1)
		if cmd.IsFlagSet("frequency") {
			f = utils.ParseShorthand(cmd.GetFlagValue("frequency"), abbrvs.FREQUENCY)
		}

		reactance := 1 / (2 * math.Pi * f * c1)
		vout = supply_voltage * (reactance / (r1 + reactance))
	} else {
		panic("unsupported: voltage divider type")
	}

	switch cmd.GetFlagValue("format") {
	case "json":
		return fmt.Sprintf(`{"voltage":%s,"voltageAbbreviated":"%sV"}`, utils.FormatFloat(vout), utils.GetAbbreviatedValue(vout))
	case "raw":
		return fmt.Sprintf("voltage=%sV", utils.FormatFloat(vout))
	default:
		return fmt.Sprintf("voltage=%sV", utils.GetAbbreviatedValue(vout))
	}
}
