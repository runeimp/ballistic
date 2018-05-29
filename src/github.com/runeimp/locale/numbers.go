/**
 * locale.numbers
 */

//
// PACKAGES
//
package locale


//
// IMPORTS
//
import (
 	// "fmt"
 	// "log"
	"regexp"
	"strconv"
	"strings"
)


//
// TYPES
//

type NumberFormatData struct {
	Separatrix string
	Decimal_Grouping []int
	Decimal_GroupMarks []string
	Fractional_Grouping []int
	Fractional_GroupMarks []string
}

type CountryCodesAndNumbers struct {
	CountryNames map[string]string
	CountryAlpha2 string
	CountryAlpha3 string
	Adjective []string
	SingularNoun []string
	PluralNoun []string
	// Separatrix string
	// Decimal_Grouping []int
	// Decimal_GroupMarks []string
	// Fractional_Grouping []int
	// Fractional_GroupMarks []string
	NumberFormat NumberFormatData
}

//
// CONSTANTS
//
const DELIMITER_APOSTROPHE = "'"
const DELIMITER_BULLET = "•"
const DELIMITER_COMMA = ","
const DELIMITER_INTERPUNCT = "·" // AKA Decimal Point, Mid Dot, Point
const DELIMITER_POINT = "." // AKA Full Stop
const DELIMITER_SEMICOLON = ";"
const DELIMITER_SPACE = " "
const DELIMITER_UNDERSCORE = "_"
const DELIMITER_VBAR = "|"

/**
 * @see [MATHEMATICAL NOTATION COMPARISONS BETWEEN U.S. AND LATIN AMERICAN COUNTRIES]: http://www.csus.edu/indiv/o/oreyd/acp.htm_files/todos.operation.description.pdf
 * @see [ISO 31-0]: https://en.wikipedia.org/wiki/ISO_31-0#Numbers
 * @see [Indian Numbering System]: https://en.wikipedia.org/wiki/Indian_numbering_system
 * @see [SI]: https://en.wikipedia.org/wiki/International_System_of_Units
 */




var /* const */ NUMBER_FORMAT_POINT_DECIMAL_COMMA3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_POINT,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_COMMA},
}

var /* const */ NUMBER_FORMAT_POINT_DECIMAL_SPACE3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_POINT,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_SPACE},
}

var /* const */ NUMBER_FORMAT_COMMA_DECIMAL_SPACE3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_COMMA,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_SPACE},
}

var /* const */ NUMBER_FORMAT_INTERPUNCT_DECIMAL_COMMA3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_INTERPUNCT,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_COMMA},
}

var /* const */ NUMBER_FORMAT_COMMA_DECIMAL_POINT3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_COMMA,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_POINT},
}

// Bangladesh, India (see Indian Numbering System)
var /* const */ NUMBER_FORMAT_POINT_DECIMAL_COMMA32 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_POINT,
	Decimal_Grouping: []int{3, 2},
	Decimal_GroupMarks: []string{DELIMITER_COMMA, DELIMITER_COMMA},
}

// Test Format Expanded from Indian Numbering System to include Fractional delimiters
var /* const */ NUMBER_FORMAT_POINT_DECIMAL_COMMA32_FRAC_COMMA32 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_POINT,
	Decimal_Grouping: []int{3, 2},
	Decimal_GroupMarks: []string{DELIMITER_COMMA, DELIMITER_COMMA},
	Fractional_Grouping: []int{3, 2},
	Fractional_GroupMarks: []string{DELIMITER_COMMA, DELIMITER_COMMA},
}

// Switzerland (computing), Liechtenstein
var /* const */ NUMBER_FORMAT_POINT_DECIMAL_APOSTROPHE3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_POINT,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_APOSTROPHE},
}

// Switzerland (handwriting)
var /* const */ NUMBER_FORMAT_COMMA_DECIMAL_APOSTROPHE3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_COMMA,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_APOSTROPHE},
}

// Spain (handwriting)
var /* const */ NUMBER_FORMAT_APOSTROPHE_DECIMAL_POINT3 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_APOSTROPHE,
	Decimal_Grouping: []int{3},
	Decimal_GroupMarks: []string{DELIMITER_POINT},
}

// China
var /* const */ NUMBER_FORMAT_POINT_DECIMAL_COMMA4 NumberFormatData = NumberFormatData{
	Separatrix: DELIMITER_POINT,
	Decimal_Grouping: []int{4}, // myriads
	Decimal_GroupMarks: []string{DELIMITER_COMMA},
}


