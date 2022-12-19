package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" bson:"fist_name" validate:"required,min=2,max=100"`
	Last_Name       *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Password        *string            `json:"password" bson:"password" validate:"required,min=2,max=100"`
	Email           *string            `json:"email" bson:"email" validate:"required,min=2,max=100"`
	Phone           *string            `json:"phone" bson:"phone"`
	Token           *string            `json:"token" bson:"token"`
	Refresh_Token   *string            `json:"refresh_token" bson:"refresh_token"`
	Created_At      time.Time          `json:"created_at" bson:"created_at"`
	Updated_At      time.Time          `json:"updated_at" bson:"updated_at"`
	User_ID         *string            `json:"user_id" bson:"user_id"`
	UserCart        []Product_User     `json:"user_cart" bson:"user_cart"`
	Address_Details []Address          `json:"address_details" bson:"address_details"`
	Order_Status    []Order            `json:"order_status" bson:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        *uint64            `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Product_User struct {
	Product_ID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        *uint64            `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"address_id"`
	House      *string            `json:"house" bson:"house"`
	Street     *string            `json:"street" bson:"street"`
	City       *string            `json:"city" bson:"city"`
	Pincode    *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []Product_User     `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          uint64             `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"image"`
}

type Payment struct {
	Digital bool
	COD     bool
}
