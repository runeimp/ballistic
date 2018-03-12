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
type CountryCodesAndNumbers struct {
	CountryNames []string
	CountryAlpha2 string
	CountryAlpha3 string
	Adjective string
	SingularNoun string
	PluralNoun string
	Number_Separatrix string
	Decimal_Grouping []int
	Decimal_GroupDelimiters []string
	Fractional_Grouping []int
	Fractional_GroupDelimiters []string
}


//
// CONSTANTS
//
const SEPARATRIX_APOSTROPHE = "'"
const SEPARATRIX_BULLET = "•"
const SEPARATRIX_COMMA = ","
const SEPARATRIX_INTERPUNCT = "·" // AKA Decimal Point, Mid Dot, Point
const SEPARATRIX_POINT = "." // AKA Full Stop
const SEPARATRIX_SEMICOLON = ";"
const SEPARATRIX_SPACE = " "
const SEPARATRIX_UNDERSCORE = "_"
const SEPARATRIX_VBAR = "|"

/**
 * @see [MATHEMATICAL NOTATION COMPARISONS BETWEEN U.S. AND LATIN AMERICAN COUNTRIES]: http://www.csus.edu/indiv/o/oreyd/acp.htm_files/todos.operation.description.pdf
 * @see [ISO 31-0]: https://en.wikipedia.org/wiki/ISO_31-0#Numbers
 * @see [Indian Numbering System]: https://en.wikipedia.org/wiki/Indian_numbering_system
 * @see [SI]: https://en.wikipedia.org/wiki/International_System_of_Units
 */
const COUNTRY_CA_NUMBER_FORMAT_STD = "### ### ###(,)###"   // Canada
const COUNTRY_DE_NUMBER_FORMAT_STD = "### ###.###(,)###"   // Germany
const COUNTRY_DK_NUMBER_FORMAT_STD = "### ### ###(,)###"   // Denmark
const COUNTRY_ES_NUMBER_FORMAT_STD = "###.###.###(,)###"   // Spain
const COUNTRY_FI_NUMBER_FORMAT_STD = "### ### ###(,)###"   // Finland
const COUNTRY_FR_NUMBER_FORMAT_STD = "### ### ###(,)###"   // France
const COUNTRY_GB_NUMBER_FORMAT_OLD = "###?###?###(|)###"   // Great Britain. Old standard.
const COUNTRY_GB_NUMBER_FORMAT_STD = "###,###,###(.)###"   // Great Britain
const COUNTRY_IN_NUMBER_FORMAT_STD = "##,##,###(.)###"     // Italy
const COUNTRY_IT_NUMBER_FORMAT_STD = "###.###.###(,)###"   // Italy
const COUNTRY_MX_NUMBER_FORMAT_ALT = "###;###,###(.)###"   // Mexico Alternate
const COUNTRY_MX_NUMBER_FORMAT_STD = "###'###,###(.)###"   // Mexico
const COUNTRY_NO_NUMBER_FORMAT_STD = "###.###.###(,)###"   // Norway
const COUNTRY_SE_NUMBER_FORMAT_STD = "###.###.###(,)###"   // Sweden
const COUNTRY_TH_NUMBER_FORMAT_STD = "###,###,###(.)###"   // Thailand
const COUNTRY_UK_NUMBER_FORMAT_OLD = "###?###?###(|)###"   // United Kingdom (Great Britain). Old standard.
const COUNTRY_UK_NUMBER_FORMAT_STD = "###,###,###(.)###"   // United Kingdom (Great Britain)
const COUNTRY_US_NUMBER_FORMAT_STD = "###,###,###(.)###"   // United States
const COUNTRY_ZA_NUMBER_FORMAT_STD = "### ### ###(,)###"   // South Africa

