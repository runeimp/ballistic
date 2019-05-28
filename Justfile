
PROJECT_NAME := 'Ballistic'
DISTRO_NAME := 'Ballistic'
SOURCE_NAME := 'main.go'
BINARY_NAME := 'ballistic'
VERSION_FILE := 'main.go'

alias ver := version

@_default:
	just _term-lw "{{PROJECT_NAME}}"
	just --list


@build arg='app':
	just _term-lw "{{PROJECT_NAME}}"
	rm -rf ./bin
	# echo "build-{{arg}}"
	just _build-{{arg}}

_build binpath='macos' goos='darwin' goarch='amd64' ext='':
	GOOS={{goos}} GOARCH={{goarch}} go build -o bin/{{binpath}}/{{BINARY_NAME}}{{ext}} {{SOURCE_NAME}}


# Build all OS/Architecture binaries
@build-all:
	just _term-wipe
	just _build-linux-amd64
	just _build-linux-386
	just _build-linux-arm7
	just _build-macos-amd64
	just _build-osx-386
	just _build-pi
	just _build-windows-386
	just _build-windows-amd64
	just _list-dir 'bin/*'


# Build the app for the current OS/Architecture
@build-app:
	just _term-wipe
	just _build-app
	
_build-app:
	#!/usr/bin/env sh
	if [ '{{os()}}' = 'windows' ] && [ '{{arch()}}' != 'x86_64' ]; then
		just build-windows-386
	else
		just build-{{os()}}
		if [ -d "${HOME}/dev/bin" ]; then
			cp bin/{{os()}}/{{BINARY_NAME}} "${HOME}/dev/bin/"
		fi
	fi


_build-arm binpath='raspberry-pi' goarm='5' goos='linux' ext='':
	GOOS={{goos}} GOARCH=arm GOARM={{goarm}} go build -o bin/{{binpath}}/{{BINARY_NAME}}{{ext}} {{SOURCE_NAME}}


# Build the Linux (32-bit) binary
@build-linux-386:
	just _term-we "Building Linux (386) binary..."
	just _build-linux-386

@_build-linux-386:
	just _build linux-386 linux 386


# Build the Linux (64-bit) binary
@build-linux-amd64:
	just _term-we "Building Linux (64-bit) binary..."
	just _build-linux-amd64

@_build-linux-amd64:
	just _build linux-amd64 linux amd64


# Build the Linux (ARM7) binary
@build-linux-arm7:
	just _term-we "Building Linux (ARM7) binary..."
	just _build-linux-arm7

@_build-linux-arm7:
	just _build-arm linux-arm7 7


# Build the macOS/OS X Lion+ (64-bit) binary
@build-macos-amd64:
	just _term-we "Building macOS (64-bit) binary..."
	just _build-macos

@_build-macos-amd64:
	just _build macos-amd64 darwin amd64


# Build most OS/Architecture binaries
@build-most:
	just _term-wipe
	just _build-most

@_build-most:
	just _term-wipe
	just _build-linux-amd64
	just _build-macos-amd64
	just _build-windows-386
	just _build-windows-amd64
	just _list-bin


# Build the OS X (32-bit) binary
@build-osx-386:
	just _term-we "Building OS X (32-bit) binary..."
	just _build-osx-386

@_build-osx-386:
	just _build osx-386 darwin 386


# Build the Raspberry Pi binary
@build-pi:
	just _term-we "Building Raspberry Pi binary..."
	just _build-pi

@_build-pi:
	@just _build-arm raspberry-pi 5


# Build the Windows (32-bit) binary
@build-windows-386:
	just _term-we "Building Windows (Win32) binary..."
	just _build-windows-386

@_build-windows-386:
	just _build windows-386 windows 386 '.exe'

# Build the Windows (64-bit) binary
@build-windows-amd64:
	just _term-we "Building Windows (amd64) binary..."
	just _build-windows-amd64
	
@_build-windows-amd64:
	just _build windows-amd64 windows amd64 '.exe'


# Clean, Build, Distro
@cbd +args='':
	just _term-wipe
	# just clean
	just _build-app
	echo
	just distro


# Clean, Build, Run
@cbr +args='':
	just _term-wipe
	# just clean
	just _build-app
	echo
	just _run {{args}}


# Cleanup build artifacts
@clean:
	just _term-wipe
	echo "Cleaning up..."
	rm -f {{BINARY_NAME}}
	rm -rf bin
	just _list-dir


