package seeds

import (
	"time"

	"github.com/fendilaw41/absensifendiray/app/src/absensi"

	"gorm.io/datatypes"
)

func (s Seed) AbsensiSeed() {

	absensi := []absensi.Absensi{
		{
			TanggalAbsen: datatypes.Date(time.Now()),
			JamAbsen:     datatypes.NewTime(10, 30, 00, 00),
			UserId:       1,
			FirstName:    "Muhamad",
			LastName:     "Efendy",
			FullName:     "Muhamad Efendy Ray",
			Picture:      "default.jpg",
			CheckAbsen:   "Check-IN",
			Status:       "Terlambat",
			CreatedBy:    1,
			UpdatedBy:    1,
		},
		{
			TanggalAbsen: datatypes.Date(time.Now()),
			JamAbsen:     datatypes.NewTime(10, 30, 00, 00),
			UserId:       1,
			FirstName:    "Muhamad",
			LastName:     "Efendy",
			FullName:     "Muhamad Efendy Ray",
			Picture:      "default.jpg",
			CheckAbsen:   "Check-OUT",
			Status:       "Terlambat",
			CreatedBy:    1,
			UpdatedBy:    1,
		},
		{
			TanggalAbsen: datatypes.Date(time.Now()),
			JamAbsen:     datatypes.NewTime(10, 30, 00, 00),
			UserId:       3,
			FirstName:    "Muhamad",
			LastName:     "Keenan",
			FullName:     "Muhamad Keenan Athariz",
			Picture:      "default.jpg",
			CheckAbsen:   "Check-IN",
			Status:       "Terlambat",
			CreatedBy:    1,
			UpdatedBy:    1,
		},
	}

	s.db.Create(&absensi)
}
