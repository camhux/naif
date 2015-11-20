# *naïf*
A small utility for nvm users to keep their SublimeText build systems synched with their installed versions.

## How it works:
naif uses your `NVM_DIR` environment variable to scan and record the installed versions of both forks (Node/io.js). Next, it outputs a file named "Node (naif).sublime-build" into your `Sublime Text (2|3)/Packages/User` directory. (If you have both ST2 and ST3 installed, it will target ST3.) This build system selects for `.js` files and will, by default, invoke the latest Node version present in your .nvm/versions. All of your installed versions are listed as variants of the build system and can be specifically invoked with `Shift + Cmd + B`.

## How to get it:
OSX, 64bit only.

If you want to build from source using the go tool: `go get github.com/camhux/naif`

If you don't: download the binary from the releases tab and run from wherever you like: `$ ./naif`.

For convenience, you can shove it somewhere on your path and use it as a command: `$ naif`.

## How to use it:
Right now, just run it. You have no options. The output path is logged to console and the output file is marked with `(naif)`, so all effects of the program should be transparent.

## Why it's in Go, and not just a shell script:
For fun, portability, practice, and (maybe if you've installed every version of Node possible) speed.
