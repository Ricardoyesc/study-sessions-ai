package domain

type A2UISurface struct {
	SurfaceID     string               `json:"surfaceId"`
	RootComponent string               `json:"rootComponent"`
	Components    map[string]A2UIComponent `json:"components"`
	DataModel     A2UIDataModel        `json:"dataModel"`
}

type A2UIComponent struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`
	Children []string               `json:"children,omitempty"`
	Props    map[string]interface{} `json:"props"`
	Events   map[string]string      `json:"events,omitempty"`
}

type A2UIDataModel struct {
	Theme         string  `json:"theme"`
	FontFamily    string  `json:"fontFamily"`
	FontScale     float64 `json:"fontScale"`
	ColorPalette  string  `json:"colorPalette"`
	HighContrast  bool    `json:"highContrast"`
	ReducedMotion bool    `json:"reducedMotion"`
	Language      string  `json:"language"`
}
