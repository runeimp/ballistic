

@_list-recipes:
	term-wipe
	just --list


# Build the app for the current OS/Architecture
@build-app:
	#!/usr/bin/env sh
	term-wipe
	if [ '{{os()}}' = 'windows' ] && [ '{{arch()}}' != 'x86_64' ]; then
		just build-win32
	else
		just build-{{os()}}
	fi
	

@build arg='app':
	term-wipe
	# go build ballistic.go
	echo "build-{{arg}}"


# Build all OS/Architecture binarys
@build-all:
	just build-linux
	just build-linux-32bit
	just build-linux-arm7
	just build-macos
	just build-osx
	just build-pi
	just build-win32
	just build-windows
	ls -AhlG bin/*

# Build the Linux (32-bit) binary
build-linux-32bit:
	@echo "Building Linux (386) binary..."
	GOOS=linux GOARCH=386 go build -o bin/linux-386/ballistic ballistic.go

# Build the Linux (64-bit) binary
build-linux:
	@echo "Building Linux (64-bit) binary..."
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/ballistic ballistic.go

# Build the Linux (ARM7) binary
build-linux-arm7:
	@echo "Building Linux (ARM7) binary..."
	GOOS=linux GOARCH=arm GOARM=7 go build -o bin/linux-arm7/ballistic ballistic.go

# Build the macOS (64-bit) binary
build-macos:
	@echo "Building macOS (64-bit) binary..."
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/ballistic ballistic.go

# Build most OS/Architecture binarys
build-most:
	@just build-linux
	@just build-macos
	@just build-win32
	@just build-windows
	just _list-dir 'bin/*'

# Build the OS X (32-bit) binary
build-osx:
	@echo "Building OS X (32-bit) binary..."
	GOOS=darwin GOARCH=386 go build -o bin/osx/ballistic ballistic.go

# Build the Raspberry Pi binary
build-pi:
	@echo "Building Raspberry Pi binary..."
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/raspberry-pi/ballistic ballistic.go

# Build the Windows (32-bit) binary
build-win32:
	@echo "Building Windows (Win32) binary..."
	GOOS=windows GOARCH=386 go build -o bin/win32/ballistic.exe ballistic.go

# Build the Windows (64-bit) binary
build-windows:
	@echo "Building Windows (amd64) binary..."
	GOOS=windows GOARCH=amd64 go build -o bin/win64/ballistic.exe ballistic.go


# Clean, Build, Run
@cbr:
	term-wipe
	# just clean
	just build
	just app-run


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
run:
	@term-wipe
	./ballistic


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

