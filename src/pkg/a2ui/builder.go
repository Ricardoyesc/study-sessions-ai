package a2ui

import (
	"sai-server/internal/domain"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func NewText(id, content, variant string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeText,
		Props: map[string]interface{}{
			"content": content,
			"variant": variant,
		},
	}
}

func NewRichText(id, markdown string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeRichText,
		Props: map[string]interface{}{
			"markdown":   markdown,
			"accessible": true,
		},
	}
}

func NewImage(id, url, altText string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeImage,
		Props: map[string]interface{}{
			"url":     url,
			"altText": altText,
		},
	}
}

func NewAudioPlayer(id, url string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeAudioPlayer,
		Props: map[string]interface{}{
			"url":     url,
			"autoPlay": false,
		},
	}
}

func NewColumn(id string, children []string, props map[string]interface{}) domain.A2UIComponent {
	if props == nil {
		props = map[string]interface{}{"gap": 16, "padding": 24}
	}
	return domain.A2UIComponent{
		ID:       id,
		Type:     ComponentTypeColumn,
		Children: children,
		Props:    props,
	}
}

func NewRow(id string, children []string, props map[string]interface{}) domain.A2UIComponent {
	if props == nil {
		props = map[string]interface{}{"alignment": "center"}
	}
	return domain.A2UIComponent{
		ID:       id,
		Type:     ComponentTypeRow,
		Children: children,
		Props:    props,
	}
}

func NewCard(id string, children []string, props map[string]interface{}) domain.A2UIComponent {
	if props == nil {
		props = map[string]interface{}{"elevation": 2}
	}
	return domain.A2UIComponent{
		ID:       id,
		Type:     ComponentTypeCard,
		Children: children,
		Props:    props,
	}
}

func NewQuizCard(id, question string, options []string, mode, endpoint string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeQuizCard,
		Props: map[string]interface{}{
			"question": question,
			"options":  options,
			"mode":     mode,
		},
		Events: map[string]string{
			"onSubmit": endpoint,
		},
	}
}

func NewSocraticDialog(id, prompt, context, endpoint string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeSocraticDialog,
		Props: map[string]interface{}{
			"prompt":  prompt,
			"context": context,
		},
		Events: map[string]string{
			"onSubmit": endpoint,
		},
	}
}

func NewProgressBar(id string, value, max float64) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeProgressBar,
		Props: map[string]interface{}{
			"value": value,
			"max":   max,
		},
	}
}

func NewButton(id, label, variant string) domain.A2UIComponent {
	return domain.A2UIComponent{
		ID:   id,
		Type: ComponentTypeButton,
		Props: map[string]interface{}{
			"label":   label,
			"variant": variant,
		},
	}
}

func DefaultDataModel() domain.A2UIDataModel {
	return domain.A2UIDataModel{
		Theme:         "system",
		FontFamily:    "sans-serif",
		FontScale:     1.0,
		ColorPalette:  "default",
		HighContrast:  false,
		ReducedMotion: false,
		Language:      "es",
	}
}

func (b *Builder) BuildCapsuleSurface(id, topic string, components map[string]domain.A2UIComponent, dataModel domain.A2UIDataModel) *domain.A2UISurface {
	return &domain.A2UISurface{
		SurfaceID:     id,
		RootComponent: "root",
		Components:    components,
		DataModel:     dataModel,
	}
}

func (b *Builder) BuildSocraticSurface(id string, components map[string]domain.A2UIComponent, dataModel domain.A2UIDataModel) *domain.A2UISurface {
	return &domain.A2UISurface{
		SurfaceID:     id,
		RootComponent: "root",
		Components:    components,
		DataModel:     dataModel,
	}
}
