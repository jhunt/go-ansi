package main

import (
	fmt "github.com/jhunt/go-ansi"
	"os"
)

func main() {
	fmt.Printf("  @k{@k is Black █████}   @K{@K is Black (bold) █████}\n")
	fmt.Printf("  @r{@r is Red ███████}   @R{@R is Red (bold) ███████}\n")
	fmt.Printf("  @g{@g is Green █████}   @G{@G is Green (bold) █████}\n")
	fmt.Printf("  @y{@y is Yellow ████}   @Y{@Y is Yellow (bold) ████}\n")
	fmt.Printf("  @b{@b is Blue ██████}   @B{@B is Blue (bold) ██████}\n")
	fmt.Printf("  @m{@m is Magenta ███}   @M{@M is Magenta (bold) ███}\n")
	fmt.Printf("  @c{@c is Cyan ██████}   @C{@C is Cyan (bold) ██████}\n")
	fmt.Printf("  @w{@w is White █████}   @W{@W is White (bold) █████}\n\n")
	fmt.Printf("  @*{@* is RAINBOW MODE █████████████████████████}\n\n")

	fmt.Fprintf(os.Stderr, "  @k{@k is Black █████}   @K{@K is Black (bold) █████}\n")
	fmt.Fprintf(os.Stderr, "  @r{@r is Red ███████}   @R{@R is Red (bold) ███████}\n")
	fmt.Fprintf(os.Stderr, "  @g{@g is Green █████}   @G{@G is Green (bold) █████}\n")
	fmt.Fprintf(os.Stderr, "  @y{@y is Yellow ████}   @Y{@Y is Yellow (bold) ████}\n")
	fmt.Fprintf(os.Stderr, "  @b{@b is Blue ██████}   @B{@B is Blue (bold) ██████}\n")
	fmt.Fprintf(os.Stderr, "  @m{@m is Magenta ███}   @M{@M is Magenta (bold) ███}\n")
	fmt.Fprintf(os.Stderr, "  @c{@c is Cyan ██████}   @C{@C is Cyan (bold) ██████}\n")
	fmt.Fprintf(os.Stderr, "  @w{@w is White █████}   @W{@W is White (bold) █████}\n\n")
	fmt.Fprintf(os.Stderr, "  @*{@* is RAINBOW MODE █████████████████████████}\n\n")
}
