package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"bitsCarPool_back/internal/crud/notifications"
	"bitsCarPool_back/internal/crud/trips"
	"bitsCarPool_back/internal/database"
	"bitsCarPool_back/internal/models"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	//trips
	r.POST("/trips", s.createTripsHandler)
	r.GET("/trips", s.searchTripsHandler)

	//notfications
	r.POST("/notification", s.createNotificationHandler)
	r.POST("/notification_response", s.notificationResponeHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) createTripsHandler(c *gin.Context) {
	var trip_body models.Trip
	if err := c.ShouldBindJSON(&trip_body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	id, err := trip.NewTripService(database.New().GetClient()).CreateTrip(&trip_body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"id": id})
}

func (s *Server) searchTripsHandler(c *gin.Context){
	var search models.TripSearch
	if err := c.ShouldBindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trips, err := trip.NewTripService(database.New().GetClient()).SearchTrips(search)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trips": trips})
}

func (s *Server) createNotificationHandler(c *gin.Context){
	var notification_body models.NotificationModel
	
	if err := c.ShouldBindJSON(&notification_body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := notification.NewNotificationService(database.New().GetClient()).CreateNotification(&notification_body)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (s *Server) notificationResponeHandler(c *gin.Context){
	var update_notification models.UpdateNotificationModel
	
	if err := c.ShouldBindJSON(&update_notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := notification.NewNotificationService(database.New().GetClient()).UpdateNotification(&update_notification)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{"id": msg})

}