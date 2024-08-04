package chart

type Options struct {
	Responsive *bool     `json:"responsive,omitempty"`
	Title      *Title    `json:"title,omitempty"`
	Legend     *Legend   `json:"legend,omitempty"`
	Tooltips   *Tooltips `json:"tooltips,omitempty"`
	Scales     *Scales   `json:"scales,omitempty"`
}

type Title struct {
	Display *bool   `json:"display,omitempty"`
	Text    *string `json:"text,omitempty"`
}

type Legend struct {
	Display  *bool   `json:"display,omitempty"`
	Position *string `json:"position,omitempty"`
}

type Tooltips struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type Axes struct {
	Display    *bool   `json:"display,omitempty"`
	ScaleLabel *string `json:"scaleLabel,omitempty"`
}

type Scales struct {
	XAxes []Axes `json:"xAxes,omitempty"`
	YAxes []Axes `json:"yAxes,omitempty"`
}
