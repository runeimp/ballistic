Ballistic v0.2.0
================

Command line ballistics calculator


Features
--------

- Calculate:
	- Projectile drop at distance
	- Projectile energy and velocity at distance
	- Projectile weight given distance and velocity
	- Projectile momentum at distance
	- Projectile penetration reference
	- <abbr title="Maximum Point Blank Range">MPBR</abbr>: Maximum Point Blank Range or Battle Zero is a military term refering the maximum distance a weapon can be fired to hit the torso of a human target (roughly 18&times;9 inches) every time (baring extreme weather or cover conditions) when aiming at the center of mass.
	- Compound Bow (modern cambered bow): projectile energy and velocity 
- Output
	- Human formated for interactive usage
	- JSON formated for easy scripting


Usage
-----

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



Formulas
--------

Note: These formulas are here for my tracking purposes. They are not necessarily used in the app.

```
PE = mass * gravity * height

Energy = pressure * volume / efficiency
```

PE = Potential Energy


- [PSI to Energy][]




[PSI to Energy]: https://www.physicsforums.com/threads/psi-to-kw-conversion.700882/