const LANG_DE_NUMBER_FORMAT = "### ###.###(,)###"          // German
const LANG_DK_NUMBER_FORMAT = "### ### ###(,)###"          // Danish
const LANG_EN_NUMBER_FORMAT_CA = "### ### ###(,)###"       // Canadian English
const LANG_EN_NUMBER_FORMAT_GB = "###,###,###(.)###"       // British English
const LANG_EN_NUMBER_FORMAT_UK = "###,###,###(.)###"       // British English
const LANG_EN_NUMBER_FORMAT_US = "###,###,###(.)###"       // American English
const LANG_EO_NUMBER_FORMAT = "###?###?###(,)###"          // Esperanto
const LANG_ES_NUMBER_FORMAT_ES_STD = "###.###.###(,)###"   // Spanish
const LANG_ES_NUMBER_FORMAT_MX_ALT = "###;###,###(.)###"   // Spanish
const LANG_ES_NUMBER_FORMAT_MX_STD = "###'###,###(.)###"   // Spanish
const LANG_FI_NUMBER_FORMAT = "### ### ###(,)###"          // Finnish
const LANG_FR_NUMBER_FORMAT = "### ### ###(,)###"          // French
const LANG_IO_NUMBER_FORMAT = "###.###.###(,)###"          // Ido
const LANG_IT_NUMBER_FORMAT = "###.###.###(,)###"          // Italian
const LANG_NO_NUMBER_FORMAT = "###.###.###(,)###"          // Norwegian
const LANG_SE_NUMBER_FORMAT = "###.###.###(,)###"          // Swedish
const LANG_TH_NUMBER_FORMAT = "###,###,###(.)###"          // Thai

// COUNTRY_{ISO 3166-1 alpha-2 code}_NUMBER_SEPARATRIX
const COUNTRY_AU_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Australia
const COUNTRY_BD_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Bangladesh
const COUNTRY_BN_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Brunei Darussalam / CN: Brunei
const COUNTRY_BOT_NUMBER_SEPARATRIX = SEPARATRIX_POINT        // BOT (British Overseas Territories) / AKA: UKOT (United Kingdom Overseas Territories) / AKA: BWI (British West Indies); includes Akrotiri and Dhekelia, Anguilla, Bermuda, British Antarctic Territory, British Indian Ocean Territory, British Virgin Islands, Cayman Islands, Falkland Islands, Gibraltar, Montserrat, (Pitcairn, Henderson, Ducie and Oeno Islands), (Saint Helena, Ascension and Tristan da Cunha), (South Georgia and the South Sandwich Islands), (Turks and Caicos Islands)
const COUNTRY_BW_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Botswana
const COUNTRY_CA_EN_NUMBER_SEPARATRIX = SEPARATRIX_POINT      // Canada (English)
const COUNTRY_CA_EN_NUMBER_THOUSANDS = SEPARATRIX_SPACE       // Canada (English)
const COUNTRY_CA_FR_NUMBER_SEPARATRIX = SEPARATRIX_COMMA      // Canada (French)
const COUNTRY_CA_FR_NUMBER_THOUSANDS = SEPARATRIX_SPACE       // Canada (French)
const COUNTRY_CH_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Switzerland. Code taken from name in Latin: Confoederatio Helvetica
const COUNTRY_CN_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // China
const COUNTRY_DE_NUMBER_MILLIONS = SEPARATRIX_SPACE           // Germany
const COUNTRY_DE_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // Germany
const COUNTRY_DE_NUMBER_THOUSANDS = SEPARATRIX_POINT          // Germany
const COUNTRY_DO_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Dominican Republic
const COUNTRY_EG_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Egypt
const COUNTRY_ES_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // Spain
const COUNTRY_ES_NUMBER_THOUSANDS = SEPARATRIX_POINT          // Spain
const COUNTRY_FI_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // Finland
const COUNTRY_FI_NUMBER_THOUSANDS = SEPARATRIX_SPACE          // Finland
const COUNTRY_FR_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // France
const COUNTRY_FR_NUMBER_THOUSANDS = SEPARATRIX_SPACE          // France
const COUNTRY_GB_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // United Kingdom of Great Britain and Northern Ireland
const COUNTRY_GB_NUMBER_THOUSANDS = SEPARATRIX_COMMA          // United Kingdom of Great Britain and Northern Ireland
const COUNTRY_GH_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Ghana
const COUNTRY_GR_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // Greece
const COUNTRY_GT_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Guatemala
const COUNTRY_HK_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Hong Kong
const COUNTRY_HN_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Honduras
const COUNTRY_IE_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Ireland
const COUNTRY_IL_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Israel
const COUNTRY_IT_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // Italy
const COUNTRY_IT_NUMBER_THOUSANDS = SEPARATRIX_POINT          // Italy
const COUNTRY_JO_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Jordan
const COUNTRY_JP_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Japan
const COUNTRY_KE_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Kenya
const COUNTRY_KH_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Cambodia
const COUNTRY_KP_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Democratic People's Republic of Korea / CN: North Korea
const COUNTRY_KR_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Republic of Korea / CN: South Korea
const COUNTRY_LB_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Lebanon
const COUNTRY_LI_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Liechtenstein
const COUNTRY_LK_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Sri Lanka
const COUNTRY_LU_NUMBER_SEPARATRIX_ALT = SEPARATRIX_COMMA     // Luxembourg
const COUNTRY_LU_NUMBER_SEPARATRIX_STD = SEPARATRIX_POINT     // Luxembourg
const COUNTRY_MM_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Myanmar
const COUNTRY_MO_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Macao / Formerly: Macau
const COUNTRY_MT_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Malta
const COUNTRY_MV_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Maldives
const COUNTRY_MX_NUMBER_MILLION_ALT = SEPARATRIX_SEMICOLON    // Mexico
const COUNTRY_MX_NUMBER_MILLION_STD = SEPARATRIX_APOSTROPHE   // Mexico
const COUNTRY_MX_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Mexico
const COUNTRY_MX_NUMBER_THOUSAND = SEPARATRIX_COMMA           // Mexico
const COUNTRY_MY_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Malaysia
const COUNTRY_NA_NUMBER_SEPARATRIX_ALT = SEPARATRIX_COMMA     // Namibia
const COUNTRY_NA_NUMBER_SEPARATRIX_STD = SEPARATRIX_POINT     // Namibia
const COUNTRY_NG_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Nigeria
const COUNTRY_NI_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Nicaragua
const COUNTRY_NP_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Nepal
const COUNTRY_NZ_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // New Zealand
const COUNTRY_PA_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Panama
const COUNTRY_PH_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Philippines
const COUNTRY_PK_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Pakistan
const COUNTRY_PR_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Puerto Rico
const COUNTRY_PS_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // State of Palestine / CN: Palestine
const COUNTRY_SG_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Singapore
const COUNTRY_SV_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // El Salvador
const COUNTRY_TH_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Thailand
const COUNTRY_TH_NUMBER_THOUSANDS = SEPARATRIX_COMMA          // Thailand
const COUNTRY_TW_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // ROC (Republic of China) / Taiwan, Province of China / CN: Taiwan
const COUNTRY_TZ_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Tanzania
const COUNTRY_UG_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Uganda
const COUNTRY_UK_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // United Kingdom
const COUNTRY_UK_NUMBER_SEPARATRIX_OLD = SEPARATRIX_VBAR      // United Kingdom. Old standard.
const COUNTRY_UK_NUMBER_THOUSANDS = SEPARATRIX_COMMA          // United Kingdom
const COUNTRY_US_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // United States
const COUNTRY_US_NUMBER_THOUSANDS = SEPARATRIX_COMMA          // United States
const COUNTRY_ZA_NUMBER_SEPARATRIX_ALT = SEPARATRIX_POINT     // South Africa
const COUNTRY_ZA_NUMBER_SEPARATRIX_STD = SEPARATRIX_COMMA     // South Africa
const COUNTRY_ZW_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // Zimbabwe