# Setup distrobutions
distro:
	#!/usr/bin/env sh
	just _term-lw "{{PROJECT_NAME}}"
	rm -rf ./distro
	for binpath in ./bin/*/{{BINARY_NAME}}*; do
		pathname=`dirname "${binpath}"`
		distname="{{DISTRO_NAME}}-v$(just version)"
		distpath="./distro/${distname}/${distname}-${pathname:6}"
		# echo " $binpath: ${binpath}"
		# echo "$pathname: ${pathname}"
		# echo "$distname: ${distname}"
		# echo "$distpath: ${distpath}"
		
		mkdir -p "${distpath}"
		echo
		cp "${binpath}" "${distpath}/"
		cp "README.md" "${distpath}/"
		cp "LICENSE" "${distpath}/"
		just _list-dir "${distpath}"
		just _dirzip "${distpath}"
		echo

		# echo "${pathname}"
	done
	just _list-dir ./distro


_dirzip path:
	#!/usr/bin/env sh
	child={{path}}
	parent=.
	echo "DirZip: {{path}}"
	# echo "  dirzip path: {{path}}"
	# echo " dirzip child: ${child}"
	# echo "dirzip parent: ${parent}"
	cd "${parent}"
	ditto -ck --keepParent --zlibCompressionLevel 9 --norsrc --noqtn --nohfsCompression "${child}" "${child}.zip"


# Justfile Environment Variables
@env:
	just _term-lw "{{PROJECT_NAME}}"
	echo "  Justfile Environment Variables:"
	env | sort


# Just info
@info:
	just _term-lw "{{PROJECT_NAME}}"
	echo "os_family(): {{os_family()}}"
	echo "       os(): {{os()}}"
	echo "     arch(): {{arch()}}"


_list-bin:
	#!/usr/bin/env sh
	if [ '{{os()}}' = 'macos' ]; then
		ls -AlhG bin/*
	else
		ls -Alh --color bin/*
	fi

_list-dir path='.':
	#!/usr/bin/env sh
	if [ '{{os()}}' = 'macos' ]; then
		echo '$ ls -AlhG "{{path}}"'
		ls -AlhG "{{path}}"
	else
		echo '$ ls -Alh --color "{{path}}"'
		ls -Alh --color "{{path}}"
	fi


# List all paths in $PATH
path:
	just _term-wipe
	echo ${PATH} | tr ":" "\n"
	@echo


# Run Helper
@run id='app' +args='':
	just _term-lw "{{PROJECT_NAME}}"
	just run-{{id}} {{args}}

# Run the app
@run-app +args='':
	echo "$ {{BINARY_NAME}} {{args}}"
	go run {{SOURCE_NAME}} {{args}}


# Test Helper
@test id='app' +args='-m 5g -v 2000fps':
	just _term-lw "{{PROJECT_NAME}}"
	just test-{{id}} {{args}}

# Test the app
test-app +args='':
	@# echo "$ {{BINARY_NAME}} {{args}}"
	go run {{SOURCE_NAME}} {{args}}
	LC_CTYPE=DE go run {{SOURCE_NAME}} {{args}}
	LC_CTYPE=IN go run {{SOURCE_NAME}} {{args}}

test-go +args='':
	go test

# Run time tests with timeit
speed:
	just _term-wipe
	timeit ./{{BINARY_NAME}} RuneImp "./{{BINARY_NAME}} RuneImp 'Command Line'"


# Terminal Helper
@term +args='wipe':
	#!/usr/bin/env sh
	if [ '{{args}}' = 'wipe' ]; then
		just term-{{args}}
	else
		just _term-{{args}}
	fi

# Helper recipe to change the terminal label
@_term-label label:
	printf "\033]0;{{label}}\007"

# Helper recipe to change the terminal label, and echo
@_term-le label:
	just _term-label "{{label}}"
	echo "{{label}}"

# Helper recipe to echo, and wipe the buffer
@_term-we label:
	just _term-wipe
	echo "{{label}}"

# Helper recipe to change the terminal label, echo, and wipe the buffer
@_term-lwe label:
	just _term-label "{{label}}"
	just _term-wipe
	echo "{{label}}"

# Helper recipe to change the terminal label and wipe the buffer
@_term-lw label:
	just _term-label "{{label}}"
	just _term-wipe

# Wipe the terminal buffer
@_term-wipe:
	#!/bin/sh
	if [[ ${#VISUAL_STUDIO_CODE} -gt 0 ]]; then
		clear
	elif [[ ${KITTY_WINDOW_ID} -gt 0 ]] || [[ ${#TMUX} -gt 0 ]] || [[ "${TERM_PROGRAM}" = 'vscode' ]]; then
		printf '\033c'
	elif [[ "$(uname)" == 'Darwin' ]] || [[ "${TERM_PROGRAM}" = 'Apple_Terminal' ]] || [[ "${TERM_PROGRAM}" = 'iTerm.app' ]]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [[ -x "$(which tput)" ]]; then
		tput reset
	elif [[ -x "$(which reset)" ]]; then
		reset
	else
		clear
	fi


# Prints the compiler or interpreter version(s)
@version:
	cat {{VERSION_FILE}} | grep -F 'const APP_VERSION' | cut -d'"' -f2
	# go version
