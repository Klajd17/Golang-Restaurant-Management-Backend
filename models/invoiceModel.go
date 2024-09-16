package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=credit_card|eq=debit_card|eq=paypal|eq=CARD|eq=CASH"`
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Create_at        time.Time          `json:"create_at"`
	Update_at        time.Time          `json:"update_at"`
}
