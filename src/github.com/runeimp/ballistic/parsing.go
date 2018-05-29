/**
 * Ballistic.parsing
 */

//
// PACKAGES
//
package ballistic


//
// IMPORTS
//
import (
	// . "github.com/runeimp/ballistic" // Import ballistic into this namespace for constants, etc.
	"log"
	"strconv"
	"strings"
)


//
// Structs
//
type InputUnits struct {
	Angle string
	Length string
	Mass string
	Metric bool
	Velocity string
	Weight string
}

type ParsedData struct {
	Label string
	Value float64
	UserLabel string
	UserValue float64
}


//
// VARIABLES
//
var InputData InputUnits
var output_debug bool = false // NOTE: Temporary!!


/** Parse user input value and normalize it for internal use */
func ParseValue(value, value_type string) (parsed_data ParsedData) {
	// log.Printf("ParseValue()  <| value: %s | value_type: %s", value, value_type)

	if len(value) > 0 {
		value_match := VALUE_RE.FindStringSubmatch(value)

		number, _ := strconv.ParseFloat(value_match[1], 64)
		suffix := strings.ToLower(value_match[2])

		var designation string
		var norm_type string
		var norm_value float64

		switch value_type {
		case VALUE_TYPE_ANGLE:
			norm_type = "degrees"

			switch suffix {
			case "degrees", "degree", "deg", "d":
				norm_value = number * 1.0
				designation = ANGLE_LABEL_DEGREES
				// InputData.Metric = false
			case "radians", "radian", "rad", "r":
				norm_value = number * 1.0
				designation = ANGLE_LABEL_RADIANS
				// InputData.Metric = false
			}

			InputData.Angle = designation
		case VALUE_TYPE_LENGTH:
			norm_type = "meter"

			switch suffix {
			case "feet", "foot", "ft", "f":
				norm_value = number * LENGTH_FROM_FEET_TO_METERS
				designation = LENGTH_LABEL_FOOT
				InputData.Metric = false
			case "inches", "inch", "in", "i":
				norm_value = number * LENGTH_FROM_INCHES_TO_METERS
				designation = LENGTH_LABEL_INCH
				InputData.Metric = false
			case "nmi", "nm", "M":
				// Actual Values: M, NM, Nm, nm, nmi
				norm_value = number * LENGTH_FROM_NAUTICAL_MILES_TO_METERS
				designation = LENGTH_LABEL_NAUTICAL_MILE
			case "yards", "yard", "yrd", "yd", "y":
				norm_value = number * LENGTH_FROM_YARDS_TO_METERS
				designation = LENGTH_LABEL_YARD
				InputData.Metric = false
			case "kilometers", "kilometer", "kilo", "km", "k":
				norm_value = number * LENGTH_FROM_KILOMETERS_TO_METERS
				designation = LENGTH_LABEL_KILOMETER
				InputData.Metric = true
			case "meters", "m", "":
				norm_value = number
				designation = LENGTH_LABEL_METER
				InputData.Metric = true
			case "centimeters", "centimeter", "centi", "cm", "c":
				norm_value = number * LENGTH_FROM_CENTIMETERS_TO_METERS
				designation = LENGTH_LABEL_CENTIMETER
				InputData.Metric = true
			case "millimeters", "millimeter", "milli", "mm":
				norm_value = number * LENGTH_FROM_MILLIMETERS_TO_METERS
				designation = LENGTH_LABEL_MILLIMETER
				InputData.Metric = true
			}

			InputData.Length = designation
		case VALUE_TYPE_MASS:
			norm_type = "kilogram"

			switch suffix {
			case "grams", "g", "":
				norm_value = number * MASS_FROM_GRAMS_TO_KILOGRAMS
				designation = MASS_LABEL_GRAMS
				InputData.Metric = true
			case "grains", "gr":
				norm_value = number * MASS_FROM_GRAINS_TO_KILOGRAMS
				designation = MASS_LABEL_GRAINS
				InputData.Metric = false
			case "pounds", "#", "lb", "lbs":
				norm_value = number * MASS_FROM_POUNDS_TO_KILOGRAMS
				designation = MASS_LABEL_POUNDS
				InputData.Metric = false
			case "stone", "st":
				norm_value = number * MASS_FROM_STONE_TO_KILOGRAMS
				designation = MASS_LABEL_STONE
				InputData.Metric = false
			case "ton":
				norm_value = number * MASS_FROM_TONS_SHORT_TO_KILOGRAMS
				designation = MASS_LABEL_SHORT_TON
				InputData.Metric = false
			case "lt":
				norm_value = number * MASS_FROM_TONS_LONG_TO_KILOGRAMS
				designation = MASS_LABEL_LONG_TON
				InputData.Metric = false
			case "mt":
				norm_value = number * MASS_FROM_TONS_METRIC_TO_KILOGRAMS
				designation = MASS_LABEL_METRIC_TONNE
				InputData.Metric = true
			}

			InputData.Mass = designation
		case VALUE_TYPE_VELOCITY:
			norm_type = "meters per second"

			switch suffix {
			case "fps":
				norm_value = number * VELOCITY_FROM_FPS_TO_MPS
				designation = VELOCITY_LABEL_FPS
				InputData.Metric = false
			case "knots", "knot", "kn", "kt":
				norm_value = number * VELOCITY_FROM_KNOTS_TO_MPS
				designation = VELOCITY_LABEL_KNOTS
			case "kmph", "k":
				norm_value = number * VELOCITY_FROM_KMPH_TO_MPS
				designation = VELOCITY_LABEL_KMPH
				InputData.Metric = true
			case "mph":
				norm_value = number * VELOCITY_FROM_MPH_TO_MPS
				designation = VELOCITY_LABEL_MPH
				InputData.Metric = false
			case "mps", "":
				norm_value = number
				designation = VELOCITY_LABEL_MPS
				InputData.Metric = true
			}

			InputData.Velocity = designation
		}

		if output_debug {
			// log.Printf("ParseValue()   <| value_match: %s", value_match)
			// log.Printf("ParseValue()   <|      number: %f", number)
			// log.Printf("ParseValue()   <|      suffix: %s", suffix)
			// log.Printf("ParseValue()    | designation: %s", designation)
			// log.Printf("ParseValue()    |  norm_value: %f", norm_value)
			// log.Printf("ParseValue()    |   norm_type: %s", norm_type)

			// log.Printf("ParseValue()   <| %8s value: %17.6f %s (%s)", value_type, number, suffix, designation )
			// log.Printf("ParseValue()    | %8s normilazed: %12.6f %s", value_type, norm_value, norm_type)

			suffix_designation := suffix + " (" + designation + ")"
			log.Printf("ParseValue()    | %8s: %12.6f %-20s | %12.6f %s", value_type, number, suffix_designation, norm_value, norm_type)
		}

		parsed_data.Label = norm_type
		parsed_data.Value = norm_value
		parsed_data.UserLabel = designation
		parsed_data.UserValue = number
	}

	return parsed_data
}


/** Initialize Package */
func init() {
	// Nada
}

