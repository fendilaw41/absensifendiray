package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fendilaw41/absensifendiray/app/controllers"
	"github.com/fendilaw41/absensifendiray/config/database"
	"github.com/fendilaw41/absensifendiray/config/database/seeds"
	"github.com/fendilaw41/absensifendiray/config/middleware"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// test deploy
func main() {
	errenv := godotenv.Load(".env")
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	repeat, err := strconv.Atoi(tStr)
	if err != nil {
		log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
		repeat = 5
	}
	database.ConfigDB()
	// database.DbMigrateFreshSeed()
	route := gin.Default()

	route.GET("/repeat", repeatHandler(repeat))
	route.GET("/seeder", RunSeeder())
	route.GET("/migration", RunMigration())

	api := route.Group("/api")
	api.GET("/", controllers.Hello)

	api.POST("/login", controllers.Login)
	api.GET("/logout", controllers.Logout)
	api.POST("/register", controllers.Signup)
	api.GET("/profile", middleware.Auth(), controllers.Profile)
	// TODO: ROUTE User
	api.GET("/user", middleware.Auth(), controllers.GetAllUserPaginate)
	// api.GET("/user", middleware.Auth(), controllers.GetAllUser)
	api.GET("/user/:id", middleware.Auth(), controllers.GetDetailUser)
	api.POST("/user", middleware.Auth(), controllers.StoreUser)
	api.PUT("/user/:id", middleware.Auth(), controllers.UpdateUser)
	api.DELETE("/user/:id", middleware.Auth(), controllers.DeleteUser)
	// TODO: ROUTE Role
	api.GET("/role", middleware.Auth(), controllers.GetAllRolePaginate)
	api.GET("/role/:id", middleware.Auth(), controllers.GetDetailRole)
	api.POST("/role", middleware.Auth(), controllers.StoreRole)
	api.PUT("/role/:id", middleware.Auth(), controllers.UpdateRole)
	api.DELETE("/role/:id", middleware.Auth(), controllers.DeleteRole)
	// TODO: ROUTE Permission
	api.GET("/permission", middleware.Auth(), controllers.GetAllPermissionPaginate)
	api.GET("/permission/:id", middleware.Auth(), controllers.GetDetailPermission)
	api.POST("/permission", middleware.Auth(), controllers.StorePermission)
	api.PUT("/permission/:id", middleware.Auth(), controllers.UpdatePermission)
	api.DELETE("/permission/:id", middleware.Auth(), controllers.DeletePermission)
	// # Divisi
	api.GET("/divisi", middleware.Auth(), controllers.GetAllDivisi)
	api.GET("/divisi_paginate", middleware.Auth(), controllers.GetAllDivisiPaginate)
	api.GET("/divisi/:id", middleware.Auth(), controllers.GetDetailDivisi)
	api.POST("/divisi", middleware.Auth(), controllers.StoreDivisi)
	api.PUT("/divisi/:id", middleware.Auth(), controllers.UpdateDivisi)
	api.DELETE("/divisi/:id", middleware.Auth(), controllers.DeleteDivisi)
	// # Departement
	api.GET("/departement", controllers.GetAllDepartement)
	api.GET("/departement_paginate", middleware.Auth(), controllers.GetAllDepartementPaginate)
	api.GET("/departement/:id", middleware.Auth(), controllers.GetDetailDepartement)
	api.POST("/departement", middleware.Auth(), controllers.StoreDepartement)
	api.PUT("/departement/:id", middleware.Auth(), controllers.UpdateDepartement)
	api.DELETE("/departement/:id", middleware.Auth(), controllers.DeleteDepartement)
	// # Absensi
	api.GET("/absensi", middleware.Auth(), controllers.GetAllAbsensi)
	api.GET("/riwayat_absensi", middleware.Auth(), controllers.RiwayatAbsensi)
	api.GET("/absensi/:id", middleware.Auth(), controllers.GetByIdAbsensi)
	api.POST("/checkin", middleware.Auth(), controllers.CheckIn)
	api.POST("/checkout", middleware.Auth(), controllers.CheckOut)
	api.PUT("/absensi/:id", middleware.Auth(), controllers.UpdateAbsensi)
	api.DELETE("/absensi/:id", middleware.Auth(), controllers.DeleteAbsensi)
	// # Aktifitas
	api.GET("/aktifitas", middleware.Auth(), controllers.GetAllAktifitas)
	api.GET("/riwayat_aktifitas", middleware.Auth(), controllers.RiwayatAktifitas)
	api.GET("/aktifitas/:id", middleware.Auth(), controllers.GetDetailAktifitas)
	api.POST("/aktifitas", middleware.Auth(), controllers.StoreAktifitas)
	api.PUT("/aktifitas/:id", middleware.Auth(), controllers.UpdateAktifitas)
	api.DELETE("/aktifitas/:id", middleware.Auth(), controllers.DeleteAktifitas)

	route.Run(":" + port) // PORT 8080
}

func RunSeeder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		database.DbMigrateFresh()
		db, err := database.DbSetup()
		if err != nil {
			log.Fatalf("Error opening DB: %v", err)
		}
		seeds.Execute(db, "seeder sukses")
		fmt.Println("=======Seeders Success=======")
		ctx.JSON(200, "=======Migration Seeder table sukses=======")
	}
}

func RunMigration() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		database.DbMigrateFresh()
		ctx.JSON(200, "=======Migration table sukses=======")
	}
}

func repeatHandler(r int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer
		for i := 0; i < r; i++ {
			buffer.WriteString("Hello from Go!\n")
		}
		c.String(http.StatusOK, buffer.String())
	}
}
