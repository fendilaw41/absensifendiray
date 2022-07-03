package seeds

import (
	"github.com/fendilaw41/absensifendiray/app/src/aktifitasImage"
)

func (s Seed) AktifitasImageSeed() {

	AktifitasSeed := []aktifitasImage.AktifitasImage{
		{
			AktifitasId:  1,
			Filename:     "Picture1.jpg",
			FilePath:     `\storage\aktifitas\Picture1.jpg`,
			OriginalName: "Picture1.jpg",
			FileSize:     18585,
		},
		{
			AktifitasId:  1,
			Filename:     "Picture2.jpg",
			FilePath:     `\storage\aktifitas\Picture2.jpg`,
			OriginalName: "Picture2.jpg",
			FileSize:     18585,
		},
	}

	s.db.Create(&AktifitasSeed)
}
