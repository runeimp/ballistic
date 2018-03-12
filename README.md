Ballistic v0.4.1
================

Command line ballistics calculator


Features
--------

- Calculate:
	- Projectile energy
	- Projectile momentum
	- Projectile velocity
	- MPBR: Maximum Point Blank Range or Battle Zero is a military term refering the maximum distance a weapon can be fired to hit the torso of a human target (roughly 18&times;9 inches) every time (baring extreme weather or cover conditions) when aiming at the center of mass.
- Output
	- Human formated for interactive usage
	- JSON formated for easy scripting


Usage
-----

### Typical Gun Ballistics

```text
$ ballistic -m 300gr -v 800fps -f 4

  Projectile Velocity:             800.0000 feet per second
    Projectile Energy:           1,155.8428 joules
  Projectile Momentum:               4.7402 meter kilogram per second
Max Point Blank Range:             242.3544 feet

```

### Archery or Mechanical Ballistics with JSON output (pretty printed)

```text
$ ballistic --mass 42g --draw-weight 80lb --draw-length 0.72m --json --pretty
{
    "energy": {
        "label": "joules",
        "value": 128.10867801984
    },
    "momentum": {
        "label": "meter kilogram per second",
        "value": 2.319604379378794
    },
    "mpbr": {
        "label": "meters",
        "value": 16.731140499999988
    },
    "velocity": {
        "label": "meters per second",
        "value": 55.22867569949509
    }
}
```

### Help Info

```text
$ ballistic

NAME:
   Ballistic - Calculates what it can based on provided input.

USAGE:
   ballistic [global options]

VERSION:
   0.2.0

GLOBAL OPTIONS:
   --debug, -d                                       Output debug info
   --draw-length LENGTH, --length LENGTH, -l LENGTH  Bow or sling shot draw LENGTH. Used to calculate projectile velocity, energy, etc.
   --draw-weight WEIGHT, --weight WEIGHT, -w WEIGHT  Bow or sling shot draw WEIGHT. Used to calculate projectile velocity, energy, etc.
   --json, -j                                        Output JSON data
   --mass MASS, -m MASS                              Projectile MASS (weight). Used to calculate projectile velocity, energy, etc.
   --pretty-print, --pretty, -p                      Pretty printed JSON output
   --radius RADIUS, -r RADIUS                        The RADIUS of the target area. Used to calculate MPBR (Maximum Point Blank Range).
   --velocity VELOCITY, -v VELOCITY                  The projectile VELOCITY (speed). Used to calculate projectile energy, momentum, etc.
   --help, -h                                        Print this help info
   --version, -V                                     Print the ballistic version

VALUE SUFFIXES:
  All input values may be suffixed to allow for broader input selection.

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
```


Rational
--------

There are many online resources to help with calculating ballistics for guns and more archaic weapons such as bows, slings, etc. But there are few command line tools. I like command line tools. They don't require a working Internet connection. They can also be piped into a toolchain to help with automating processes. That's often difficult or next to impossible with most online tools.


### `Justfile`

This is my choice over using `make` to facilitates building the binaries and creating distrobutions. I use [`just`][] which is such a great command runner in the style of `make`! But without the many issues associated with using one of the many, many different versions of `make`. I can barely express the love I have for such tools. I highly recommend it!


Internationalization/Locale
---------------------------

Ballistic can format the human output numbers per locale norms. It currently checks for locale settings via the environment variables LC_CTYPE and LANG. It also allows you to specify a locale such as `en_CA` or `FR-CA` for English or French Canada for instance or simply `EN` for general English speakers, `DE` the German language or country, etc. via the `--locale` option.

### Locales Supported

- `AU` Australia/Australian -- Country and language standard
- `CN` China -- Country standard when using arabic numerals
- `DE` Germany/German -- Country and language standard
- `EN-SIU` English International System of Units standard
- `EN_CA` English Canadian standard
- `EN` English -- The same for the USA, UK, and many others. Who would have thunk?
- `FR-SIU` French International System of Units standard
- `FR_CA` French Canadian standard
- `HK` Hong Kong -- Country standard when using arabic numerals
- `IE` Ireland/Irish -- Country and language standard
- `IN` India/Indian -- Country and language standard
- `IS` Israel/Israeli -- Country and language standard when using arabic numerals
- `JP` Japan/Japanese -- Country and language standard when using arabic numerals

