package tui

type ColorPalette struct {
	BrightGrey    string
	DimGrey       string
	BrightRed     string
	DimRed        string
	BrightGreen   string
	DimGreen      string
	BrightYellow  string
	DimYellow     string
	BrightBlue    string
	DimBlue       string
	BrightMagenta string
	DimMagenta    string
	BrightCyan    string
	DimCyan       string
	BrightWhite   string
	DimWhite      string
}

var darkPalette = ColorPalette{
	BrightGrey:    "#929292",
	BrightRed:     "#E27373",
	BrightGreen:   "#94B979",
	BrightYellow:  "#FFBA7B",
	BrightBlue:    "#97BEDC",
	BrightMagenta: "#E1C0FA",
	BrightCyan:    "#00988E",
	BrightWhite:   "#DEDEDE",

	DimGrey:    "#BDBDBD",
	DimRed:     "#FFA1A1",
	DimGreen:   "#BDDEAB",
	DimYellow:  "#FFDCA0",
	DimBlue:    "#B1D8F6",
	DimMagenta: "#FBDAFF",
	DimCyan:    "#1AB2A8",
	DimWhite:   "#FFFFFF",
}

var lightPalette = ColorPalette{
	BrightGrey:    "#24292f",
	BrightRed:     "#cf222e",
	BrightGreen:   "#1a7f37",
	BrightYellow:  "#9a6700",
	BrightBlue:    "#0969da",
	BrightMagenta: "#8250df",
	BrightCyan:    "#1b7c83",
	BrightWhite:   "#6e7781",
	DimGrey:       "#57606a",
	DimRed:        "#a40e26",
	DimGreen:      "#2da44e",
	DimYellow:     "#bf8700",
	DimBlue:       "#218bff",
	DimMagenta:    "#a475f9",
	DimCyan:       "#3192aa",
	DimWhite:      "#8c959f",
}
