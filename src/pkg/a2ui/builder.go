package a2ui

import "sai-server/internal/domain"

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) BuildCapsuleSurface(capsuleID, topic string, components map[string]domain.A2UIComponent, dataModel domain.A2UIDataModel) *domain.A2UISurface {
	return &domain.A2UISurface{
		SurfaceID:     capsuleID,
		RootComponent: "root",
		Components:    components,
		DataModel:     dataModel,
	}
}

func (b *Builder) BuildQuizSurface(sessionID string, components map[string]domain.A2UIComponent, dataModel domain.A2UIDataModel) *domain.A2UISurface {
	return &domain.A2UISurface{
		SurfaceID:     sessionID,
		RootComponent: "root",
		Components:    components,
		DataModel:     dataModel,
	}
}

func (b *Builder) BuildSocraticSurface(sessionID string, components map[string]domain.A2UIComponent, dataModel domain.A2UIDataModel) *domain.A2UISurface {
	return &domain.A2UISurface{
		SurfaceID:     sessionID,
		RootComponent: "root",
		Components:    components,
		DataModel:     dataModel,
	}
}
