package mail

import (
	"bytes"
	"html/template"
	"os"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/model"
	mailgun "github.com/mailgun/mailgun-go"
)

func OrderNotify(
	orderID string,
	cart model.Cart,
	customer model.Customer,
	promotion model.Promotion) (string, error) {

	mg := mailgun.NewMailgun(DOMAIN, API_KEY, PUBLIC_API_KEY)
	emails := append(ADMIN_EMAILS, customer.Email)
	message := mg.NewMessage(
		os.Getenv("EMAIL_SENDER"),
		"Thank you for Your Order - "+orderID,
		"We've received your order.",
		emails...,
	)

	html := orderEmailHtml(orderID, cart, customer, promotion)
	message.SetHtml(html)

	_, id, err := mg.Send(message)
	return id, err
}

type OrderTmplData struct {
	OrderID   string
	Cart      model.Cart
	Customer  model.Customer
	Promotion model.Promotion
}

func orderEmailHtml(
	orderID string,
	cart model.Cart,
	customer model.Customer,
	promotion model.Promotion) string {

	var content bytes.Buffer
	t := template.New("email").Funcs(handler.BaseFuncMap())
	t = template.Must(t.ParseFiles("templates/emails/order.tmpl"))
	t.ExecuteTemplate(&content, "email", OrderTmplData{orderID, cart, customer, promotion})
	return content.String()
}
