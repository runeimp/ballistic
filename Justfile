

@_list-recipes:
	term-wipe
	just --list


# Build the app for the current OS/Architecture
@build-app:
	term-wipe
	just _build-app
	
_build-app:
	#!/usr/bin/env sh
	if [ '{{os()}}' = 'windows' ] && [ '{{arch()}}' != 'x86_64' ]; then
		just build-win32
	else
		just build-{{os()}}
	fi

@build arg='app':
	term-wipe
	# echo "build-{{arg}}"
	just _build-{{arg}}


_build binpath='macos' goos='darwin' goarch='amd64' ext='':
	GOOS={{goos}} GOARCH={{goarch}} go build -o bin/{{binpath}}/ballistic{{ext}} ballistic.go

_build-arm binpath='raspberry-pi' goarm='5' goos='linux' ext='':
	GOOS={{goos}} GOARCH=arm GOARM={{goarm}} go build -o bin/{{binpath}}/ballistic{{ext}} ballistic.go

# Build all OS/Architecture binarys
@build-all:
	term-wipe
	just _build-linux
	just _build-linux-32bit
	just _build-linux-arm7
	just _build-macos
	just _build-osx
	just _build-pi
	just _build-win32
	just _build-windows
	just _list-dir 'bin/*'

# Build the Linux (32-bit) binary
build-linux-32bit:
	@echo "Building Linux (386) binary..."
	@# GOOS=linux GOARCH=386 go build -o bin/linux-386/ballistic ballistic.go
	@just _build linux-386 linux 386

# Build the Linux (64-bit) binary
build-linux:
	@echo "Building Linux (64-bit) binary..."
	@# GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/ballistic ballistic.go
	@just _build linux-amd64 linux amd64

# Build the Linux (ARM7) binary
build-linux-arm7:
	@echo "Building Linux (ARM7) binary..."
	@# GOOS=linux GOARCH=arm GOARM=7 go build -o bin/linux-arm7/ballistic ballistic.go
	@just _build-arm linux-arm7 7

# Build the macOS (64-bit) binary
build-macos:
	@echo "Building macOS (64-bit) binary..."
	@# GOOS=darwin GOARCH=amd64 go build -o bin/macos/ballistic ballistic.go
	@just _build macos darwin amd64

# Build most OS/Architecture binarys
@build-most:
	just build-linux
	just _build-macos
	just _build-win32
	just _build-windows
	just _list-dir 'bin/*'

# Build the OS X (32-bit) binary
build-osx:
	@echo "Building OS X (32-bit) binary..."
	@# GOOS=darwin GOARCH=386 go build -o bin/osx/ballistic ballistic.go
	@just _build osx darwin 386

# Build the Raspberry Pi binary
build-pi:
	@echo "Building Raspberry Pi binary..."
	@# GOOS=linux GOARCH=arm GOARM=5 go build -o bin/raspberry-pi/ballistic ballistic.go
	@just _build-arm raspberry-pi 5

# Build the Windows (32-bit) binary
build-win32:
	@echo "Building Windows (Win32) binary..."
	GOOS=windows GOARCH=386 go build -o bin/win32/ballistic.exe ballistic.go
	@just _build win32 windows 386

# Build the Windows (64-bit) binary
build-windows:
	@echo "Building Windows (amd64) binary..."
	GOOS=windows GOARCH=amd64 go build -o bin/windows/ballistic.exe ballistic.go
	@just _build windows windows amd64


# Clean, Build, Run
@cbr +args='':
	term-wipe
	# just clean
	just _build-app
	echo
	just _run {{args}}


# Cleanup build artifacts
@clean:
	term-wipe
	echo "Cleaning up..."
	rm -f ballistic
	rm -rf bin
	just _list-dir

# Just info
@info:
	term-wipe
	echo "os_family(): {{os_family()}}"
	echo "os(): {{os()}}"
	echo "arch(): {{arch()}}"

_list-dir path='.':
	#!/usr/bin/env sh
	if [ '{{os()}}' = 'macos' ]; then
		ls -AlhG "{{path}}"
	else
		ls -Alh --color "{{path}}"
	fi

# Run the app
run +args='':
	@term-wipe
	@just _run {{args}}

@_run +args='':
	# BALLISTIC_WEIGHT=grams bin/{{os()}}/ballistic
	# bin/{{os()}}/ballistic
	echo ballistic {{args}}
	go run ballistic.go {{args}}


# Run time tests with timeit
speed:
	@term-wipe
	timeit ./ballistic RuneImp "./ballistic RuneImp 'Command Line'"


# Test all examples
test:
	@term-wipe
	@echo
	./ballistic 'Command Line'
	@echo


# Prints the compiler or interpreter version(s)
ver:
	@term-wipe
	@go version

