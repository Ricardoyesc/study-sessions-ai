package a2ui_test

import (
	"testing"

	"sai-server/pkg/a2ui"
)

func TestDiffProps_setsComponentIDAndProps(t *testing.T) {
	props := map[string]interface{}{"color": "red", "size": 14}
	got := a2ui.DiffProps("comp1", props)
	if got.ComponentID != "comp1" {
		t.Errorf("ComponentID: want comp1 got %s", got.ComponentID)
	}
	if got.Props["color"] != "red" {
		t.Errorf("color: want red got %v", got.Props["color"])
	}
}

func TestDiffProps_emptyProps_stillValid(t *testing.T) {
	got := a2ui.DiffProps("comp2", map[string]interface{}{})
	if got.ComponentID != "comp2" {
		t.Errorf("ComponentID: want comp2 got %s", got.ComponentID)
	}
	if len(got.Props) != 0 {
		t.Errorf("Props: want empty got %v", got.Props)
	}
}

func TestNewA2UIUpdate_wrapsUpdatesSlice(t *testing.T) {
	u1 := a2ui.DiffProps("c1", map[string]interface{}{"x": 1})
	u2 := a2ui.DiffProps("c2", map[string]interface{}{"y": 2})
	payload := a2ui.NewA2UIUpdate(u1, u2)
	if len(payload.Updates) != 2 {
		t.Fatalf("Updates: want 2 got %d", len(payload.Updates))
	}
	if payload.Updates[0].ComponentID != "c1" {
		t.Errorf("Updates[0].ComponentID: want c1 got %s", payload.Updates[0].ComponentID)
	}
	if payload.Updates[1].ComponentID != "c2" {
		t.Errorf("Updates[1].ComponentID: want c2 got %s", payload.Updates[1].ComponentID)
	}
}

func TestNewA2UIUpdate_noArgs_emptySlice(t *testing.T) {
	payload := a2ui.NewA2UIUpdate()
	if payload.Updates == nil {
		t.Error("Updates: want empty slice got nil")
	}
	if len(payload.Updates) != 0 {
		t.Errorf("Updates: want 0 got %d", len(payload.Updates))
	}
}
