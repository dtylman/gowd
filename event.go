package gowd

//EventElement represents the DOM element sending an event
type EventElement struct {
	Properties map[string]string `json:"properties"`
}

//Event represents a DOM event
type Event struct {
	Name   string         `json:"name"`
	Sender EventElement   `json:"sender"`
	Inputs []EventElement `json:"inputs"`
}

//EventHandler handler for DOM event.
type EventHandler func(sender *Element, event *EventElement)

const (
	//OnClick onclick event
	OnClick = "onclick"
	//OnChange onchange event
	OnChange = "onchange"
	//OnKeyPress onkeypress event
	OnKeyPress = "onkeypress"
)

//GetID get the id of the event sender.
func (e *EventElement) GetID() string {
	id, _ := e.Properties["id"]
	return id
}

//GetValue gets the value of the event sender.
func (e *EventElement) GetValue() string {
	id, _ := e.Properties["value"]
	return id
}
