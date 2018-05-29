Ballistic v0.5.0
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
$ ballistic -m 123gr -v 50000fps

  Projectile Velocity:    50,000.001600 feet per second
    Projectile Energy: 1,851,154.550587 joules
  Projectile Momentum:       121.466834 meter kilogram per second
Max Point Blank Range:    15,147.150313 feet

$ ballistic -m 123gr -v 50000fps --locale IN

  Projectile Velocity:    50,000.001600 feet per second
    Projectile Energy: 18,51,154.550587 joules
  Projectile Momentum:       121.466834 meter kilogram per second
Max Point Blank Range:    15,147.150313 feet

```

Note that the first run defaults to my locale of `en_US.UTF-8`. If a locale is not found or supported (yet) the default is `en_US`.


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

### Calculate initial velocity and MPBR based on projection angel and distance (on a horizontal plan)

```text
$ ballistic -a 45 -d 100m

  Projectile Velocity: 790.133355 meters per second
Max Point Blank Range: 239.365366 meters

```

#### Calculate everything by adding projectile mass

```text
$ ballistic -a 45 -d 100m -m 300gr

  Projectile Velocity:   790.133355 meters per second
    Projectile Energy: 6,068.197170 joules
  Projectile Momentum:    15.359932 meter kilogram per second
Max Point Blank Range:   239.365366 meters

```

### Help Info

```text
$ ballistic

NAME:
   Ballistic - Calculates what it can based on provided input.

USAGE:
   ballistic [global options]

VERSION:
   0.5.0

GLOBAL OPTIONS:
   --debug, -D                                             Output debug info
   --draw-length LENGTH, --length LENGTH, -l LENGTH        Bow or sling shot draw LENGTH. Used to calculate projectile velocity, energy, etc.
   --draw-weight WEIGHT, --weight WEIGHT, -w WEIGHT        Bow or sling shot draw WEIGHT. Used to calculate projectile velocity, energy, etc.
   --json, -j                                              Output JSON data
   --locale LOCALE, --local LOCALE                         The LOCALE to format number output for. (default: "en_US") [$LC_CTYPE, $LANG]
   --precision PRECISION, --float PRECISION, -f PRECISION  The output floating point PRECISION (numbers after decimal mark). (default: "6")
   --pretty-print, --pretty, -p                            Pretty printed JSON output
   --projectile MASS, --mass MASS, -m MASS                 Projectile MASS (weight). Used to calculate projectile velocity, energy, etc.
   --projectile-range value, --distance value, -d value    The distance the projectile traveled
   --projection-angle value, --angle value, -a value       The projection angle or trajectory of projectile
   --radius RADIUS, -r RADIUS                              The RADIUS of the target area. Used to calculate MPBR (Maximum Point Blank Range). (default: "225mm")
   --velocity VELOCITY, -v VELOCITY                        The projectile VELOCITY (speed). Used to calculate projectile energy, momentum, etc.
   --help, -h                                              Output this help info
   --version, -V                                           Output the ballistic app version

VALUE SUFFIXES:
  All input values may be suffixed to allow for broader input selection.

  ANGEL
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

```


Rational
--------

There are many online resources to help with calculating ballistics for guns and more archaic weapons such as bows, slings, etc. But there are few command line tools. I like command line tools. They don't require a working Internet connection. They can also be piped into a toolchain to help with automating processes. That's often difficult or next to impossible with most online tools.


### `Justfile`

This is my choice over using `make` to facilitates building the binaries and creating distrobutions. I use [`just`][] which is such a great command runner in the style of `make`! But without the many issues associated with using one of the many, many different versions of `make`. I can barely express the love I have for such tools. I highly recommend it!


Internationalization/Locale
---------------------------

Ballistic can format the human output numbers per locale norms. It currently checks for locale settings via the environment variables `LC_CTYPE` and `LANG`. It also allows you to specify a locale such as `en_CA` or `FR-CA` for English or French Canada for instance or simply `EN` for general English speakers, `DE` the German language or country, etc. via the `--locale` option.

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



[`just`]: https://github.com/casey/just
[Decimal Separator - Wikipedia]: https://en.wikipedia.org/wiki/Decimal_separator
[Decimal and Thousands Separators (International Language Environments Guide)]: https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/
[SI (International System of Units) - Wikipedia]: https://en.wikipedia.org/wiki/International_System_of_Units
[ISO 31-0: Numbers - Wikipedia]: https://en.wikipedia.org/wiki/ISO_31-0#Numbers


