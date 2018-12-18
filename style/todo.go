package style

var (
	ListContainerStyle = Style(
		Prop{"width", "100%"},
		Prop{"display", "flex"},
		Prop{"flex-wrap", "wrap"},
	)

	ListStyle = Style(
		Prop{"width", "50%"},
		Prop{"max-width", "250px"},
	)
)
