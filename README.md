# justify
![Justify Gopher Logo](assets/justify-gopher-logo-128x128.png)  
A Go package to justify text (left and right alignment), old-school typewriter style.

## Usage

```go
package main

import (
	"fmt"

	"github.com/GwynethLlewelyn/justify"
)

func main() {
	fmt.Println(justify.Justify(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum`, 40))	
}
```

produces the following output:

```
Lorem  ipsum dolor sit amet, consectetur
adipiscing  elit,  sed do eiusmod tempor
incididunt  ut  labore  et  dolore magna
aliqua.  Ut  enim  ad minim veniam, quis
nostrud   exercitation  ullamco  laboris
nisi ut aliquip ex ea commodo consequat.
Duis  aute  irure dolor in reprehenderit
in voluptate velit esse cillum dolore eu
fugiat  nulla  pariatur.  Excepteur sint
occaecat cupidatat non proident, sunt in
culpa  qui  officia deserunt mollit anim
id est laborum.
```

## Functions

`Justify(string, int)` assumes that the string represents a block of text, with multiple lines of any length (separated by newlines, `\n`). Each line corresponds to a paragraph, which gets formatted individually (and may therefore add additional lines as needed). It returns a string including the `\n`; use `strings.Split()` (or something equivalent) to retrieve an array of indvidual lines, if needed.  

`JustifyLine(string, int)` will assume that all the provided text should be formatted (up to a line width), considering *all* whitespace — including newlines, tabs, etc. — to be irrelevant. The text gets formatted as if the input were single line (backslashed formatting commands have no special meaning and are returned literally). Note that `Justify()` actually calls `JustifyLine()` repeatedly, with some additional formatting.

## Caveat

This module assumes the Unix convention of using just a CR (`\n`) to terminate a line, and not the CSLF (`\n\r`) convention used by Windows. There was an attempt made to use an universal solution (including using all Unicode points specifying new lines or new paragraphs), but its implementation was not trivial. WiP (PRs welcome!).

## Variables

To avoid outputting the last newline (whenever calling this function from somewhere that already properly inserts a newline at the end), just set `justify.EOL = false`. Not goroutine-safe. 

## Tip

On a CLI application, if you wish to justify some text to the current width of the console, you can use something like the following code:

```go
package main

import (
	"fmt"
	"os"

	"github.com/GwynethLlewelyn/justify"
	"golang.org/x/term"
)

func main() {
	termWidth := 80 // default terminal width
	var err error

	if term.IsTerminal(int(os.Stdin.Fd())) {
		if termWidth, _, err = term.GetSize(int(os.Stdin.Fd())); err != nil {
			fmt.Println("could not figure out terminal width; error was", err) // handle (or ignore) error
		}
		fmt.Println("Terminal width set to", termWidth)
	} else {
		fmt.Println("Not connected to a TTY; using default width of", termWidth)
	}

	fmt.Println(justify.Justify("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\nUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum", termWidth))
}
```
[![Go Playground](https://img.shields.io/badge/Go_Playground-%2300ADD8?style=plastic&logo=go&logoColor=%2300ADD8&labelColor=%23FFFFFF)](https://go.dev/play/p/hswg1Fl1h_V)


[![Go](https://github.com/GwynethLlewelyn/justify/actions/workflows/go.yml/badge.svg)](https://github.com/GwynethLlewelyn/justify/actions/workflows/go.yml)