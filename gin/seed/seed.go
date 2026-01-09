package seed

import (
	"log"
	"time"

	"github.com/unifuu/ditto2/gin/model/user"
	userSvc "github.com/unifuu/ditto2/gin/svc/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// SeedData initializes the database with default data
func SeedData() {
	log.Println("Starting database seeding...")

	// Seed admin user
	seedAdminUser()

	log.Println("Database seeding completed!")
}

func seedAdminUser() {
	userService := userSvc.NewService()

	// Check if admin user already exists
	existingUser := userService.ByUsername("admin")
	if len(existingUser.Username) > 0 {
		log.Println("Admin user already exists, skipping creation")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return
	}

	// Create admin user
	adminUser := user.User{
		ID:        primitive.NewObjectID(),
		Username:  "admin",
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert admin user
	err = userService.SignUp(adminUser)
	if err != nil {
		log.Printf("Error creating admin user: %v", err)
		return
	}

	log.Println("Admin user created successfully (username: admin, password: admin)")
}
