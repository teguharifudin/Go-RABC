package main

// load required packages
import (
	"fmt"
	"log"
	"macbookpro/go-rbac/controller"
	"macbookpro/go-rbac/database"
	"macbookpro/go-rbac/model"
	"macbookpro/go-rbac/util"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	loadDatabase()
	// start the server
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

// run database migrations and add seed data
func loadDatabase() {
	database.InitDb()
	database.Db.AutoMigrate(&model.Role{})
	database.Db.AutoMigrate(&model.User{})
	seedData()
}

func serveApplication() {
	router := gin.Default()
	authRoutes := router.Group("/auth/user")
	// registration route
	authRoutes.POST("/register", controller.Register)
	// login route
	authRoutes.POST("/login", controller.Login)
	// jwt auth
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(util.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	adminRoutes.POST("/user/role", controller.CreateRole)
	adminRoutes.GET("/user/roles", controller.GetRoles)
	adminRoutes.PUT("/user/role/:id", controller.UpdateRole)

	router.Run(":6000")
	fmt.Println("Server running on port 6000")
}

// load seed data into the database
func seedData() {
	var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
	database.Db.Save(&roles)
	database.Db.Save(&user)
}