Note that locales with two sets of letters can be seperated by a hyphen or underscore. Both are valid and are interspersed above just for illustrative purposes.


ToDo
----

- Calculate:
	- Projectile drop at distance
	- Projectile momentum at distance
	- Projectile mass given distance and velocity
	- Projectile penetration reference (WIP momentum + ballistic coefficient)
	- Compound Bow (modern cambered bow): projectile energy and velocity
- Automagically:
	- Set output units based on locale. At least switching intelligently between metric and imperial units.


 * * *

Here be Dragons! (my research)
------------------------------

### Articles & References

- [Decimal and Thousands Separators (International Language Environments Guide)][]
- [Decimal Separator - Wikipedia][]
- [ISO 31-0: Numbers - Wikipedia][]
- [SI (International System of Units) - Wikipedia][]


### Symbols

| Category      | Symbol     | Unicode  | Decimal     | Escape     | Numbers  | Description                                                                                 |
| :------:      | :----:     | :-----:  | :-----:     | :----:     | :-----:  | -----------                                                                                 |
| Comma         | &#x0027;   | U+0027   | `&#39;`     | `\u0027`   |          | [Unicode APOSTROPHE][]                                                                      |
| Comma         | &#x002C;   | U+002C   | `&#44;`     | `\u002C`   | &check;  | [Unicode COMMA][]                                                                           |
| Comma         | &#x01144D; | U+01144D | `&#70733;`  | `\u01144D` |          | [Unicode NEWA COMMA][]                                                                      |
| Comma         | &#x01DA87; | U+01DA87 | `&#121479;` | `\u01DA87` |          | [Unicode SIGNWRITING COMMA][]                                                               |
| Comma         | &#x02BD;   | U+02BD   | `&#701;`    | `\u02BD`   |          | [Unicode MODIFIER LETTER REVERSED COMMA][]                                                  |
| Comma         | &#x0312;   | U+0312   | `&#786;`    | `\u0312`   |          | [Unicode COMBINING TURNED COMMA ABOVE][]                                                    |
| Comma         | &#x0314;   | U+0314   | `&#788;`    | `\u0314`   |          | [Unicode COMBINING REVERSED COMMA ABOVE][]                                                  |
| Comma         | &#x0315;   | U+0315   | `&#789;`    | `\u0315`   |          | [Unicode COMBINING COMMA ABOVE RIGHT][]                                                     |
| Comma         | &#x060C;   | U+060C   | `&#1548;`   | `\u060C`   |          | [Unicode ARABIC COMMA][]                                                                    |
| Comma         | &#x275B;   | U+275B   | `&#10075;`  | `\u275B`   |          | [Unicode HEAVY SINGLE TURNED COMMA QUOTATION MARK ORNAMENT][]                               |
| Comma         | &#x275C;   | U+275C   | `&#10076;`  | `\u275C`   |          | [Unicode HEAVY SINGLE COMMA QUOTATION MARK ORNAMENT][]                                      |
| Comma         | &#x2E32;   | U+2E32   | `&#11826;`  | `\u2E32`   |          | [Unicode TURNED COMMA][]                                                                    |
| Comma         | &#x2E34;   | U+2E34   | `&#11828;`  | `\u2E34`   |          | [Unicode RAISED COMMA][]                                                                    |
| Comma         | &#x2E41;   | U+2E41   | `&#11841;`  | `\u2E41`   |          | [Unicode REVERSED COMMA][]. Used in Sindhi.                                                 |
| Comma         | &#x3001;   | U+3001   | `&#12289;`  | `\u3001`   |          | [Unicode IDEOGRAPHIC COMMA][]                                                               |
| Comma         | &#xFE50;   | U+FE50   | `&#65104;`  | `\uFE50`   |          | [Unicode SMALL COMMA][]                                                                     |
| Comma         | &#xFE51;   | U+FE51   | `&#65105;`  | `\uFE51`   |          | [Unicode SMALL IDEOGRAPHIC COMMA][]                                                         |
| Comma         | &#xFF0C;   | U+FF0C   | `&#65292;`  | `\uFF0C`   |          | [Unicode FULLWIDTH COMMA][]                                                                 |
| Comma         | &#xFF64;   | U+FF64   | `&#65380;`  | `\uFF64`   |          | [Unicode HALFWIDTH IDEOGRAPHIC COMMA][]                                                     |
| Decimal Mark  | &#x066B;   | U+066B   | `&#1643;`   | `\u066B`   | &check;  | [Unicode ARABIC DECIMAL SEPARATOR][]                                                        |
| Decimal Mark  | &#x2396;   | U+2396   | `&#9110;`   | `\u2396`   | Keyboard | [Unicode DECIMAL SEPARATOR KEY SYMBOL][]. Noted in [ISO/IEC 9995][ISO/IEC 9995 - Wikipedia] |
| Decimal Mark? | &#x2E30;   | U+2E30   | `&#11824;`  | `\u2E30`   |          | [Unicode RING POINT][]                                                                      |
| Group Mark    | &#x066C;   | U+066C   | `&#1644;`   | `\u066C`   | &check;  | [Unicode ARABIC THOUSANDS SEPARATOR][]                                                      |
| Interpunct    | &#x00B7;   | U+00B7   | `&#183;`    | `\u00B7`   | &check;  | [Unicode MIDDLE DOT][]                                                                      |
| Interpunct    | &#x2022;   | U+2022   | `&#8226;`   | `\u2022`   |          | [Unicode BULLET][]                                                                          |
| Interpunct    | &#x2023;   | U+2023   | `&#8227;`   | `\u2023`   |          | [Unicode TRIANGULAR BULLET][]                                                               |
| Interpunct    | &#x2219;   | U+2219   | `&#8729;`   | `\u2219`   |          | [Unicode BULLET OPERATOR][]                                                                 |
| Interpunct    | &#x22C5;   | U+22C5   | `&#8901;`   | `\u22C5`   |          | [Unicode DOT OPERATOR][]                                                                    |
| Interpunct    | &#x25E6;   | U+25E6   | `&#9702;`   | `\u25E6`   |          | [Unicode WHITE BULLET][]                                                                    |
| Interpunct    | &#x2981;   | U+2981   | `&#10625;`  | `\u2981`   | Math     | [Unicode Z NOTATION SPOT][]                                                                 |
| Interpunct    | &#x2E31;   | U+2E31   | `&#11825;`  | `\u2E31`   |          | [Unicode WORD SEPARATOR MIDDLE DOT][]                                                       |
| Misc.         | &#x1F340;  | U+1F340  | `&#127808;` | `\u1F340`  |          | [Unicode FOUR LEAF CLOVER][]                                                                |

