package app

//
// IMPORTS
//
import (
	"regexp"
)

/*
 * CONSTANTS
 */
const (
	EnergyFromJoulesToFootPounds = 0.737562
	EnergyLabelFootPounds        = "foot-pounds"
	EnergyLabelJoules            = "joules"

	ForceFromKilogramsToNewtons float64 = 9.80665 // kg times meters per second squared
	ForceLabelNewtons                   = "newtons"
	ForceLabelFootpounds                = "foot-pounds"
	GravityMPS                  float64 = 9.80665 // meters per second squared
)

const HelpTemplate = `
NAME:
   {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

GLOBAL OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{.Copyright}}{{end}}

VALUE SUFFIXES:
  All input values may be suffixed to allow for broader input selection.

  ANGLE
    d, deg, degree, degrees
    r, rad, radian, radians
  LENGTH
    c, cm, centi, centimeter, centimeters
    f, ft, foot, feet
    i, in, inch, inches
    k, km, kilo, kilometer, kilometers
    m, meters †
    M, NM, Nm, nm, nmi  (Nautical Miles)
    mm, milli, millimeter, millimeters
    y, yd, yrd, yard, yards
  MASS
    #, lb, lbs, pound, pounds
    g, gram, grams
    gr, grain, grains
    kg, kilo, kilogram, kilograms †
    lt, long-ton
    mt, tonne, metric-tonne
    st, stone
    t, ton, short-ton
  VELOCITY
    fps, feet-per-second
    kmph, kilometers-per-hour
    kn, kt, knot, knots
    mph, miles-per-hour
    mps, meters-per-second †

†  This is the default and will be used if no suffix is specified

If most or all of the input values are in imperial units then the output will use imperial units as well.

`

// ‡  This is the default if you set BALLISTIC_UNITS to imperial instead of metric
// The environment variable BALLISTIC_UNITS can be defined as imperial or metric. If it is defined the output units will always be of that system. If BALLISTIC_UNITS is not defined and most or all of the input values are in imperial units then the output will use imperial units as well. Otherwise ballistic defaults to metric.

const (
	AngleDegreesToRadians float64 = 0.0174533
	AngleLabelDegrees             = "degrees"
	AngleLabelRadians             = "radians"

	LengthFromCentimetersToMeters   float64 = 0.01
	LengthFromFeetToMeters          float64 = 0.3048
	LengthFromInchesToCM            float64 = 2.54
	LengthFromInchesToMeters        float64 = 0.0254
	LengthFromKilometersToMeters    float64 = 1000.0
	LengthFromMetersToCentimeters   float64 = 100
	LengthFromMetersToFeet          float64 = 3.28084
	LengthFromMetersToInches        float64 = 39.3701
	LengthFromMetersToKilometers    float64 = 0.001
	LengthFromMetersToMiles         float64 = 0.000621371
	LengthFromMetersToMillimeters   float64 = 1000
	LengthFromMetersToNauticalMiles float64 = 0.000539957
	LengthFromMilesToMeters         float64 = 1609.34
	LengthFromMillimetersToMeters   float64 = 0.001
	LengthFromNauticalMilesToMeters float64 = 1852
	LengthFromYardsToMeters         float64 = 0.9144
	LengthLabelCentimeter                   = "centimeters"
	LengthLabelFoot                         = "feet"
	LengthLabelInch                         = "inches"
	LengthLabelKilometer                    = "kilometers"
	LengthLabelMeter                        = "meters"
	LengthLabelMile                         = "miles"
	LengthLabelMillimeter                   = "millimeters"
	LengthLabelNauticalMile                 = "nautical miles"
	LengthLabelYard                         = "yards"

	MassFromGrainsToGrams         float64 = 0.0647989
	MassFromGrainsToKilograms     float64 = 0.0000647989
	MassFromGramsToKilograms      float64 = 0.001
	MassFromKilogramsToPounds     float64 = 2.20462
	MassFromPoundsToGrams         float64 = 453.592
	MassFromPoundsToKilograms     float64 = 0.453592
	MassFromStoneToGrams          float64 = 6350.288
	MassFromStoneToKilograms      float64 = 6.350288
	MassFromStoneToPounds         float64 = 14.0
	MassFromTonsLongToGrams       float64 = 1016047.203454
	MassFromTonsLongToKilograms   float64 = 1016.047203454
	MassFromTonsMetricToGrams     float64 = 1000000.0
	MassFromTonsMetricToKilograms float64 = 1000.0
	MassFromTonsShortToGrams      float64 = 907185.0
	MassFromTonsShortToKilograms  float64 = 907.185
	MassLabelGrains                       = "grains"
	MassLabelGrams                        = "grams"
	MassLabelLongTon                      = "long ton"
	MassLabelMetricTonne                  = "metric tonne"
	MassLabelPounds                       = "pounds"
	MassLabelShortTon                     = "short ton"
	MassLabelStone                        = "stone"

	MomentumLabelFPS = "foot-pound per second"
	MomentumLabelMKS = "meter kilogram per second"
	MomentumLabelNS  = "newton second"
)

const (
	VelocityFromFPSToMPS    float64 = 0.3048
	VelocityFromKMPHToMPS   float64 = 0.277778
	VelocityFromKnotsToKMPH float64 = 1.852
	VelocityFromKnotsToMPS  float64 = 0.514444
	VelocityFromMPHToMPS    float64 = 0.44704
	VelocityFromMPSToFPS    float64 = 3.28084
	VelocityFromMPSToKMPH   float64 = 3.6
	VelocityFromMPSToKnots  float64 = 1.94384
	VelocityFromMPSToMPH    float64 = 2.23694
	VelocityLabelFPS                = "feet per second"
	VelocityLabelKMPH               = "kilometers per hour"
	VelocityLabelKnots              = "knots"
	VelocityLabelMPH                = "miles per hour"
	VelocityLabelMPS                = "meters per second"
)

var /* const */ VALUE_RE = regexp.MustCompile("([0-9]*[0-9.]?[0-9]*)([a-z#]*)")

const (
	ValueTypeAngle    string = "angle"
	ValueTypeLength   string = "length"
	ValueTypeMass     string = "weight"
	ValueTypeVelocity string = "velocity"
)
