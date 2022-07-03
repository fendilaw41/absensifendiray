package absensi

type AbsensiRequest struct {
	// TanggalAbsen string `json:"tanggal_absen" form:"tanggal_absen" binding:"required"`
	Hours   int64  `json:"hours" form:"hours" `
	Minutes int64  `json:"minutes" form:"minutes" `
	Seconds int64  `json:"seconds" form:"seconds" `
	Status  string `json:"status"  form:"status"`
}