<!--
|               | &#x____;   | U+____   | `&#____;`   | `\u____`   |          | []                                                                                          |
|               | &#x____;   | U+____   | `&#____;`   | `\u____`   |          | []                                                                                          |
|               | &#x____;   | U+____   | `&#____;`   | `\u____`   |          | []                                                                                          |
|               | &#x____;   | U+____   | `&#____;`   | `\u____`   |          | []                                                                                          |
|               | &#x____;   | U+____   | `&#____;`   | `\u____`   |          | []                                                                                          |
-->



### Formulas

Note: These formulas are here for my tracking purposes. They are not necessarily used in the app.

```
PE = mass * gravity * height

Energy = pressure * volume / efficiency
```

PE = Potential Energy


- [PSI to Energy][]


<!--
[Unicode ____]: https://www.fileformat.info/info/unicode/char/____/
[Unicode ____]: https://www.fileformat.info/info/unicode/char/____/
[Unicode ____]: https://www.fileformat.info/info/unicode/char/____/
-->

[`just`]: https://github.com/casey/just
[Decimal Separator - Wikipedia]: https://en.wikipedia.org/wiki/Decimal_separator
[ISO/IEC 9995 - Wikipedia]: https://en.wikipedia.org/wiki/ISO/IEC_9995
[Decimal and Thousands Separators (International Language Environments Guide)]: https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/
[SI (International System of Units) - Wikipedia]: https://en.wikipedia.org/wiki/International_System_of_Units
[ISO 31-0: Numbers - Wikipedia]: https://en.wikipedia.org/wiki/ISO_31-0#Numbers
[PSI to Energy]: https://www.physicsforums.com/threads/psi-to-kw-conversion.700882/
[Unicode APOSTROPHE]: https://www.fileformat.info/info/unicode/char/0027/
[Unicode ARABIC COMMA]: https://www.fileformat.info/info/unicode/char/060C/
[Unicode ARABIC DECIMAL SEPARATOR]: https://www.fileformat.info/info/unicode/char/066B/
[Unicode ARABIC THOUSANDS SEPARATOR]: https://www.fileformat.info/info/unicode/char/066C/
[Unicode BULLET OPERATOR]: https://www.fileformat.info/info/unicode/char/2219/
[Unicode BULLET]: https://www.fileformat.info/info/unicode/char/2022/
[Unicode COMBINING COMMA ABOVE RIGHT]: https://www.fileformat.info/info/unicode/char/0315/
[Unicode COMBINING REVERSED COMMA ABOVE]: https://www.fileformat.info/info/unicode/char/0314/
[Unicode COMBINING TURNED COMMA ABOVE]: https://www.fileformat.info/info/unicode/char/0312/
[Unicode COMMA]: https://www.fileformat.info/info/unicode/char/002C/
[Unicode DECIMAL SEPARATOR KEY SYMBOL]: https://www.fileformat.info/info/unicode/char/2396/
[Unicode DOT OPERATOR]: https://www.fileformat.info/info/unicode/char/22C5/
[Unicode FOUR LEAF CLOVER]: https://www.fileformat.info/info/unicode/char/1F340/
[Unicode FULLWIDTH COMMA]: https://www.fileformat.info/info/unicode/char/FF0C/
[Unicode HALFWIDTH IDEOGRAPHIC COMMA]: https://www.fileformat.info/info/unicode/char/FF64/
[Unicode HEAVY SINGLE COMMA QUOTATION MARK ORNAMENT]: https://www.fileformat.info/info/unicode/char/275C/
[Unicode HEAVY SINGLE TURNED COMMA QUOTATION MARK ORNAMENT]: https://www.fileformat.info/info/unicode/char/275B/
[Unicode IDEOGRAPHIC COMMA]: https://www.fileformat.info/info/unicode/char/3001/
[Unicode MIDDLE DOT]: https://www.fileformat.info/info/unicode/char/00B7/
[Unicode MODIFIER LETTER REVERSED COMMA]: https://www.fileformat.info/info/unicode/char/02BD/
[Unicode NEWA COMMA]: https://www.fileformat.info/info/unicode/char/1144D/
[Unicode RAISED COMMA]: https://www.fileformat.info/info/unicode/char/2E34/
[Unicode REVERSED COMMA]: https://www.fileformat.info/info/unicode/char/2E41/
[Unicode RING POINT]: https://www.fileformat.info/info/unicode/char/2E30/
[Unicode SIGNWRITING COMMA]: https://www.fileformat.info/info/unicode/char/1DA87/
[Unicode SMALL COMMA]: https://www.fileformat.info/info/unicode/char/FE50/
[Unicode SMALL IDEOGRAPHIC COMMA]: https://www.fileformat.info/info/unicode/char/FE51/
[Unicode TRIANGULAR BULLET]: https://www.fileformat.info/info/unicode/char/2023/
[Unicode TURNED COMMA]: https://www.fileformat.info/info/unicode/char/2E32/
[Unicode WHITE BULLET]: https://www.fileformat.info/info/unicode/char/25E6/
[Unicode WORD SEPARATOR MIDDLE DOT]: https://www.fileformat.info/info/unicode/char/2E31/
[Unicode Z NOTATION SPOT]: https://www.fileformat.info/info/unicode/char/2981/

