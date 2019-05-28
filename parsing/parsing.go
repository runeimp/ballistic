package parsing

//
// IMPORTS
//
import (
	"log"
	"strconv"
	"strings"

	"github.com/runeimp/ballistic/app"
)

/*
 * Structs
 */

// InputUnits defines the values a user may input to the app
type InputUnits struct {
	Angle    string
	Length   string
	Mass     string
	Metric   bool
	Velocity string
	Weight   string
}

// ParsedData represents the parsed version of the input units
type ParsedData struct {
	Label     string
	Value     float64
	UserLabel string
	UserValue float64
}

/*
 * VARIABLES
 */
var (
	InputData   InputUnits
	outputDebug = false // NOTE: Temporary!!
)

// ParseValue will parse user input value and normalize it for internal use
func ParseValue(value, valueType string) (parsedData ParsedData) {
	// log.Printf("ParseValue()  <| value: %s | valueType: %s", value, valueType)

	if len(value) > 0 {
		valueMatch := app.VALUE_RE.FindStringSubmatch(value)

		number, _ := strconv.ParseFloat(valueMatch[1], 64)
		suffix := strings.ToLower(valueMatch[2])

		var designation string
		var normType string
		var normValue float64

		switch valueType {
		case app.ValueTypeAngle:
			normType = "degrees"

			switch suffix {
			case "degrees", "degree", "deg", "d":
				normValue = number * 1.0
				designation = app.AngleLabelDegrees
				// InputData.Metric = false
			case "radians", "radian", "rad", "r":
				normValue = number * 1.0
				designation = app.AngleLabelRadians
				// InputData.Metric = false
			}

			InputData.Angle = designation
		case app.ValueTypeLength:
			normType = "meter"

			switch suffix {
			case "feet", "foot", "ft", "f":
				normValue = number * app.LengthFromFeetToMeters
				designation = app.LengthLabelFoot
				InputData.Metric = false
			case "inches", "inch", "in", "i":
				normValue = number * app.LengthFromInchesToMeters
				designation = app.LengthLabelInch
				InputData.Metric = false
			case "nmi", "nm", "M":
				// Actual Values: M, NM, Nm, nm, nmi
				normValue = number * app.LengthFromNauticalMilesToMeters
				designation = app.LengthLabelNauticalMile
			case "yards", "yard", "yrd", "yd", "y":
				normValue = number * app.LengthFromYardsToMeters
				designation = app.LengthLabelYard
				InputData.Metric = false
			case "kilometers", "kilometer", "kilo", "km", "k":
				normValue = number * app.LengthFromKilometersToMeters
				designation = app.LengthLabelKilometer
				InputData.Metric = true
			case "meters", "m", "":
				normValue = number
				designation = app.LengthLabelMeter
				InputData.Metric = true
			case "centimeters", "centimeter", "centi", "cm", "c":
				normValue = number * app.LengthFromCentimetersToMeters
				designation = app.LengthLabelCentimeter
				InputData.Metric = true
			case "millimeters", "millimeter", "milli", "mm":
				normValue = number * app.LengthFromMillimetersToMeters
				designation = app.LengthLabelMillimeter
				InputData.Metric = true
			}

			InputData.Length = designation
		case app.ValueTypeMass:
			normType = "kilogram"

			switch suffix {
			case "grams", "g", "":
				normValue = number * app.MassFromGramsToKilograms
				designation = app.MassLabelGrams
				InputData.Metric = true
			case "grains", "gr":
				normValue = number * app.MassFromGrainsToKilograms
				designation = app.MassLabelGrains
				InputData.Metric = false
			case "pounds", "#", "lb", "lbs":
				normValue = number * app.MassFromPoundsToKilograms
				designation = app.MassLabelPounds
				InputData.Metric = false
			case "stone", "st":
				normValue = number * app.MassFromStoneToKilograms
				designation = app.MassLabelStone
				InputData.Metric = false
			case "ton":
				normValue = number * app.MassFromTonsShortToKilograms
				designation = app.MassLabelShortTon
				InputData.Metric = false
			case "lt":
				normValue = number * app.MassFromTonsLongToKilograms
				designation = app.MassLabelLongTon
				InputData.Metric = false
			case "mt":
				normValue = number * app.MassFromTonsMetricToKilograms
				designation = app.MassLabelMetricTonne
				InputData.Metric = true
			}

			InputData.Mass = designation
		case app.ValueTypeVelocity:
			normType = "meters per second"

			switch suffix {
			case "fps":
				normValue = number * app.VelocityFromFPSToMPS
				designation = app.VelocityLabelFPS
				InputData.Metric = false
			case "knots", "knot", "kn", "kt":
				normValue = number * app.VelocityFromKnotsToMPS
				designation = app.VelocityLabelKnots
			case "kmph", "k":
				normValue = number * app.VelocityFromKMPHToMPS
				designation = app.VelocityLabelKMPH
				InputData.Metric = true
			case "mph":
				normValue = number * app.VelocityFromMPHToMPS
				designation = app.VelocityLabelMPH
				InputData.Metric = false
			case "mps", "":
				normValue = number
				designation = app.VelocityLabelMPS
				InputData.Metric = true
			}

			InputData.Velocity = designation
		}

		if outputDebug {
			// log.Printf("ParseValue()   <| valueMatch: %s", valueMatch)
			// log.Printf("ParseValue()   <|      number: %f", number)
			// log.Printf("ParseValue()   <|      suffix: %s", suffix)
			// log.Printf("ParseValue()    | designation: %s", designation)
			// log.Printf("ParseValue()    |  normValue: %f", normValue)
			// log.Printf("ParseValue()    |   normType: %s", normType)

			// log.Printf("ParseValue()   <| %8s value: %17.6f %s (%s)", valueType, number, suffix, designation )
			// log.Printf("ParseValue()    | %8s normilazed: %12.6f %s", valueType, normValue, normType)

			suffixDesignation := suffix + " (" + designation + ")"
			log.Printf("ParseValue()    | %8s: %12.6f %-20s | %12.6f %s", valueType, number, suffixDesignation, normValue, normType)
		}

		parsedData.Label = normType
		parsedData.Value = normValue
		parsedData.UserLabel = designation
		parsedData.UserValue = number
	}

	return parsedData
}

/** Initialize Package */
func init() {
	// Nada
}
