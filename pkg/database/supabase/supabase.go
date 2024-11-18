package supabase

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type Interface interface {
	Upload(file *multipart.FileHeader, folderName string) (string, error)
	Delete(link string) error
}

type supabaseStorage struct {
	client *storage_go.Client
}

// Delete implements Interface3.
func (s *supabaseStorage) Delete(filePath string) error {
	// Memastikan path lengkap file yang ingin dihapus
	_, err := s.client.RemoveFile("foto-profil", []string{filePath})
	if err != nil {
		// Jika terjadi error, bisa jadi file tidak ada
		return fmt.Errorf("gagal menghapus file dengan path %s: %v", filePath, err)
	}

	return nil
}


func (s *supabaseStorage) Upload(file *multipart.FileHeader, folderName string) (string, error) {
	// Membuka file
	fileBody, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileBody.Close() // Pastikan file ditutup setelah digunakan

	// Mendapatkan nama bucket dan file
	bucket := os.Getenv("SUPABASE_BUCKET")
	if bucket == "" {
		return "", errors.New("bucket not defined")
	}

	// Membuat path lengkap untuk file (folder + file)
	fileName := file.Filename
	contentType := file.Header.Get("Content-Type")
	filePath := fmt.Sprintf("%s/%s", folderName, fileName)

	// Melakukan upload file ke Supabase
	_, err = s.client.UploadFile(bucket, filePath, fileBody, storage_go.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}

	// Mengambil URL publik untuk file yang di-upload
	url := s.client.GetPublicUrl(bucket, filePath).SignedURL
	return url, nil
}



func Init() Interface {
	storageClient := storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_TOKEN"), nil)
	return &supabaseStorage{
		client: storageClient,
	}
}
