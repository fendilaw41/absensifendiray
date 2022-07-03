package seeds

import (
	"log"
	"reflect"

	"gorm.io/gorm"
)

type Seed struct {
	db *gorm.DB
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}
	seed(s, "DivisiSeed")
	seed(s, "DepartementSeed")
	seed(s, "UserSeed")
	seed(s, "AbsensiSeed")
	seed(s, "RoleSeed")
	seed(s, "UserRoleSeed") // untuk pivot
	seed(s, "PermissionSeed")
	seed(s, "RolePermissionSeed") // untuk pivot
	seed(s, "AktifitasSeed")
	seed(s, "AktifitasImageSeed")
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "success")
}