//
// VARIABLES
//
var LocaleData map[string]CountryCodesAndNumbers = map[string]CountryCodesAndNumbers{
	"DE": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Germany"},
		CountryAlpha2: "DE",
		CountryAlpha3: "DEU",
		Adjective: []string{"German"},
		SingularNoun: []string{"German"},
		PluralNoun: []string{"Germans"},
		NumberFormat: NUMBER_FORMAT_COMMA_DECIMAL_POINT3,
	},
	"IN": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "India"},
		CountryAlpha2: "IN",
		CountryAlpha3: "IND",
		Adjective: []string{"Indian"},
		SingularNoun: []string{"Indian"},
		PluralNoun: []string{"Indians"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA32,
	},
	"SIU_EN": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "International System of Units", "FR": "Système international (d'unités)"},
		CountryAlpha2: "SI",
		CountryAlpha3: "SIU",
		Adjective: []string{"International System of Units"},
		SingularNoun: []string{"International System of Units"},
		PluralNoun: []string{"International System of Units"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_SPACE3,
	},
	"SIU_FR": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "International System of Units", "FR": "Système international (d'unités)"},
		CountryAlpha2: "SI",
		CountryAlpha3: "SIU",
		Adjective: []string{"Système international (d'unités)"},
		SingularNoun: []string{"Système international (d'unités)"},
		PluralNoun: []string{"Système international (d'unités)"},
		NumberFormat: NUMBER_FORMAT_COMMA_DECIMAL_SPACE3,
	},
	"TESTONE": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Test One"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA32_FRAC_COMMA32,
	},
	"EN": CountryCodesAndNumbers{
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"US": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "America", "Official": "The United States of America", "Continent": "North America"},
		CountryAlpha2: "US",
		CountryAlpha3: "USA",
		Adjective: []string{"American"},
		SingularNoun: []string{"American"},
		PluralNoun: []string{"Americans"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"AU": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Australia", "Official": "Australia", "UN": "Australia"},
		CountryAlpha2: "AU",
		CountryAlpha3: "AUS",
		Adjective: []string{"Australian"},
		SingularNoun: []string{"Australian"},
		PluralNoun: []string{"Australians"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"CA_EN": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Canada", "Official": "Canada", "UN": "Canada"},
		CountryAlpha2: "CA",
		CountryAlpha3: "CAN",
		Adjective: []string{"Canadian"},
		SingularNoun: []string{"Canadian"},
		PluralNoun: []string{"Canadians"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"CN": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "China", "Official": "People's Republic of China", "UN": "China"},
		CountryAlpha2: "CN",
		CountryAlpha3: "CHN",
		Adjective: []string{"Chinese"},
		SingularNoun: []string{"Chinese"},
		PluralNoun: []string{"Chinese"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"HK": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Hong Kong", "Official": "Hong Kong Special Administrative Region of the People's Republic of China"},
		CountryAlpha2: "HK",
		CountryAlpha3: "HKG",
		Adjective: []string{"Hongkonger", "Hong Kongese"},
		SingularNoun: []string{"Hongkonger", "Hong Kongese"},
		PluralNoun: []string{"Hongkongers", "Hong Kongese"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"IE": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Ireland", "Official": "Republic of Ireland", "UN": "Ireland"},
		CountryAlpha2: "IE",
		CountryAlpha3: "IRL",
		Adjective: []string{"Irish"},
		SingularNoun: []string{"Irishman", "Irishwoman"},
		PluralNoun: []string{"Irish"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"IL": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Israel", "Official": "State of Israel", "UN": "Israel"},
		CountryAlpha2: "IL",
		CountryAlpha3: "ISR",
		Adjective: []string{"Israeli"},
		SingularNoun: []string{"Israeli"},
		PluralNoun: []string{"Israelis"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	"JP": CountryCodesAndNumbers{
		CountryNames: map[string]string{"CN": "Japan", "Official": "State of Japan", "UN": "Japan"},
		CountryAlpha2: "JP",
		CountryAlpha3: "JPN",
		Adjective: []string{"Japanese"},
		SingularNoun: []string{"Japanese"},
		PluralNoun: []string{"Japanese"},
		NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	},
	// "__": CountryCodesAndNumbers{
	// 	CountryNames: map[string]string{"CN": "____", "Official": "____", "UN": "____"},
	// 	CountryAlpha2: "__",
	// 	CountryAlpha3: "___",
	// 	Adjective: []string{"____"},
	// 	SingularNoun: []string{"____"},
	// 	PluralNoun: []string{"____"},
	// 	NumberFormat: NUMBER_FORMAT_POINT_DECIMAL_COMMA3,
	// },
}

// Korea, Malaysia, Mexico, New Zealand, Pakistan, Philippines, Singapore, Taiwan, Thailand, United Kingdom, United States


var CountryNameByAlpha2 map[string]string = map[string]string{
	"CN": "China",
	"DE": "Germany",
	"US": "The United States of America",
}

var CountryNameByAlpha3 map[string]string = map[string]string{
	"CHN": "China",
	"DEU": "Germany",
	"USA": "The United States of America",
}

var CountryAlpha2ByAlpha3 map[string]string = map[string]string{
	"CHN": "CN",
	"DEU": "DE",
	"USA": "US",
}

var CountryAlpha3ByAlpha2 map[string]string = map[string]string{
	"CN": "CHN",
	"DE": "DEU",
	"US": "USA",
}


/*
country        Greece   Denmark  Spain     Turkey    Germany  Ireland
adjective      Grecian  Danish   Spanish   Turkish   German   Irish
singular noun  Greek    Dane     Spaniard  Turk      German   Irish(man/woman)
plural noun    Greeks   Danes    Spanish   Turks     Germans  Irish

ENV:LANG=en_US.UTF-8
ENV:LC_CTYPE=en_US.UTF-8

	  10,000 = Myriad (Greek value in American notation)
	1,00,000 = lakh (Indian)
	 100,000 = lakh (Indian value in American notation)
 1,00,00,000 = crore (Indian)
  10,000,000 = crore (Indian value in American notation)
   100000000 = 10,000 * 10,000
 100,000,000 = 1 hundred thousand US
10,00,00,000 = 10 crore or 1,000 lakh IN
*/


// var /* const */ LOCALE_RE = regexp.MustCompile("([a-z]{2}[_-][a-z]{2})\\.?.*")
var /* const */ LOCALE_RE = regexp.MustCompile("(?P<lang>[[:alpha:]]+)[_-]?(?P<country>[[:alpha:]]*)\\.?(?P<encoding>.*)")



//
// FUNCTIONS
//


/**
 * Regular Expression Submatch to Map
 *
 * @see https://stackoverflow.com/questions/20750843/using-named-matches-from-go-regex#answer-46202939
 * @see https://play.golang.org/p/zpLJe0iFwJ
 */
func reSubMatchMap(r *regexp.Regexp, str string) (map[string]string) {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	return subMatchMap
}


func NumberFormatter(locale_str string) func(number float64, scale int) (result string) {
	// locale_match := LOCALE_RE.FindStringSubmatch(locale_str)
	locale_match := reSubMatchMap(LOCALE_RE, locale_str)
	if len(locale_match["lang"]) > 0 {
		locale_match["lang"] = strings.ToUpper(locale_match["lang"])
	}
	// log.Printf("NumberFormatter() | locale_str: %s\n", locale_str)
	// log.Printf("NumberFormatter() | locale_match: %v\n", locale_match)
	// log.Printf("NumberFormatter() | locale_match[\"lang\"]: %v\n", locale_match["lang"])

	return func(number float64, scale int) (result string) {
		// log.Printf("NumberFormatter func() | number: %f | scale: %d\n", number, scale)

		var locale_alpha2 string = "SIU"
		var locale_data CountryCodesAndNumbers
		var locale_found bool
		var locale_normalized string = ""
		var locale_empty bool = false

		if len(locale_match["country"]) > 1 {
			locale_alpha2 = locale_match["country"]
			if len(locale_match["lang"]) > 1 {
				locale_normalized = locale_match["country"] + "_" + locale_match["lang"]
			}
		} else if len(locale_match["lang"]) > 1 {
			locale_alpha2 = locale_match["lang"]
		}

		if len(locale_normalized) > 0 {
			locale_data, locale_found = LocaleData[locale_normalized]
		}
		if ! locale_found || len(locale_data.NumberFormat.Separatrix) == 0 {
			locale_data, locale_found = LocaleData[locale_alpha2]
		}
		locale_empty = (locale_data.NumberFormat.Separatrix == "")

		if locale_empty || ! locale_found {
			locale_data, locale_found = LocaleData["EN"]
		}

		// log.Printf("NumberFormatter func() | locale_alpha2: %s\n", locale_alpha2)
		// log.Printf("NumberFormatter func() | locale_normalized: %s\n", locale_normalized)
		// log.Printf("NumberFormatter func() | empty: %v | locale_data: %v\n", (locale_data.NumberFormat.Separatrix == ""), locale_data)

		separatrix := locale_data.NumberFormat.Separatrix
		grouping := locale_data.NumberFormat.Decimal_Grouping
		delimiters := locale_data.NumberFormat.Decimal_GroupMarks
		// log.Printf("NumberFormatter func() | locale separatrix: %v\n", separatrix)
		// log.Printf("NumberFormatter func() | locale grouping: %v\n", grouping)
		// log.Printf("NumberFormatter func() | locale delimiters: %v\n", delimiters)

		// str_float := fmt.Sprintf("%.9f", number)
		str_float := strconv.FormatFloat(number, 'f', scale, 64)
		float_parts := strings.Split(str_float, ".")

		str_whole := float_parts[0] // decimal or integral
		str_scale := float_parts[1] // fractional

		// log.Printf("NumberFormatter func() | number: %f\n", number)
		// log.Printf("NumberFormatter func() | str_float: %s\n", str_float)
		// log.Printf("NumberFormatter func() | float_parts: %v\n", float_parts)
		// log.Printf("NumberFormatter func() | str_whole: %v\n", str_whole)
		// log.Printf("NumberFormatter func() | str_scale: %v\n", str_scale)


		var count int = 1
		var delimiter string = delimiters[0]
		var group int = 0
		var group_size int = grouping[0]
		var max_group int = len(grouping) - 1
		var num string = "0"

		// log.Printf("NumberFormatter func() | group: %d | max_group: %d | group_size: %d\n", group, max_group, group_size)

		if len(str_whole) > group_size {
			for i := len(str_whole) - 1; i >= 0; i-- {
				num = string(str_whole[i])
				// log.Printf("NumberFormatter func() | str_whole | i: %d | num: %v | group: %d | group_size: %d\n", i, num, group, group_size)
				result = num + result
				if count >= group_size && i != 0 {
					result = delimiter + result
					count = 1
					if group < max_group {
						group += 1
					}
					delimiter = delimiters[group]
					group_size = grouping[group]
				} else {
					count += 1
				}
			}
		} else {
			result = str_whole
		}

		if len(result) > 0 {
			result += separatrix
		}

		grouping = locale_data.NumberFormat.Fractional_Grouping
		delimiters = locale_data.NumberFormat.Fractional_GroupMarks
		if len(grouping) > 0 {
			// log.Printf("NumberFormatter func() | len(grouping) > 0: %d\n", len(grouping))
			
			count = 1
			delimiter = delimiters[0]
			delimiters_added := 0
			group = 0
			group_size = grouping[0]
			max_group = len(grouping) - 1
			scale_length := len(str_scale)
			// scale_length := scale
			result_scale := ""

			if scale < scale_length {
				scale_length = scale
			}

			max_length := scale_length - 1

			// log.Printf("NumberFormatter func() | scale_length: %d\n", scale_length)
			// log.Printf("NumberFormatter func() | max_length: %d\n", max_length)
			// log.Printf("NumberFormatter func() | group_size: %d\n", group_size)
			for i := 0; i < scale_length; i++ {
				num = string(str_scale[i])
				// log.Printf("NumberFormatter func() | str_scale | i: %d | num: %v | group: %d | group_size: %d\n", i, num, group, group_size)
				result_scale += num
				if count >= group_size && i < max_length {
					result_scale += delimiter
					delimiters_added += 1
					count = 1
					if group < max_group {
						group += 1
					}
					delimiter = delimiters[group]
					group_size = grouping[group]
				} else {
					count += 1
				}
				// log.Printf("NumberFormatter func() | count: %d\n", count)
			}
			// log.Printf("NumberFormatter func() | delimiters_added: %d\n", delimiters_added)

			if scale > scale_length {
				if delimiters_added == 0 {
					count = len(result_scale) + 1
					max_length = scale_length
				} else {
					count = len(result_scale) - delimiters_added
					max_length = scale_length + delimiters_added
				}
				delimiters_added = 0
				var i int = len(result_scale)

				for (i - delimiters_added) < scale {
					// log.Printf("NumberFormatter func() | scale > scale_length | count: %d | i: %d\n", count, i)
					// log.Printf("NumberFormatter func() | scale > scale_length | i: %d | count: %v | group: %d | group_size: %d | delimiters_added: %d\n", i, count, group, group_size, delimiters_added)

					result_scale += "0"
					if count >= group_size {
						result_scale += delimiter
						delimiters_added += 1
						count = 1
						if group < max_group {
							group += 1
						}
						delimiter = delimiters[group]
						group_size = grouping[group]
					} else {
						count += 1
					}
					i = len(result_scale)
				}
			}
			result += result_scale
		} else {
			result_scale := str_scale
			scale_length := len(result_scale)
			// log.Printf("NumberFormatter func() | scale: %d | scale_length: %d | result_scale: %s\n", scale, scale_length, result_scale)
			
			// Zero pad short fractional number
			for scale_length < scale {
				result_scale += "0"
				scale_length = len(result_scale)
				// log.Printf("NumberFormatter func() | scale: %d | scale_length: %d\n", scale, scale_length)
			}

			result += result_scale
		}

		// log.Printf("NumberFormatter func() | result: %v\n", result)
		
		return result
	}
}

/** Initialize Package */
func init() {
	// Nada
}

