package justify

import (
	"os"
	"testing"

	"golang.org/x/term"
)

const loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
const loremipsumBlock = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestInvalidInput_ZeroWidth(t *testing.T) {
	got := JustifyLine(loremipsum, 0)
	if got != "" {
		t.Errorf("expected empty string with line width of zero, got: %q", got)
	}
}

func TestInvalidInput_EmptyText(t *testing.T) {
	got := JustifyLine("", 160)
	if got != "" {
		t.Errorf("expected empty string with empty input, got: %q", got)
	}
}

func TestInvalidInput_BlockZeroWidth(t *testing.T) {
	got := Justify(loremipsum, 0)
	if got != "" {
		t.Errorf("expected empty block with line width of zero, got: %q", got)
	}
}

func TestInvalidInput_BlockEmptyText(t *testing.T) {
	got := Justify("", 160)
	if got != "" {
		t.Errorf("expected empty block with empty input, got: %q", got)
	}
}

func TestBasicJustifyLine(t *testing.T) {
	testCases := []struct {
		testName  string
		formatted string
		width     int
	}{
		{"Line width 40",
			`Lorem  ipsum dolor sit amet, consectetur
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
`, 40},
		{"Line width 7",
			`Lorem
ipsum
dolor
sit
amet,
consectetur
adipiscing
elit,
sed  do
eiusmod
tempor
incididunt
ut
labore
et
dolore
magna
aliqua.
Ut enim
ad
minim
veniam,
quis
nostrud
exercitation
ullamco
laboris
nisi ut
aliquip
ex   ea
commodo
consequat.
Duis
aute
irure
dolor
in
reprehenderit
in
voluptate
velit
esse
cillum
dolore
eu
fugiat
nulla
pariatur.
Excepteur
sint
occaecat
cupidatat
non
proident,
sunt in
culpa
qui
officia
deserunt
mollit
anim id
est
laborum.
`, 7},
		{"Line width 80",
			`Lorem  ipsum  dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt  ut  labore  et  dolore  magna  aliqua. Ut enim ad minim veniam, quis
nostrud  exercitation  ullamco  laboris nisi ut aliquip ex ea commodo consequat.
Duis  aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
fugiat  nulla  pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.
`, 80},
		{"single line with huge line size", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n", 1000000007},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			expected := tc.formatted
			got := JustifyLine(loremipsum, tc.width)

			if got != expected {
				t.Errorf("mismatch for line=%d, got:\n------\n%s\n------\nexpected:\n------\n%s\n------\n", tc.width, got, expected)
			}
		})
	}
}

func TestBasicJustifyBlockWithEOL(t *testing.T) {
	const expected = `Lorem  ipsum  dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua.
Ut  enim  ad  minim  veniam,  quis  nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat.
Duis  aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
fugiat nulla pariatur.
Excepteur  sint  occaecat  cupidatat  non  proident,  sunt  in culpa qui officia
deserunt mollit anim id est laborum.
`

	//	t.Logf("Expected block of text with EOL:\n------\n%s\n------\n", expected)
	got := Justify(loremipsumBlock, 80)
	//	t.Logf("Justified block:\n------\n%s\n------\n", got)
	if got != expected {
		t.Errorf("mismatch for block of text with EOL, got:\n------\n%s\n------\nexpected:\n------\n%s\n------\n",
			got, expected)
	}
}

func TestBasicJustifyBlockWithNewlinesNoEOL(t *testing.T) {
	const expected = `Lorem  ipsum  dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua.
Ut  enim  ad  minim  veniam,  quis  nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat.
Duis  aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
fugiat nulla pariatur.
Excepteur  sint  occaecat  cupidatat  non  proident,  sunt  in culpa qui officia
deserunt mollit anim id est laborum.`
	EOL = false

	//	t.Logf("Expected block of text without EOL:\n------\n%s\n------\n", expected)
	got := Justify(loremipsumBlock, 80)
	//	t.Logf("Justified block:\n------\n%s\n------\n", got)
	if got != expected {
		t.Errorf("mismatch for block of text without EOL, got:\n------\n%s\n------\nexpected:\n------\n%s\n------\n",
			got, expected)
	}
}

// Simple terminal test (note: should not work under testing battery).
func TestTerminal(t *testing.T) {
	termWidth := 80 // default terminal width
	var err error

	if term.IsTerminal(int(os.Stdin.Fd())) {
		if termWidth, _, err = term.GetSize(int(os.Stdin.Fd())); err != nil {
			t.Error("could not figure out terminal width; error was", err) // handle (or ignore) error
		}
		t.Log("Terminal width set to", termWidth)
	} else {
		t.Log("Not connected to a TTY; using default width of", termWidth)
	}
}
