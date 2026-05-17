package internal

import (
	"booking-system/lvl1/internal/config"
	"booking-system/lvl1/internal/handler"
	"booking-system/lvl1/internal/middleware"
	"booking-system/lvl1/internal/repository"
	"booking-system/lvl1/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, db *config.DB) {
	router.Use(middleware.CORS())
	api := router.Group("/api")
	{
		AuthRoutes(api, db)
		UserRoutes(api, db)
		HotelRoutes(api, db)
		// BookingRoutes(api, db)
	}

}

func AuthRoutes(router *gin.RouterGroup, db *config.DB) {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/refresh", authHandler.RefreshToken)

	}
}

func UserRoutes(router *gin.RouterGroup, db *config.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	user := router.Group("/users")
	{
		// user.GET("/", userHandler.GetUsers)
		user.POST("", userHandler.CreateUser)
		// user.PUT("/:id", userHandler.UpdateUser)
		// user.DELETE("/:id", userHandler.DeleteUser)
	}
}

func HotelRoutes(router *gin.RouterGroup, db *config.DB) {
	hotelRepo := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepo)
	hotelHandler := handler.NewHotelHandler(hotelService)

	roomRepo := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	hotel := router.Group("/hotels")
	{
		hotel.GET("", hotelHandler.GetAllHotels)
		hotel.GET("/:hotelId", hotelHandler.GetHotelByID)

		protected := hotel.Group("")
		protected.Use(middleware.Auth())
		{
			protected.POST("", hotelHandler.CreateHotel)
			protected.PUT("/:hotelId", hotelHandler.UpdateHotel)

			// rooms nested under hotels
			rooms := protected.Group("/:hotelId/rooms")
			{
				rooms.POST("", roomHandler.CreateRoom)
				rooms.GET("", roomHandler.GetRoomsByHotel)
				rooms.GET("/:roomId", roomHandler.GetRoomByID)
				rooms.PUT("/:roomId", roomHandler.UpdateRoom)
				rooms.DELETE("/:roomId", roomHandler.DeleteRoom)
			}
		}
	}
}
