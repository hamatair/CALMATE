package supabase

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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
// func (s *supabaseStorage) Delete(filePath string) error {
// 	// Memastikan path lengkap file yang ingin dihapus
// 	_, err := s.client.RemoveFile("foto-profil", []string{filePath})
//     fmt.Println(filePath)
// 	if err != nil {
// 		// Jika terjadi error, bisa jadi file tidak ada
// 		return fmt.Errorf("gagal menghapus file dengan path %s: %v", filePath, err)
// 	}

// 	return nil
// }

func (s *supabaseStorage) Delete(filePath string) error {
	bucketName := os.Getenv("SUPABASE_BUCKET")
	if bucketName == "" {
		return errors.New("bucket name is not defined")
	}

	// URL endpoint Supabase untuk menghapus file
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", os.Getenv("SUPABASE_URL"), bucketName, filePath)

	// Membuat request HTTP DELETE
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("gagal membuat request HTTP: %v", err)
	}

	// Menambahkan header autentikasi
	req.Header.Set("apikey", os.Getenv("SUPABASE_APIKEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_TOKEN"))

	// Eksekusi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirim request ke Supabase: %v", err)
	}
	defer resp.Body.Close()

	// Membaca isi body respons
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("gagal membaca respons dari Supabase: %v", err)
	}

	// Mengecek status kode HTTP
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("file tidak ditemukan: %s (status code: 404), response: %s", filePath, string(body))
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("gagal menghapus file %s (status code: %d), response: %s", filePath, resp.StatusCode, string(body))
	}

	fmt.Printf("File %s berhasil dihapus dari bucket %s\n", filePath, bucketName)
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
	storageClient := storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_APIKEY"), nil)
	return &supabaseStorage{
		client: storageClient,
	}
}
