package aktifitas

type AktifitasRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	AssignId    int    `json:"assign_id" form:"assign_id" binding:"required"`
	Subject     string `json:"subject" form:"subject" binding:"required"`
	Description string `json:"description" form:"description"`
}

type AktifitasRequestPUT struct {
	Name        string `json:"name" form:"name"`
	AssignId    int    `json:"assign_id" form:"assign_id" `
	Subject     string `json:"subject" form:"subject"`
	Description string `json:"description" form:"description"`
}
