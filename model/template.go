package model

// Global Template Data
type GTD struct {
	QueryText string
	Title     string
	Cart      Cart
}

type Tpl struct {
	QueryText string
	Title     string
	Cart      Cart
}

// Set Template Data from session
// notes: delete session by 'delete(session.Values, key)'
func (gtd *GTD) InitWithSession(session *Session) {
	if session == nil {
		return
	}
	cart, ok := session.Values["cart"].(Cart)
	if ok {
		gtd.Cart = cart
	}
}
