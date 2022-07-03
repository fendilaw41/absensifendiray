package main

import (
	"github.com/fendilaw41/absensifendiray/app/controllers"
	"github.com/fendilaw41/absensifendiray/config/database"
	"github.com/fendilaw41/absensifendiray/config/middleware"

	"github.com/gin-gonic/gin"
)

// test deploy
func main() {
	database.ConfigDB()
	// database.DbMigrateFreshSeed()
	route := gin.Default()
	api := route.Group("/api")

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
	api.GET("/departement", middleware.Auth(), controllers.GetAllDepartement)
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

	route.Run() // PORT 8080
}
