package routes

import (
	"clinicQueue/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	patientRoutes := r.Group("/patients")
	{
		patientRoutes.GET("", controllers.GetAllPatients)
		patientRoutes.POST("", controllers.CreatePatient)
		patientRoutes.GET("/:id", controllers.GetPatient)
	}

	doctorRoutes := r.Group("/doctors")
	{
		doctorRoutes.GET("/", controllers.GetAllDoctors)
		doctorRoutes.POST("/", controllers.CreateDoctor)
		doctorRoutes.GET("/:id", controllers.GetDoctor)
		doctorRoutes.PUT("/:id", controllers.UpdateDoctor)
	}

	adminRoutes := r.Group("/admins")
	{
		adminRoutes.POST("/", controllers.CreateAdmin)
		adminRoutes.GET("/:id", controllers.GetAdmin)
		adminRoutes.PUT("/:id", controllers.UpdateAdmin)
	}

	ticketRoutes := r.Group("/tickets")
	{
		//ticketRoutes.GET("/", controllers.GetAllTickets)
		ticketRoutes.POST("", controllers.CreateTicket)
		// ticketRoutes.GET("/:id", controllers.GetTicket)
		// ticketRoutes.PUT("/:id", controllers.UpdateTicket)
		// ticketRoutes.DELETE("/:id", controllers.DeleteTicket)
		// ticketRoutes.GET("/active", controllers.ActiveTicket)
		// ticketRoutes.GET("/progress", controllers.ProgressTicket)
	}

	// userRoutes := r.Group("/users")
	// {
	// 	userRoutes.GET("/", controllers.GetAllUsers)
	// 	userRoutes.POST("/", controllers.CreateUser)
	// 	userRoutes.GET("/:id", controllers.GetUser)
	// 	userRoutes.PUT("/:id", controllers.UpdateUser)
	// }

	queueRoutes := r.Group("/queue")
	{
		//queueRoutes.GET("/", controllers.GetAllQueue)
		//queueRoutes.POST("/:type", controllers.AddQueue)
		queueRoutes.GET("/:id", controllers.GetDoctorQueue)
	}

	registerRoutes := r.Group("/register")
	{
		registerRoutes.POST("/doctor", controllers.RegisterDoctor)
		registerRoutes.POST("/admin", controllers.RegisterAdmin)
		registerRoutes.POST("/login", controllers.Login)
	}

	specializationRoutes := r.Group("/specialization")
	{
		specializationRoutes.GET("", controllers.GetAllSpecialization)
	}
}
