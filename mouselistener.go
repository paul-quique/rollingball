package rollingball

const (
	MousePressed = iota
	MouseReleased
	MouseMoved
)

var Listeners map[string]MouseListener = map[string]MouseListener{}

//Clickable should be implemented by any clickable ebiten item
type MouseListener interface {
	ID() string
	Update(e MouseEvent)
	IsHovered(e MouseEvent) bool
	IsFocused() bool
}

//ClickEvent gives information about the location of the event and its type (MousePress, MouseReleased)
type MouseEvent struct {
	Location *Point
	Type     int
}

//AddListener adds a listener to Listeners
func AddMouseListener(m MouseListener) {
	Listeners[m.ID()] = m
}

//RemoveListener removes a listener from Listeners
func RemoveMouseListener(id string) {
	delete(Listeners, id)
}

//At most one listener is focused so the function will
//return as soon as it has found and updated a listener which
//IsFocused() function returned true.
//If no listener is focused, it will do the exact same thing
//with the IsHovered() function
//If no listener is focused and no listener is hovered, the
//function will update zero listener and return false

//UpdateListeners updates AT MOST ONE listener, returns true if one is updated, false otherwise
func UpdateListeners(e MouseEvent) bool {
	for _, l := range Listeners {
		if l.IsFocused() {
			l.Update(e)
			return true
		}
	}

	for _, l := range Listeners {
		if l.IsHovered(e) {
			l.Update(e)
			return true
		}
	}

	return false
}
