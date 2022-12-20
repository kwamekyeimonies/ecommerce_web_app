package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/ecommerce_web_app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := Validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phoneNumber": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone Number has already been used"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.User_ID.Hex()
		token, refresh_token, _ := generate.TokenGenerator(*&user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refresh_token
		user.UserCart = make([]models.Product_User, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Account Creation unsucessfull...."})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, "Account Succesfully Created.....")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Product_View() {

}

func Search_Products() {

}

func Search_Product_By_Query() {

}
