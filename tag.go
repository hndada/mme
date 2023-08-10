package mme

// Anchor tag
func NewATag() Element {
	return Element{
		Tag: "a",
		Text: Text{
			Color:     Blue,
			Underline: true,
		},
		Cursor: CursorPointer,
	}
}

// Strong tag
func NewStrongTag() Element {
	return Element{
		Tag: "strong",
		Text: Text{
			Font: Font{
				Weight: FontWeightBold,
			},
		},
	}
}

// Emphasis tag
func NewEmTag() Element {
	return Element{
		Tag: "em",
		Text: Text{
			Font: Font{
				Style: FontStyleItalic,
			},
		},
	}
}

func NewHorizontalRuleTag() Element {
	return Element{
		Tag: "hr",
		Size: Size{
			Height: "1px",
		},
		Border: Border{
			Margin: VerticalHorizontal("10px", "0"),
		},
		Background: Background{
			FillColor: "#ccc",
		},
	}
}
func NewAddressTag() Element {
	return Element{
		Tag: "address",
		Text: Text{
			Font: Font{
				Style: FontStyleItalic,
			},
		},
	}
}

func NewPreTag() Element {
	return Element{
		Tag: "pre",
		Border: Border{
			Radius:  AllCorners("3px"),
			Padding: AllSides("10px"), // Unit(10, Pixel)
		},
		Background: Background{
			FillColor: Gray94,
		},
		Text: Text{
			Font:       NewFont("monospace"),
			WhiteSpace: WhiteSpacePre,
		},
	}
}

func NewCodeTag() Element {
	return Element{
		Tag: "code",
		Border: Border{
			Radius:  AllCorners("3px"),
			Padding: VerticalHorizontal("2px", "4px"),
		},
		Background: Background{
			FillColor: Gray94,
		},
		Text: Text{
			Font: NewFont("monospace"),
		},
	}
}
