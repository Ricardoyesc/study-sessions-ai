package a2ui

type DiffUpdate struct {
	ComponentID string                 `json:"componentId"`
	Props       map[string]interface{} `json:"props"`
}

type A2UIUpdatePayload struct {
	Updates []DiffUpdate `json:"updates"`
}

func NewA2UIUpdate(updates ...DiffUpdate) A2UIUpdatePayload {
	return A2UIUpdatePayload{Updates: updates}
}

func DiffProps(id string, props map[string]interface{}) DiffUpdate {
	return DiffUpdate{ComponentID: id, Props: props}
}
