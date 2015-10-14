PACKAGE DOCUMENTATION
## Utilities to exit a program

A small set of utility function to terminate a program with an (error) message,
based on the concept of Perl's [die](http://perldoc.perl.org/functions/die.html).

### Functions

```go
func Die(a ...interface{})
    // Print a message @a to stderr and exit with 1 - fmt.Println variant.

func Dief(format string, a ...interface{})
    // Like Die(), but with a format string - fmt.Printf variant.

func Error(a ...interface{})
    // Die with an ERROR message - fmt.Println variant.

func Errorf(format string, a ...interface{})
    // Die with an ERROR message - fmt.Printf variant.

func Fatal(a ...interface{})
    // Die with a FATAL message - fmt.Println variant.

func Fatalf(format string, a ...interface{})
    // Die with a FATAL message - fmt.Printf / log.Fatalf variant.

func PDie(a ...interface{})
    // Like Die(), but preprend the name of the program.

func PDief(format string, a ...interface{})
    // Like Dief(), but preprend the name of the program.
```
