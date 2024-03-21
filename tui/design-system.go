package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func StartDesignSystem() {

	fmt.Println(renderColorPalette(darkPalette))
	fmt.Println()
	fmt.Println(renderColorPalette(lightPalette))
}

func renderColorPalette(p ColorPalette) string {
	t := "AaBbMmYyZz"
	b := lipgloss.NewStyle().Width(4).Height(1)
	tc := lipgloss.NewStyle().Padding(0, 1).Margin(0, 2)

	boxes := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			b.Copy().Background(lipgloss.Color(p.BrightGrey)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightRed)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightGreen)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightYellow)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightBlue)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightMagenta)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightCyan)).Render(""),
			b.Copy().Background(lipgloss.Color(p.BrightWhite)).Render(""),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			b.Copy().Background(lipgloss.Color(p.DimGrey)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimRed)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimGreen)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimYellow)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimBlue)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimMagenta)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimCyan)).Render(""),
			b.Copy().Background(lipgloss.Color(p.DimWhite)).Render(""),
		),
	)

	texts := lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Left,
			tc.Copy().Foreground(lipgloss.Color(p.BrightGrey)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightRed)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightGreen)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightYellow)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightBlue)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightMagenta)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightCyan)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.BrightWhite)).Render(t),
		),
		lipgloss.JoinVertical(
			lipgloss.Left,
			tc.Copy().Foreground(lipgloss.Color(p.DimGrey)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimRed)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimGreen)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimYellow)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimBlue)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimMagenta)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimCyan)).Render(t),
			tc.Copy().Foreground(lipgloss.Color(p.DimWhite)).Render(t),
		),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		boxes,
		lipgloss.NewStyle().Margin(1, 0).Render(""),
		texts,
	)
}
