package seeds

import (
	"github.com/fendilaw41/absensifendiray/app/src/departement"
)

func (s Seed) DepartementSeed() {

	depart := []departement.Departement{
		{
			Name:      "IT",
			CreatedBy: 1,
		},
		{
			Name:      "Marketing",
			CreatedBy: 1,
		},
	}

	s.db.Create(&depart)
}