const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____
const COUNTRY_XX_NUMBER_SEPARATRIX = SEPARATRIX_COMMA         // _____

const COUNTRY_IN_NUMBER_SEPARATRIX = SEPARATRIX_POINT         // India
const COUNTRY_IN_NUMBER_THOUSAND = SEPARATRIX_COMMA           // India
const COUNTRY_IN_NUMBER_HUNDRED_THOUSAND = SEPARATRIX_COMMA   // Indian lakh
const COUNTRY_IN_NUMBER_MYRIAD_TEN = SEPARATRIX_COMMA         // Indian lakh
const COUNTRY_IN_NUMBER_MYRIAD_THOUSAND = SEPARATRIX_COMMA    // Indian crore

const LANG_EO_NUMBER_SEPARATRIX = SEPARATRIX_COMMA   // Esperanto
const LANG_IA_NUMBER_SEPARATRIX = SEPARATRIX_COMMA   // Interlingua
const LANG_IO_NUMBER_SEPARATRIX = SEPARATRIX_COMMA   // Ido
const LANG_IO_NUMBER_THOUSANDS = SEPARATRIX_POINT    // Ido

const STANDARD_ISO_NUMBER_SEPARATRIX_ALT = SEPARATRIX_COMMA   // ISO 31-0
const STANDARD_ISO_NUMBER_SEPARATRIX_STD = SEPARATRIX_POINT   // ISO 31-0
const STANDARD_ISO_NUMBER_DELIMITER = SEPARATRIX_SPACE        // ISO 31-0
const STANDARD_SI_NUMBER_SEPARATRIX_ALT = SEPARATRIX_COMMA    // SI, Système international (d'unités), AKA International System of Units. Adopted by all countries except United States, Liberia, and Burma.
const STANDARD_SI_NUMBER_SEPARATRIX_STD = SEPARATRIX_POINT
const STANDARD_SI_NUMBER_DELIMITER = SEPARATRIX_SPACE


