Ballistic v0.2.1
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
$ ballistic -m 300gr -v 800fps

  Projectile Velocity:   800.000026 feet per second
    Projectile Energy:  1155.842841 joules
  Projectile Momentum:     4.740169 meter kilogram per second
Max Point Blank Range:   242.354405 feet

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


ToDo
----

- Calculate:
	- Projectile drop at distance
	- Projectile momentum at distance
	- Projectile mass given distance and velocity
	- Projectile penetration reference (WIP momentum + ballistic coefficient)
	- Compound Bow (modern cambered bow): projectile energy and velocity


Formulas
--------

Note: These formulas are here for my tracking purposes. They are not necessarily used in the app.

```
PE = mass * gravity * height

Energy = pressure * volume / efficiency
```

PE = Potential Energy


- [PSI to Energy][]




[`just`]: https://github.com/casey/just
[PSI to Energy]: https://www.physicsforums.com/threads/psi-to-kw-conversion.700882/


