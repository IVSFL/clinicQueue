package routes

import (
	"clinicQueue/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ws", controllers.WSHandler)

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
		doctorRoutes.PUT("/:id/office", controllers.UpdateDoctorOffice)
		doctorRoutes.GET("/specialization/:specialization_id", controllers.GetDoctorsBySpecialization)
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
		queueRoutes.POST("/:id/call-next", controllers.CallNext)
		queueRoutes.POST("/:id/call/:patient_id", controllers.CallList)
		queueRoutes.POST("/:id/call-deferred/:patient_id", controllers.CallDeferredPatient)
		queueRoutes.POST("/defer/:ticket_number", controllers.DeferPatientForTicket)
		queueRoutes.GET("/:id/deferred", controllers.GetDoctorDeferredQueue)
		queueRoutes.GET("/:id", controllers.GetDoctorQueue)
		queueRoutes.POST("/complete/:ticket_number", controllers.CompletePatient)
		queueRoutes.POST("/transfer/:ticket_number", controllers.TransferPatient)
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
