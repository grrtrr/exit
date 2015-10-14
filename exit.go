/*
 * Utilities to exit a program
 */
package exit

import (
	"strings"
	"path"
	"fmt"
	"os"
)

// Print a message @a to stderr and exit with 1 - fmt.Println variant.
func Die(a ...interface{}) {
	if len(a) > 0 {
		fmt.Fprintln(os.Stderr, a...)
	} else {
		fmt.Fprintln(os.Stderr, "Program terminated.")
	}
	os.Exit(1)
}

// Like Die(), but with a format string - fmt.Printf variant.
func Dief(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, strings.TrimSpace(format) + "\n", a...)
	os.Exit(1)
}

// Like Die(), but preprend the name of the program.
func PDie(a ...interface{}) {
	fmt.Fprint(os.Stderr, path.Base(os.Args[0]) + ": ")
	if len(a) > 0 {
		Die(a...)
	} else {
		Die("program terminated.")
	}
}

// Like Dief(), but preprend the name of the program.
func PDief(format string, a ...interface{}) {
	Dief(path.Base(os.Args[0]) + ": " + format, a...)
}


// Die with a FATAL message - fmt.Println variant.
func Fatal(a ...interface{}) {
	fmt.Fprint(os.Stderr, "FATAL: ")
	Die(a...)
}

// Die with a FATAL message - fmt.Printf / log.Fatalf variant.
func Fatalf(format string, a ...interface{}) {
	Dief("FATAL: " + format, a...)
}


// Die with an ERROR message - fmt.Println variant.
func Error(a ...interface{}) {
	fmt.Fprint(os.Stderr, "ERROR: ")
	Die(a...)
}

// Die with an ERROR message - fmt.Printf variant.
func Errorf(format string, a ...interface{}) {
	Dief("ERROR: " + format, a...)
}
