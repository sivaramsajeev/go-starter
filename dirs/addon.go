package dirs

import (
	"github.com/fatih/color"
	"fmt"
)

func AddonPrint() string {
	color.Cyan("Prints text in cyan.")
	color.Blue("Prints %s in blue.", "text")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

	// Use handy standard colors
	color.Set(color.FgYellow)

	fmt.Println("Existing text will now be in yellow")
	fmt.Printf("This one %s\n", "too")

	color.Unset() // Don't forget to unset

	// You can mix up parameters
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset() // Use it in your function

	fmt.Println("All text will now be bold magenta.")
	return "Addon printed"
}
