package tag

import . "github.com/hndada/mme"

// <a>
func Anchor() Element {
	return Element{
		Tag: "a",
		Text: Text{
			Color:     Blue,
			Underline: true,
		},
		Cursor: CursorPointer,
	}
}

// <strong>
func Strong() Element {
	return Element{
		Tag: "strong",
		Text: Text{
			Font: Font{
				Weight: FontWeightBold,
			},
		},
	}
}

// <em>
func Emphasis() Element {
	return Element{
		Tag: "em",
		Text: Text{
			Font: Font{
				Style: FontStyleItalic,
			},
		},
	}
}

// <hr>
func HorizontalRule() Element {
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

// <address>
func Address() Element {
	return Element{
		Tag: "address",
		Text: Text{
			Font: Font{
				Style: FontStyleItalic,
			},
		},
	}
}

// <pre>
func Pre() Element {
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

// <code>
func Code() Element {
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
