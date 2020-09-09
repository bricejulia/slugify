package repl

import (
	"bufio"
	"fmt"
	"github.com/mozillazg/go-unidecode"
	"io"
	"regexp"
	"strings"
)

const PROMPT string = ">> "
var re = regexp.MustCompile("[^a-z0-9]+")

// Start main repl loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		io.WriteString(out, strings.Trim(re.ReplaceAllString(strings.ToLower(unidecode.Unidecode(line)), "-"), "-"))
		io.WriteString(out, "\n")
	}
}