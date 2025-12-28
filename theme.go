package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var baseTheme = theme.DarkTheme()

var (
	//go:embed assets/fonts/AtkinsonHyperlegible-Regular.ttf
	fontAtkinsonRegular []byte
	//go:embed assets/fonts/AtkinsonHyperlegible-Bold.ttf
	fontAtkinsonBold []byte
	//go:embed assets/fonts/SpaceMono-Regular.ttf
	fontSpaceMonoRegular []byte
)

var (
	atkinsonRegular  = fyne.NewStaticResource("AtkinsonHyperlegible-Regular.ttf", fontAtkinsonRegular)
	atkinsonBold     = fyne.NewStaticResource("AtkinsonHyperlegible-Bold.ttf", fontAtkinsonBold)
	spaceMonoRegular = fyne.NewStaticResource("SpaceMono-Regular.ttf", fontSpaceMonoRegular)
)

type olivetumDarkTheme struct{}

func (olivetumDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x0B, G: 0x0F, B: 0x14, A: 0xFF}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF} // slate-900-ish
	case theme.ColorNameButton:
		return color.NRGBA{R: 0x1F, G: 0x29, B: 0x37, A: 0xFF} // slate-800
	case theme.ColorNameHover:
		return color.NRGBA{R: 0x2B, G: 0x37, B: 0x49, A: 0xFF}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0x2A, G: 0x33, B: 0x42, A: 0xFF}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xE5, G: 0xE7, B: 0xEB, A: 0xFF}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x9C, G: 0xA3, B: 0xAF, A: 0xFF}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0x6B, G: 0x72, B: 0x80, A: 0xFF}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0x1A, G: 0x22, B: 0x2E, A: 0xFF}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x7C, G: 0xB3, B: 0x42, A: 0xFF} // olive accent
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0x7C, G: 0xB3, B: 0x42, A: 0xFF}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0x22, G: 0x37, B: 0x1D, A: 0xFF}
	default:
		return baseTheme.Color(name, variant)
	}
}

func (olivetumDarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return spaceMonoRegular
	}
	if style.Bold {
		return atkinsonBold
	}
	return atkinsonRegular
}

func (olivetumDarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return baseTheme.Icon(name)
}

func (olivetumDarkTheme) Size(name fyne.ThemeSizeName) float32 {
	return baseTheme.Size(name)
}
