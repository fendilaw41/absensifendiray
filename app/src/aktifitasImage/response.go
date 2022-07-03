package aktifitasImage

type AktifitasImageResponse struct {
	Id           int    `json:"id"`
	AktifitasId  int    `json:"aktifitas_id"`
	Filename     string `json:"filename"`
	FilePath     string `json:"filepath"`
	OriginalName string `json:"original_name"`
	FileSize     uint   `json:"filesize"`
}
