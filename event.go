package gowd

type EventElement struct {
	Properties map[string]string `json:"properties"`
}

type Event struct {
	Name   string         `json:"name"`
	Sender EventElement   `json:"sender"`
	Inputs []EventElement `json:"inputs"`
}

type EventHandler func(sender *Element, event *EventElement)

const (
	OnClick  = "onclick"
	OnChange = "onchange"
)

func (e *EventElement) GetID() string {
	id, _ := e.Properties["id"]
	return id
}

func (e *EventElement) GetValue() string {
	id, _ := e.Properties["value"]
	return id
}
