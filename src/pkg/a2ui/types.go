package a2ui

type MessageType string

const (
	MsgA2UIFull       MessageType = "a2ui_full"
	MsgA2UIUpdate     MessageType = "a2ui_update"
	MsgDataModelUpdate MessageType = "data_model_update"
	MsgError          MessageType = "error"
	MsgPing           MessageType = "ping"
	MsgPong           MessageType = "pong"
)

type DataModelUpdatePayload struct {
	Path  string                 `json:"path"`
	Value interface{}            `json:"value"`
	Diff  map[string]interface{} `json:"diff"`
}

type ErrorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	ComponentTypeText           = "Text"
	ComponentTypeRichText       = "RichText"
	ComponentTypeImage          = "Image"
	ComponentTypeAudioPlayer    = "AudioPlayer"
	ComponentTypeVideoPlayer    = "VideoPlayer"
	ComponentTypeCard           = "Card"
	ComponentTypeColumn         = "Column"
	ComponentTypeRow            = "Row"
	ComponentTypeQuizCard       = "QuizCard"
	ComponentTypeSocraticDialog = "SocraticDialog"
	ComponentTypeProgressBar    = "ProgressBar"
	ComponentTypeButton         = "Button"
)
