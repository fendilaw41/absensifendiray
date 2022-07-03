package seeds

import "github.com/fendilaw41/absensifendiray/app/src/divisi"

func (s Seed) DivisiSeed() {

	Divisi := []divisi.Divisi{
		{
			Name:      "Backend Developer",
			CreatedBy: 1,
			UpdatedBy: 1,
		},
		{
			Name:      "Frontend Developer",
			CreatedBy: 1,
			UpdatedBy: 1,
		},
	}

	s.db.Create(&Divisi)
}
