package a2ui_test

import (
	"testing"

	"sai-server/pkg/a2ui"
)

func TestNewText_setsTypeAndProps(t *testing.T) {
	got := a2ui.NewText("t1", "Hello", "h1")
	if got.ID != "t1" {
		t.Errorf("ID: want t1 got %s", got.ID)
	}
	if got.Type != a2ui.ComponentTypeText {
		t.Errorf("Type: want Text got %s", got.Type)
	}
	if got.Props["content"] != "Hello" {
		t.Errorf("content: want Hello got %v", got.Props["content"])
	}
	if got.Props["variant"] != "h1" {
		t.Errorf("variant: want h1 got %v", got.Props["variant"])
	}
}

func TestNewRichText_setsMarkdownAndAccessible(t *testing.T) {
	got := a2ui.NewRichText("r1", "**bold**")
	if got.Type != a2ui.ComponentTypeRichText {
		t.Errorf("Type: want RichText got %s", got.Type)
	}
	if got.Props["markdown"] != "**bold**" {
		t.Errorf("markdown: want **bold** got %v", got.Props["markdown"])
	}
	if got.Props["accessible"] != true {
		t.Errorf("accessible: want true got %v", got.Props["accessible"])
	}
}

func TestNewImage_setsUrlAndAltText(t *testing.T) {
	got := a2ui.NewImage("i1", "https://example.com/img.png", "a cat")
	if got.Type != a2ui.ComponentTypeImage {
		t.Errorf("Type: want Image got %s", got.Type)
	}
	if got.Props["url"] != "https://example.com/img.png" {
		t.Errorf("url mismatch")
	}
	if got.Props["altText"] != "a cat" {
		t.Errorf("altText: want 'a cat' got %v", got.Props["altText"])
	}
}

func TestNewAudioPlayer_setsUrlAndAutoPlayFalse(t *testing.T) {
	got := a2ui.NewAudioPlayer("a1", "https://example.com/audio.mp3")
	if got.Type != a2ui.ComponentTypeAudioPlayer {
		t.Errorf("Type: want AudioPlayer got %s", got.Type)
	}
	if got.Props["url"] != "https://example.com/audio.mp3" {
		t.Errorf("url mismatch")
	}
	if got.Props["autoPlay"] != false {
		t.Errorf("autoPlay: want false got %v", got.Props["autoPlay"])
	}
}

func TestNewColumn_nilProps_setsDefaults(t *testing.T) {
	got := a2ui.NewColumn("col1", []string{"c1", "c2"}, nil)
	if got.Type != a2ui.ComponentTypeColumn {
		t.Errorf("Type: want Column got %s", got.Type)
	}
	if got.Props["gap"] != 16 {
		t.Errorf("gap: want 16 got %v", got.Props["gap"])
	}
	if got.Props["padding"] != 24 {
		t.Errorf("padding: want 24 got %v", got.Props["padding"])
	}
	if len(got.Children) != 2 {
		t.Errorf("Children: want 2 got %d", len(got.Children))
	}
}

func TestNewColumn_customProps_usesProvided(t *testing.T) {
	custom := map[string]interface{}{"gap": 8}
	got := a2ui.NewColumn("col2", nil, custom)
	if got.Props["gap"] != 8 {
		t.Errorf("gap: want 8 got %v", got.Props["gap"])
	}
}

func TestNewRow_nilProps_setsDefaults(t *testing.T) {
	got := a2ui.NewRow("row1", []string{"a"}, nil)
	if got.Type != a2ui.ComponentTypeRow {
		t.Errorf("Type: want Row got %s", got.Type)
	}
	if got.Props["alignment"] != "center" {
		t.Errorf("alignment: want center got %v", got.Props["alignment"])
	}
}

func TestNewCard_nilProps_setsDefaults(t *testing.T) {
	got := a2ui.NewCard("card1", []string{}, nil)
	if got.Type != a2ui.ComponentTypeCard {
		t.Errorf("Type: want Card got %s", got.Type)
	}
	if got.Props["elevation"] != 2 {
		t.Errorf("elevation: want 2 got %v", got.Props["elevation"])
	}
}

func TestNewQuizCard_setsQuestionOptionsModeAndOnSubmitEvent(t *testing.T) {
	opts := []string{"A", "B", "C"}
	got := a2ui.NewQuizCard("q1", "What is 2+2?", opts, "single", "/api/answer")
	if got.Type != a2ui.ComponentTypeQuizCard {
		t.Errorf("Type: want QuizCard got %s", got.Type)
	}
	if got.Props["question"] != "What is 2+2?" {
		t.Errorf("question mismatch")
	}
	if got.Props["mode"] != "single" {
		t.Errorf("mode: want single got %v", got.Props["mode"])
	}
	if got.Events["onSubmit"] != "/api/answer" {
		t.Errorf("onSubmit: want /api/answer got %v", got.Events["onSubmit"])
	}
}

func TestNewSocraticDialog_setsPromptContextAndOnSubmitEvent(t *testing.T) {
	got := a2ui.NewSocraticDialog("sd1", "Explain it simply", "context text", "/api/socratic")
	if got.Type != a2ui.ComponentTypeSocraticDialog {
		t.Errorf("Type: want SocraticDialog got %s", got.Type)
	}
	if got.Props["prompt"] != "Explain it simply" {
		t.Errorf("prompt mismatch")
	}
	if got.Events["onSubmit"] != "/api/socratic" {
		t.Errorf("onSubmit mismatch")
	}
}

func TestNewProgressBar_setsValueAndMax(t *testing.T) {
	got := a2ui.NewProgressBar("pb1", 0.85, 1.0)
	if got.Type != a2ui.ComponentTypeProgressBar {
		t.Errorf("Type: want ProgressBar got %s", got.Type)
	}
	if got.Props["value"] != 0.85 {
		t.Errorf("value: want 0.85 got %v", got.Props["value"])
	}
	if got.Props["max"] != 1.0 {
		t.Errorf("max: want 1.0 got %v", got.Props["max"])
	}
}

func TestNewButton_setsLabelAndVariant(t *testing.T) {
	got := a2ui.NewButton("b1", "Submit", "primary")
	if got.Type != a2ui.ComponentTypeButton {
		t.Errorf("Type: want Button got %s", got.Type)
	}
	if got.Props["label"] != "Submit" {
		t.Errorf("label: want Submit got %v", got.Props["label"])
	}
	if got.Props["variant"] != "primary" {
		t.Errorf("variant: want primary got %v", got.Props["variant"])
	}
}

func TestDefaultDataModel_setsLanguageEsAndThemeSystem(t *testing.T) {
	got := a2ui.DefaultDataModel()
	if got.Language != "es" {
		t.Errorf("Language: want es got %s", got.Language)
	}
	if got.Theme != "system" {
		t.Errorf("Theme: want system got %s", got.Theme)
	}
	if got.FontScale != 1.0 {
		t.Errorf("FontScale: want 1.0 got %f", got.FontScale)
	}
}
