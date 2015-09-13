# *naïf*
A small utility for nvm users to keep their SublimeText build systems synched with their installed versions.

## How it works:
naif uses your `NVM_DIR` shell variable to scan and record the installed versions of both forks (Node/io.js). Next, it outputs super-simple .sublime-build files into your `Sublime Text (2|3)/Packages/User` directory, each with the `path` option set to the right `bin` folder inside .nvm. (If you have both ST2 and ST3 installed, it will target ST3.) Later, if you remove versions of Node/io.js that you had installed through nvm, naif will clean up the corresponding build files next time it runs.

## How to get it:
OSX, 64bit only.
If you want to build from source using the go tool: `go get github.com/camhux/naif`
If you don't: download the binary from the releases tab and run from wherever you like: `$ ./naif`. For convenience, you can shove it somewhere on your path and use it like a command: `$ naif`.

## How to use it:
Right now, just run it. You have no options. All of the files it writes are suffixed with `-naif` before their extension, so you can easily see what's been done, and it writes only to your Sublime Text/Packages/User directory.

## Why it's in Go, and not just a shell script:
For fun, and for practice in a language I want to know better.

## If it blows up your computer or doesn't work:
I'm sorry. This was written partially as a learning exercise and, while very simple, is not extensively tested. Open an issue and hear my pleas for forgiveness.
