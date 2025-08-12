package main

// "github.com/ajonesb/user-management/backend/internal/controllers"
// "github.com/ajonesb/user-management/backend/internal/middleware"
// "github.com/ajonesb/user-management/backend/internal/repositories"
// "github.com/ajonesb/user-management/backend/internal/services"
// "github.com/ajonesb/user-management/backend/pkg/config"
// "github.com/ajonesb/user-management/backend/pkg/database"
// "github.com/gin-contrib/cors"
// "github.com/gin-gonic/gin"

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; reading environment variables.")
	}

	cfg := LoadConfigFromEnv()

	// Ensure upload dir exists
	if err := os.MkdirAll(cfg.UploadDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create upload dir: %v", err)
	}

	// DB
	db, err := InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	// Migrate
	if err := AutoMigrate(db); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// Create admin user if none exists (dev helper)
	EnsureAdminUser(db)

	// Router
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", RegisterHandler(db, cfg))
			auth.POST("/login", LoginHandler(db, cfg))
			auth.GET("/me", AuthMiddleware(cfg.JWTSecret, db), MeHandler(db))
		}

		showrooms := api.Group("/showrooms")
		{
			showrooms.GET("", ListShowroomsHandler(db))
			showrooms.GET("/:id", GetShowroomHandler(db))

			// protected
			showrooms.POST("", AuthMiddleware(cfg.JWTSecret, db), RoleMiddleware("admin"), UploadModelMiddleware(cfg), CreateShowroomHandler(db, cfg))
			showrooms.PATCH("/:id", AuthMiddleware(cfg.JWTSecret, db), RoleMiddleware("admin"), UploadModelMiddleware(cfg), UpdateShowroomHandler(db, cfg))
			showrooms.DELETE("/:id", AuthMiddleware(cfg.JWTSecret, db), RoleMiddleware("admin"), DeleteShowroomHandler(db, cfg))
		}

		admin := api.Group("/admin")
		{
			admin.Use(AuthMiddleware(cfg.JWTSecret, db), RoleMiddleware("admin"))
			admin.GET("/users", AdminListUsersHandler(db))
			admin.PATCH("/users/:id", AdminPatchUserHandler(db))
			admin.DELETE("/users/:id", AdminDeleteUserHandler(db))
			admin.GET("/bookings", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "not implemented in this example"}) })
		}
	}

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
