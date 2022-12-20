package database

import "errors"

var (
	ErrCantFindProducts   = errors.New("cant find products")
	ErrCantDecodeProducts = errors.New("cant find the product")
	ErrUserIdIsNotValid   = errors.New("User is not valid")
	ErrCantUpdateUser     = errors.New("Cannot add product to Cart")
	ErrCantRemoveItemCart = errors.New("Cannot remove this item from cart")
	ErrCantGetItem        = errors.New("Unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("Cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
