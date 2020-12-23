package setget


type Event struct{
	title string
	name string
	Date
}

// Setter methods for fields of Event type
// With Date field, All the above methods on Date type are exported to Event type
func (e *Event) SetTitle(title string){
	e.title = title
}

func (e *Event) SetName(name string){
	e.name = name
}

// Getter Methods for fields of Event type
func (e *Event) Title() string{
	return e.title
}

func (e *Event) Name() string{
	return e.name
}