package main

import (
	commentDelivery "cloud_payments/internal/comment/delivery"
	commentRepo "cloud_payments/internal/comment/repo"
	"cloud_payments/internal/comment/usecase"
	"cloud_payments/internal/middlewares"
	"cloud_payments/internal/models"
	userDelivery "cloud_payments/internal/user/delivery"
	"cloud_payments/internal/user/repo"
	usecase2 "cloud_payments/internal/user/usecase"
	"fmt"
	"github.com/centrifugal/gocent/v3"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

func main() {

	host, _ := os.Hostname()

	tablePrefix := ""

	if os.Getenv("DB_TABLE_PREFIX") != "" {
		tablePrefix = os.Getenv("DB_TABLE_PREFIX")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.Project{},
		&models.Room{},
		&models.User{},
		&models.Comment{},
		&models.Attachment{},
		&models.Vote{},
		&models.CommentHistory{},
		&models.UserAuthorization{},
		&models.UserBan{},
	)

	api := gocent.New(gocent.Config{
		Addr: os.Getenv("CENTRIFUGO_APIURL"),
		Key:  os.Getenv("CENTRIFUGO_APIKEY"),
	})

	centrifugoApi := usecase.CentrifugoAPI{
		Api: api,
	}

	commentRepo := &commentRepo.CommentRepository{
		Db: db,
	}

	userRepo := &repo.UserRepo{
		Db: db,
	}

	jwtManager := usecase2.JwtManager{
		HmacSecret:    []byte(os.Getenv("HMAC_SECRET")),
		Method:        jwt.SigningMethodHS256,
		ExpireSeconds: time.Duration(86400 * 7),
	}

	commentHandler := commentDelivery.CommentHandler{
		CommentRepo:   commentRepo,
		UserRepo:      userRepo,
		CentrifugoApi: &centrifugoApi,
	}

	userHandler := userDelivery.UserHandler{
		JwtManager: jwtManager,
		UserRepo:   userRepo,
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(gzip.Gzip(5))
	router.Use(middlewares.CORSMiddleware())

	type HelloResponse struct {
		Message string `json:"message"`
		Version string `json:"version"`
		Host    string `json:"host"`
	}
	router.GET("/", func(c *gin.Context) {

		data := HelloResponse{
			Message: "I'm CloudComments",
			Version: "1.0",
			Host:    host,
		}

		c.JSON(200, data)
	})

	//admin route
	admin := router.Group("/admin/api")

	admin.GET("/comments", commentHandler.CommentsHandler())
	admin.GET("/users", userHandler.GetUsers())

	// comment routes
	router.GET("/project/:id", commentHandler.GetCommentHandler())
	router.GET("/upload/url", commentHandler.GetUploadUrlHandler())
	router.PUT("/upload", commentHandler.GetUploadHandler())
	router.POST("/attachments/tenor", commentHandler.CreateTenorGIFAttachment())

	// user routes
	router.GET("/user/avatar", userHandler.GetSvgAvatarByName())
	router.POST("/user/register", userHandler.RegisterUserHandler())
	router.POST("/user/auth", userHandler.AuthUserHandler())

	// centrifugo routes
	router.Any("/centrifugo/connect", commentHandler.ConnectHandler())
	router.Any("/centrifugo/publish", commentHandler.PublishHandler())

	// static
	router.StaticFile("/demo", "./dist/index.html")
	router.StaticFile("/admin", "./dist/admin.html")
	router.StaticFile("/callback", "./web/blank.html")
	router.Static("/static", "./dist/static")
	router.Static("/storage", "./storage")

	pprof.Register(router)

	router.Run(":" + os.Getenv("PORT"))

}
