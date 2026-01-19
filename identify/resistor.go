package identify

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
	"math"
	"strconv"
	"strings"
)

type resistor_band struct {
	utils.ColorBand

	tolerance             *utils.Tolerance
	is_valid_temp_ce_band bool
	temp_ce               int
}

var resistor_band_mapping = map[string]*resistor_band{
	"black": {
		ColorBand:             utils.EIA_COLOR_MAPPING["black"],
		tolerance:             nil,
		is_valid_temp_ce_band: true,
		temp_ce:               250,
	},
	"brown": {
		ColorBand: utils.EIA_COLOR_MAPPING["brown"],
		tolerance: &utils.Tolerance{
			Value:    1,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               100,
	},
	"red": {
		ColorBand: utils.EIA_COLOR_MAPPING["red"],
		tolerance: &utils.Tolerance{
			Value:    2,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               50,
	},
	"orange": {
		ColorBand: utils.EIA_COLOR_MAPPING["orange"],
		tolerance: &utils.Tolerance{
			Value:    .05,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               15,
	},
	"yellow": {
		ColorBand: utils.EIA_COLOR_MAPPING["yellow"],
		tolerance: &utils.Tolerance{
			Value:    .02,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               25,
	},
	"green": {
		ColorBand: utils.EIA_COLOR_MAPPING["green"],
		tolerance: &utils.Tolerance{
			Value:    .5,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               20,
	},
	"blue": {
		ColorBand: utils.EIA_COLOR_MAPPING["blue"],
		tolerance: &utils.Tolerance{
			Value:    .25,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               10,
	},
	"violet": {
		ColorBand: utils.EIA_COLOR_MAPPING["violet"],
		tolerance: &utils.Tolerance{
			Value:    .1,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               5,
	},
	"grey": {
		ColorBand: utils.EIA_COLOR_MAPPING["grey"],
		tolerance: &utils.Tolerance{
			Value:    .01,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		is_valid_temp_ce_band: true,
		temp_ce:               1,
	},
	"white": {
		ColorBand: utils.EIA_COLOR_MAPPING["white"],
		tolerance: nil,
		temp_ce:   0,
	},
	"gold": {
		ColorBand: utils.EIA_COLOR_MAPPING["gold"],
		tolerance: &utils.Tolerance{
			Value:    5,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		temp_ce: 0,
	},
	"silver": {
		ColorBand: utils.EIA_COLOR_MAPPING["silver"],
		tolerance: &utils.Tolerance{
			Value:    10,
			UnitType: utils.UNIT_TYPE_PERCENT,
		},
		temp_ce: 0,
	},
	"pink": {
		ColorBand: utils.EIA_COLOR_MAPPING["pink"],
		tolerance: nil,
		temp_ce:   0,
	},
}

func init() {
	resistor_band_mapping[abbrvs.SI_PINK] = resistor_band_mapping["pink"]
	resistor_band_mapping[abbrvs.SI_SILVER] = resistor_band_mapping["silver"]
	resistor_band_mapping[abbrvs.SI_GOLD] = resistor_band_mapping["gold"]
	resistor_band_mapping[abbrvs.SI_BLACK] = resistor_band_mapping["black"]
	resistor_band_mapping[abbrvs.SI_BROWN] = resistor_band_mapping["brown"]
	resistor_band_mapping[abbrvs.SI_RED] = resistor_band_mapping["red"]
	resistor_band_mapping[abbrvs.SI_ORANGE] = resistor_band_mapping["orange"]
	resistor_band_mapping[abbrvs.SI_YELLOW] = resistor_band_mapping["yellow"]
	resistor_band_mapping[abbrvs.SI_GREEN] = resistor_band_mapping["green"]
	resistor_band_mapping[abbrvs.SI_BLUE] = resistor_band_mapping["blue"]
	resistor_band_mapping[abbrvs.SI_VIOLET] = resistor_band_mapping["violet"]
	resistor_band_mapping[abbrvs.SI_GREY] = resistor_band_mapping["grey"]
	resistor_band_mapping[abbrvs.SI_WHITE] = resistor_band_mapping["white"]
}

func cmd_resistor_handler(cmd *cli.Command) string {
	if cmd.ArgsLength > 6 {
		panic("too many arguments: [args...]")
	}

	significant_bands := cmd.Args[0:int(math.Ceil(float64(cmd.ArgsLength)/2.))]

	var bands_colors strings.Builder
	bands_visual := strings.Builder{}

	for i := range significant_bands {
		arg := significant_bands[i]
		key := strings.ToLower(arg)

		if key == "gold" || key == "silver" || key == "pink" || key == abbrvs.SI_GOLD || key == abbrvs.SI_SILVER || key == abbrvs.SI_PINK {
			panic(fmt.Errorf("invalid: significant digit band color can not be %s", arg))
		}

		val, ok := resistor_band_mapping[key]
		if !ok {
			panic(fmt.Errorf("invalid: significant digit band color %s", arg))
		}
		bands_colors.WriteString(strconv.Itoa(val.ColorBand.SignificantNumeral))
		bands_visual.WriteString(val.ColorBand.Ansi)
		bands_visual.WriteString("▌")
	}

	bands_value, err := strconv.ParseInt(bands_colors.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	multiplier_color := cmd.Args[2]
	tolerance_color := ""
	temp_ce_color := ""
	if len(cmd.Args) > 3 {
		tolerance_color = cmd.Args[3]
	}
	if cmd.ArgsLength > 4 {
		multiplier_color = cmd.Args[3]
		tolerance_color = cmd.Args[4]
	}
	if cmd.ArgsLength > 5 {
		temp_ce_color = cmd.Args[5]
	}

	multiplier := 0
	multiplier_band, ok := resistor_band_mapping[multiplier_color]
	if !ok {
		panic(fmt.Errorf("invalid: multiplier band color %s", multiplier_color))
	}

	multiplier = multiplier_band.ColorBand.Multiplier
	bands_visual.WriteString(multiplier_band.ColorBand.Ansi)
	bands_visual.WriteString("▌")

	tolerance := .2
	if tolerance_color != "" {
		tolerance_band, ok := resistor_band_mapping[tolerance_color]
		if !ok || tolerance_band.tolerance == nil {
			panic(fmt.Errorf("invalid: tolerance band color %s", tolerance_color))
		}

		tolerance = tolerance_band.tolerance.Value / 100

		bands_visual.WriteString(utils.ANSI_RESET)
		bands_visual.WriteString(" ")
		bands_visual.WriteString(tolerance_band.ColorBand.Ansi)
		bands_visual.WriteString("▌")
	}

	temp_ce := 0
	if temp_ce_color != "" {
		in_tempce, ok := resistor_band_mapping[temp_ce_color]
		if !ok || !in_tempce.is_valid_temp_ce_band {
			panic(fmt.Errorf("invalid: temperature coefficient band color %s", temp_ce_color))
		}
		temp_ce = in_tempce.temp_ce
		bands_visual.WriteString(in_tempce.ColorBand.Ansi)
		bands_visual.WriteString("▌")
	}

	nominal_value := float64(bands_value) * math.Pow10(multiplier)
	actual_min := nominal_value * (1 - tolerance)
	actual_max := nominal_value * (1 + tolerance)

	bands_visual.WriteString(utils.ANSI_RESET)

	switch cmd.GetFlagValue("format") {
	case "json":
		return fmt.Sprintf(`{"nominal":%s,"nominalAbbreviated":"%sΩ","actualMin":%s,"actualMinAbbreviated":"%sΩ","actualMax":%s,"actualMaxAbbreviated":"%sΩ","temperatureCoefficient":%s}`,
			utils.FormatFloat(nominal_value),
			utils.GetAbbreviatedValue(nominal_value),
			utils.FormatFloat(actual_min),
			utils.GetAbbreviatedValue(actual_min),
			utils.FormatFloat(actual_max),
			utils.GetAbbreviatedValue(actual_max),
			utils.If(temp_ce_color != "", strconv.Itoa(temp_ce), "null"),
		)
	case "raw":
		return fmt.Sprintf("%s nominal=%sΩ min=%sΩ max=%sΩ temp_coefficient=%s",
			&bands_visual,
			utils.FormatFloat(nominal_value),
			utils.FormatFloat(actual_min),
			utils.FormatFloat(actual_max),
			utils.If(temp_ce_color != "", fmt.Sprintf("%d ppm/K", temp_ce), "nil"),
		)
	default:
		return fmt.Sprintf("%s nominal=%sΩ min=%sΩ max=%sΩ temp_coefficient=%s",
			&bands_visual,
			utils.GetAbbreviatedValue(nominal_value),
			utils.GetAbbreviatedValue(actual_min),
			utils.GetAbbreviatedValue(actual_max),
			utils.If(temp_ce_color != "", fmt.Sprintf("%d ppm/K", temp_ce), "nil"),
		)
	}
}