//
// VARIABLES
//
var LocaleData map[string]CountryCodesAndNumbers = map[string]CountryCodesAndNumbers{
	"DE": CountryCodesAndNumbers{
		CountryNames: []string{"Germany"},
		CountryAlpha2: "DE",
		CountryAlpha3: "DEU",
		Adjective: "German",
		SingularNoun: "German",
		PluralNoun: "Germans",
		Number_Separatrix: SEPARATRIX_COMMA,
		Decimal_Grouping: []int{3},
		Decimal_GroupDelimiters: []string{SEPARATRIX_POINT},
	},
	"IN": CountryCodesAndNumbers{
		CountryNames: []string{"India"},
		CountryAlpha2: "IN",
		CountryAlpha3: "IND",
		Adjective: "Indian",
		SingularNoun: "Indian",
		PluralNoun: "Indians",
		Number_Separatrix: COUNTRY_IN_NUMBER_SEPARATRIX,
		Decimal_Grouping: []int{3, 2},
		Decimal_GroupDelimiters: []string{COUNTRY_IN_NUMBER_THOUSAND, COUNTRY_IN_NUMBER_HUNDRED_THOUSAND},
	},
	"SI": CountryCodesAndNumbers{
		CountryNames: []string{"International System of Units", "Système international (d'unités)"},
		CountryAlpha2: "SI",
		CountryAlpha3: "SIU",
		Adjective: "International System of Units",
		SingularNoun: "International System of Units",
		PluralNoun: "International System of Units",
		Number_Separatrix: SEPARATRIX_POINT,
		Decimal_Grouping: []int{3},
		Decimal_GroupDelimiters: []string{SEPARATRIX_SPACE},
		Fractional_Grouping: []int{3},
		Fractional_GroupDelimiters: []string{SEPARATRIX_SPACE},
	},
	"TESTONE": CountryCodesAndNumbers{
		CountryNames: []string{"Test One"},
		Number_Separatrix: COUNTRY_IN_NUMBER_SEPARATRIX,
		Decimal_Grouping: []int{3, 2},
		Decimal_GroupDelimiters: []string{COUNTRY_IN_NUMBER_THOUSAND, COUNTRY_IN_NUMBER_HUNDRED_THOUSAND},
		Fractional_Grouping: []int{3, 2},
		Fractional_GroupDelimiters: []string{COUNTRY_IN_NUMBER_THOUSAND, COUNTRY_IN_NUMBER_HUNDRED_THOUSAND},
	},
	"US": CountryCodesAndNumbers{
		CountryNames: []string{"America", "The United States of America", "North America"},
		CountryAlpha2: "US",
		CountryAlpha3: "USA",
		Adjective: "American",
		SingularNoun: "American",
		PluralNoun: "Americans",
		Number_Separatrix: SEPARATRIX_POINT,
		Decimal_Grouping: []int{3},
		Decimal_GroupDelimiters: []string{SEPARATRIX_COMMA},
	},
}


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

		var locale_data CountryCodesAndNumbers

		var locale_alpha2 = "SI"
		if len(locale_match["country"]) > 1 {
			locale_alpha2 = locale_match["country"]
		} else if len(locale_match["lang"]) > 1 {
			locale_alpha2 = locale_match["lang"]
		}
		locale_data = LocaleData[locale_alpha2]
		// log.Printf("NumberFormatter func() | locale_alpha2: %s\n", locale_alpha2)

		separatrix := locale_data.Number_Separatrix
		grouping := locale_data.Decimal_Grouping
		delimiters := locale_data.Decimal_GroupDelimiters
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

		grouping = locale_data.Fractional_Grouping
		delimiters = locale_data.Fractional_GroupDelimiters
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

