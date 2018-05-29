/**
 * Ballistic.constants
 */

//
// PACKAGES
//
package ballistic


//
// IMPORTS
//
import (
	"regexp"
)


//
// CONSTANTS
//
const ENERGY_FROM_JOULES_TO_FOOTPOUNDS = 0.737562
const ENERGY_LABEL_FOOTPOUNDS = "foot-pounds"
const ENERGY_LABEL_JOULES = "joules"

const FORCE_FROM_KILOGRAMS_TO_NEWTONS float64 = 9.80665 // kg times meters per second squared
const FORCE_LABEL_NEWTONS string = "newtons"
const FORCE_LABEL_FOOTPOUNDS = "foot-pounds"
const GRAVITY_MPS float64 = 9.80665 // meters per second squared

const HELP_TEMPLATE = `
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


const ANGLE_DEGREES_TO_RADIANS float64 = 0.0174533
const ANGLE_LABEL_DEGREES = "degrees"
const ANGLE_LABEL_RADIANS = "radians"

const LENGTH_FROM_CENTIMETERS_TO_METERS float64 = 0.01
const LENGTH_FROM_FEET_TO_METERS float64 = 0.3048
const LENGTH_FROM_INCHES_TO_CM float64 = 2.54
const LENGTH_FROM_INCHES_TO_METERS float64 = 0.0254
const LENGTH_FROM_KILOMETERS_TO_METERS float64 = 1000.0
const LENGTH_FROM_METERS_TO_CENTIMETERS float64 = 100
const LENGTH_FROM_METERS_TO_FEET float64 = 3.28084
const LENGTH_FROM_METERS_TO_INCHES float64 = 39.3701
const LENGTH_FROM_METERS_TO_KILOMETERS float64 = 0.001
const LENGTH_FROM_METERS_TO_MILES float64 = 0.000621371
const LENGTH_FROM_METERS_TO_MILLIMETERS float64 = 1000
const LENGTH_FROM_METERS_TO_NAUTICAL_MILES float64 = 0.000539957
const LENGTH_FROM_MILES_TO_METERS float64 = 1609.34
const LENGTH_FROM_MILLIMETERS_TO_METERS float64 = 0.001
const LENGTH_FROM_NAUTICAL_MILES_TO_METERS float64 = 1852
const LENGTH_FROM_YARDS_TO_METERS float64 = 0.9144
const LENGTH_LABEL_CENTIMETER = "centimeters"
const LENGTH_LABEL_FOOT = "feet"
const LENGTH_LABEL_INCH = "inches"
const LENGTH_LABEL_KILOMETER = "kilometers"
const LENGTH_LABEL_METER = "meters"
const LENGTH_LABEL_MILE = "miles"
const LENGTH_LABEL_MILLIMETER = "millimeters"
const LENGTH_LABEL_NAUTICAL_MILE = "nautical miles"
const LENGTH_LABEL_YARD = "yards"

const MASS_FROM_GRAINS_TO_GRAMS float64 = 0.0647989
const MASS_FROM_GRAINS_TO_KILOGRAMS float64 = 0.0000647989
const MASS_FROM_GRAMS_TO_KILOGRAMS float64 = 0.001
const MASS_FROM_KILOGRAMS_TO_POUNDS float64 = 2.20462
const MASS_FROM_POUNDS_TO_GRAMS float64 = 453.592
const MASS_FROM_POUNDS_TO_KILOGRAMS float64 = 0.453592
const MASS_FROM_STONE_TO_GRAMS float64 = 6350.288
const MASS_FROM_STONE_TO_KILOGRAMS float64 = 6.350288
const MASS_FROM_STONE_TO_POUNDS float64 = 14.0
const MASS_FROM_TONS_LONG_TO_GRAMS float64 = 1016047.203454
const MASS_FROM_TONS_LONG_TO_KILOGRAMS float64 = 1016.047203454
const MASS_FROM_TONS_METRIC_TO_GRAMS float64 = 1000000.0
const MASS_FROM_TONS_METRIC_TO_KILOGRAMS float64 = 1000.0
const MASS_FROM_TONS_SHORT_TO_GRAMS float64 = 907185.0
const MASS_FROM_TONS_SHORT_TO_KILOGRAMS float64 = 907.185
const MASS_LABEL_GRAINS = "grains"
const MASS_LABEL_GRAMS = "grams"
const MASS_LABEL_LONG_TON = "long ton"
const MASS_LABEL_METRIC_TONNE = "metric tonne"
const MASS_LABEL_POUNDS = "pounds"
const MASS_LABEL_SHORT_TON = "short ton"
const MASS_LABEL_STONE = "stone"

const MOMENTUM_LABEL_FPS = "foot-pound per second"
const MOMENTUM_LABEL_MKS = "meter kilogram per second"
const MOMENTUM_LABEL_NS = "newton second"

// "Kælie"


const VELOCITY_FROM_FPS_TO_MPS float64 = 0.3048
const VELOCITY_FROM_KMPH_TO_MPS float64 = 0.277778
const VELOCITY_FROM_KNOTS_TO_KMPH float64 = 1.852
const VELOCITY_FROM_KNOTS_TO_MPS float64 = 0.514444
const VELOCITY_FROM_MPH_TO_MPS float64 = 0.44704
const VELOCITY_FROM_MPS_TO_FPS float64 = 3.28084
const VELOCITY_FROM_MPS_TO_KMPH float64 = 3.6
const VELOCITY_FROM_MPS_TO_KNOTS float64 = 1.94384
const VELOCITY_FROM_MPS_TO_MPH float64 = 2.23694
const VELOCITY_LABEL_FPS = "feet per second"
const VELOCITY_LABEL_KMPH = "kilometers per hour"
const VELOCITY_LABEL_KNOTS = "knots"
const VELOCITY_LABEL_MPH = "miles per hour"
const VELOCITY_LABEL_MPS = "meters per second"


var /* const */ VALUE_RE = regexp.MustCompile("([0-9]*[0-9.]?[0-9]*)([a-z#]*)")
// var /* const */ VALUE_RE = regexp.MustCompile("([0-9.]+)([a-z#]*)")
const VALUE_TYPE_ANGLE string = "angle"
const VALUE_TYPE_LENGTH string = "length"
const VALUE_TYPE_MASS string = "weight"
const VALUE_TYPE_VELOCITY string = "velocity"


/** Initialize Package */
func init() {
	// Nada
}

