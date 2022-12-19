package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/ecommerce_web_app/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.Admin_View_Product())
	incomingRoutes.GET("/users/productsview", controllers.Search_Products())
	incomingRoutes.GET("/users/search", controllers.Search_Product_By_Query())
	incomingRoutes.GET("/addtocart", controllers.Add_To_Cart())
	incomingRoutes.GET("/removeitem", controllers.Remove_Item())
	incomingRoutes.GET("/cartcheckout", controllers.Buy_From_Cart())
	incomingRoutes.GET("/instantbuy", controllers.Instant_Buy())
}
