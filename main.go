package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"myfinnplan/entity"
	"myfinnplan/handler"
	"myfinnplan/helper"
	"myfinnplan/repository"
	"myfinnplan/service"
)

func main() {

	// dsn := "root:@tcp(127.0.0.1:3306)/myfinnplan?charset=utf8mb4&parseTime=true&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	env := entity.SetEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", env.DB_HOST, env.DB_USER, env.DB_PASS, env.DB_NAME, env.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(entity.User{}, entity.UserAccount{}, entity.BankAccount{}, entity.TrxCategory{}, entity.Transaction{})
	//Repository Region
	userAccountRepository := repository.NewUserAccountRepository(db)
	trxCategoryRepository := repository.NewTrxCategoryRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)
	bankAccountRepository := repository.NewBankAccountRepository(db)

	//Service Region
	userAccountService := service.NewUserAccountService(userAccountRepository)
	trxCategoryService := service.NewTrxCategoryService(trxCategoryRepository)
	transactionService := service.NewTransactionService(transactionRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)
	jwtService := service.NewJwtService()
	bankAccountService := service.NewBankAccountService(bankAccountRepository)
	bankService := service.NewBankService()

	//Handler Region
	userAccountHandler := handler.NewUserAccountHandler(userAccountService)
	trxCategoryHandler := handler.NewTrxCategoryHandler(trxCategoryService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService, jwtService, userService)
	bankAccountHandler := handler.NewBankAccountHandler(bankAccountService)
	bankHandler := handler.NewBankHandler(bankService)

	//Router Region
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api/v1")
	api.GET("/", func(ctx *gin.Context) {
		html, err := os.ReadFile("index.html")
		if err != nil {
			ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", html)
	})
	//auth
	api.POST("/createuserAccount", authMiddleware(jwtService, userService), userAccountHandler.CreateUserAccount)
	api.PUT("/edituserAccount", authMiddleware(jwtService, userService), userAccountHandler.EditUserAccount)
	api.GET("/getalluserAccounts", authMiddleware(jwtService, userService), userAccountHandler.GetAllUserAccounts)
	api.DELETE("/deleteuserAccount/:id", authMiddleware(jwtService, userService), userAccountHandler.DeleteUserAccount)
	api.GET("/getuserAccountbyid/:id", authMiddleware(jwtService, userService), userAccountHandler.GetUserAccountById)
	api.GET("/getuserAccountbyaccountcode/:accountCode", authMiddleware(jwtService, userService), userAccountHandler.GetUserAccountByAccountCode)
	api.GET("/getuserAccountbyaccountname/:accountName", authMiddleware(jwtService, userService), userAccountHandler.GetUserAccountByAccountName)
	api.GET("/getuserAccountbyuserid/:userId", authMiddleware(jwtService, userService), userAccountHandler.GetUserAccountByUserId)
	api.POST("/createtrxCategory", authMiddleware(jwtService, userService), trxCategoryHandler.CreateTrxCategory)
	api.PUT("/edittrxCategory", authMiddleware(jwtService, userService), trxCategoryHandler.EditTrxCategory)
	api.GET("/getalltrxCategorys", authMiddleware(jwtService, userService), trxCategoryHandler.GetAllTrxCategorys)
	api.DELETE("/deletetrxCategory/:id", authMiddleware(jwtService, userService), trxCategoryHandler.DeleteTrxCategory)
	api.GET("/gettrxCategorybyid/:id", authMiddleware(jwtService, userService), trxCategoryHandler.GetTrxCategoryById)
	api.GET("/gettrxCategorybycategorycode/:categoryCode", authMiddleware(jwtService, userService), trxCategoryHandler.GetTrxCategoryByCategoryCode)
	api.GET("/gettrxCategorybycategoryname/:categoryName", authMiddleware(jwtService, userService), trxCategoryHandler.GetTrxCategoryByCategoryName)
	api.POST("/createtransaction", authMiddleware(jwtService, userService), transactionHandler.CreateTransaction)
	api.PUT("/edittransaction", authMiddleware(jwtService, userService), transactionHandler.EditTransaction)
	api.GET("/getalltransactions", authMiddleware(jwtService, userService), transactionHandler.GetAllTransactions)
	api.GET("/deletetransaction/:id", authMiddleware(jwtService, userService), transactionHandler.DeleteTransaction)
	api.GET("/gettransactionbyid/:id", authMiddleware(jwtService, userService), transactionHandler.GetTransactionById)
	api.GET("/gettransactionbybankaccountid/:bankAccountId", authMiddleware(jwtService, userService), transactionHandler.GetTransactionByBankAccountId)
	api.GET("/gettransactionbycategorycode/:categoryCode", authMiddleware(jwtService, userService), transactionHandler.GetTransactionByCategoryCode)
	api.GET("/gettransactionbyamount/:amount", authMiddleware(jwtService, userService), transactionHandler.GetTransactionByAmount)
	api.GET("/gettransactionbynotes/:notes", authMiddleware(jwtService, userService), transactionHandler.GetTransactionByNotes)
	api.GET("/gettransactionbyuserid/:userId", authMiddleware(jwtService, userService), transactionHandler.GetTransactionByUserId)
	// api.POST("/createuser", authMiddleware(jwtService, userService), userHandler.CreateUser)
	api.PUT("/edituser", authMiddleware(jwtService, userService), userHandler.EditUser)
	api.PUT("/changepassword", authMiddleware(jwtService, userService), userHandler.ChangePassword)
	api.GET("/getallusers", authMiddleware(jwtService, userService), userHandler.GetAllUsers)
	api.GET("/deleteuser/:id", authMiddleware(jwtService, userService), userHandler.DeleteUser)
	api.POST("/register", authHandler.RegisterUser)
	api.POST("/login", authHandler.Login)
	api.GET("/getuserbyid/:id", authMiddleware(jwtService, userService), userHandler.GetUserById)
	api.GET("/getuserbyusername/:userName", authMiddleware(jwtService, userService), userHandler.GetUserByUserName)
	api.POST("/createbankAccount", authMiddleware(jwtService, userService), bankAccountHandler.CreateBankAccount)
	api.POST("/editbankAccount", authMiddleware(jwtService, userService), bankAccountHandler.EditBankAccount)
	api.GET("/getallbankAccounts", authMiddleware(jwtService, userService), bankAccountHandler.GetAllBankAccounts)
	api.GET("/deletebankAccount/:id", authMiddleware(jwtService, userService), bankAccountHandler.DeleteBankAccount)
	api.GET("/getbankAccountbyid/:id", authMiddleware(jwtService, userService), bankAccountHandler.GetBankAccountById)
	api.GET("/getbankAccountbyaccountcode/:accountCode", authMiddleware(jwtService, userService), bankAccountHandler.GetBankAccountByAccountCode)
	api.GET("/getbankAccountbyaccountidowner/:id", authMiddleware(jwtService, userService), bankAccountHandler.GetBankAccountByAccountIdOwner)
	api.GET("/getbankAccountbybankcode/:bankCode", authMiddleware(jwtService, userService), bankAccountHandler.GetBankAccountByBankCode)
	api.GET("/getbankAccountbyamount/:amount", authMiddleware(jwtService, userService), bankAccountHandler.GetBankAccountByAmount)
	api.GET("/getbankAccountbynotes/:notes", authMiddleware(jwtService, userService), bankAccountHandler.GetBankAccountByNotes)
	api.GET("/getbankdata", authMiddleware(jwtService, userService), bankHandler.GetBankData)
	api.GET("/verifiedaccount/:token", authHandler.VerifiedUser)
	api.GET("/resendemail/:id", authHandler.ReSendEmail)

	router.Run()

}

func authMiddleware(jwtService service.JwtService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.GetHeader("Authorization")

		if !strings.Contains(authheader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authheader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		dateTime, err := time.Parse(time.RFC3339Nano, claim["time"].(string))

		if err != nil {
			response := helper.APIResponse("Error Parse Date", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if dateTime.Before(time.Now()) {
			response := helper.APIResponse("Session End", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		users, err := userService.GetUserById(userId)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		user := users[0]

		c.Set("currentUser", user)
	}
}
