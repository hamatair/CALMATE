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
	_, err := s.client.RemoveFile("foto-profil", []string{filePath})
	if err != nil {
		return err
	}

	return nil
}

func (s *supabaseStorage) Upload(file *multipart.FileHeader, folderName string) (string, error) {
    // Membuka file
	fmt.Println("1")
    fileBody, err := file.Open()
    if err != nil {
        return "", err
    }
    defer fileBody.Close() // Pastikan file ditutup setelah digunakan
	fmt.Println("2")

    // Mendapatkan nama bucket dan file
    bucket := os.Getenv("SUPABASE_BUCKET")
    if bucket == "" {
        return "", errors.New("bucket not defined")
    }
	fmt.Println("3")

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
	fmt.Println("4")

    // Mengambil URL publik
    url := s.client.GetPublicUrl(bucket, fileName).SignedURL
    return url, nil
}


func Init() Interface {
	storageClient := storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_TOKEN"), nil)
	return &supabaseStorage{
		client: storageClient,
	}
}
