package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/umarkotak/twitter_local/config"
	"github.com/umarkotak/twitter_local/controllers/ping_controller"
	"github.com/umarkotak/twitter_local/controllers/post_controller"
	"github.com/umarkotak/twitter_local/controllers/user_controller"
	_ "github.com/umarkotak/twitter_local/docs"
	"github.com/umarkotak/twitter_local/middlewares"
	formatter "github.com/umarkotak/twitter_local/utils/log_formatter"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	Twitter local API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:17000
//	@BasePath	/

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&formatter.Formatter{
		FieldsOrder: []string{"component", "category"},
	})

	// Load config
	err := config.Initialize()
	if err != nil {
		logrus.Fatalf("Error initializing config: %v", err)
	}

	// Initialize migration
	driver, err := postgres.WithInstance(config.Get().MasterDB.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		logrus.Fatal(err)
	}

	// Execute migration
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		logrus.Fatal(err)
	}

	// Initialize route
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middlewares.RequestID())
	r.Use(middlewares.LogRequest())

	rUser := r.Group("/", middlewares.AuthUser())

	r.GET("/ping", ping_controller.Ping)

	r.POST("/users/register", user_controller.Register)
	r.POST("/users/login", user_controller.Login)

	r.GET("/user/profile", user_controller.MyProfile)

	rUser.GET("/user/:username/profile", user_controller.ProfileByUsername)
	rUser.GET("/user/posts/:id")
	rUser.GET("/user/posts", post_controller.GetMyPosts)
	rUser.POST("/user/posts", post_controller.Create)
	rUser.PUT("/user/posts/:id")
	rUser.DELETE("/user/posts/:id")

	// docs: https://github.com/swaggo/swag
	// open: /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":17000")
}
