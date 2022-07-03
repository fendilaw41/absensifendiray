package action

import (
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func HandleImageAktifitas(res *gin.Context, fileName string, file *multipart.FileHeader) (string, error) {
	dirPath := filepath.Join(".", "storage", "aktifitas")
	filePath := filepath.Join(dirPath, fileName)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModeDir)
		if err != nil {
			CustomError("Terjadi Error Direktori", res)
		}
	}
	if errUpload := res.SaveUploadedFile(file, filePath); errUpload != nil {
		CustomError("Gagal Upload Foto", res)
	}
	return filePath, nil
}
