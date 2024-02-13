package router

import (
	"go-mygram/controllers"
	"go-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	// users route
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}
	
	// PROTECTED ROUTE with middleware.Authentication()
	// photos route
	photoRouter := r.Group("/photos", middlewares.Authentication())
	{
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:photoID", controllers.GetOnePhoto)
		photoRouter.POST("/", controllers.CreatePhoto)

		// authorized routes
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	// comments route
	commentRouter := r.Group("/comments", middlewares.Authentication())
	{
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:commentID", controllers.GetOneComment)
		commentRouter.POST("/", controllers.CreateComment)
		
		// authorized routes
		commentRouter.PUT("/:commentID", controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", controllers.DeleteComment)	
	}
	
	// social media route
	socialMediaRouter := r.Group("/social-media", middlewares.Authentication())
	{
		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:sosmedID", controllers.GetOneSocialMedia)
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.PUT("/:sosmedID", controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:sosmedID", controllers.DeleteSocialMedia)	
	}

	return r
}
