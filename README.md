Ballistic
=========

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

```
-w | --weight WEIGHT   Weight specified as noted bellow
-v | --velocity SPEED  Speed specified as noted bellow



The weight value can be an integer or floating point value with an optional suffix from the following list. If no suffix is provided then the value is expected to be in grains.

gr = grains
g  = grams
#  = pounds
st = short tons (2000 pounds)
lt = long ton (2,240 pounds)
mt = metric tonne (1,000 kg)

The velocity value can be an integer or floating point value with an optional suffix from the following list. If no suffix is provided then the value is expected to be in fps (feet per second).

f | fps = Feet per second
m | mps = Meters per second
mph = Miles per hour
```


Rational
--------

There are many online resources to help with calculating ballistics for guns and more archaic weapons such as bows, slings, etc. But there are few command line tools. I like command line tools. They don't require a working Internet connection. They can also be piped into a toolchain to help with automating processes. That's often difficult or next to impossible with most online tools.



