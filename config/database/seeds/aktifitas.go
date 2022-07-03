package seeds

import (
	"github.com/fendilaw41/absensifendiray/app/src/aktifitas"
)

func (s Seed) AktifitasSeed() {

	AktifitasSeed := []aktifitas.Aktifitas{
		{
			UserId:      1,
			AssignedId:  1,
			Subject:     "Meeting di PT ABC",
			Name:        "Meeting Klien",
			Description: "Presentasi Aplikasi",
			// CreatedAt:   action.FormatDate("2022-07-01"),
			CreatedBy: 1,
			UpdatedBy: 1,
		},
		{
			UserId:      2,
			AssignedId:  2,
			Subject:     "Meeting di PT Maju Mundur",
			Name:        "Meeting Kolega",
			Description: "Jalin Kerjasama",
			// CreatedAt:   action.FormatDate("2022-06-20"),
			CreatedBy: 1,
			UpdatedBy: 1,
		},
	}

	s.db.Create(&AktifitasSeed)
}
