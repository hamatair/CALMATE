package supabase

import (
	"bytes"
	"encoding/json"
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

	// Step 1: List all files in the bucket or directory
	files, err := s.ListObjects(bucketName, filePath)
	if err != nil {
		return fmt.Errorf("gagal mengambil daftar file: %v", err)
	}

	// Step 2: Log files found
	if len(files) == 0 {
		fmt.Println("Tidak ada file yang ditemukan untuk dihapus.")
		return nil
	}

	fmt.Println("Daftar file yang akan dihapus:")
	for _, file := range files {
		fmt.Println(file)
	}

	// Step 3: Delete each file
	for _, file := range files {
		// URL endpoint Supabase untuk menghapus file
		url := fmt.Sprintf("%s/object/public/%s/%s", os.Getenv("SUPABASE_URL"), bucketName, file)

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
			fmt.Printf("File tidak ditemukan: %s (status code: 404), response: %s\n", file, string(body))
			continue
		}

		if resp.StatusCode != http.StatusNoContent {
			return fmt.Errorf("gagal menghapus file %s (status code: %d), response: %s", file, resp.StatusCode, string(body))
		}

		fmt.Printf("File %s berhasil dihapus dari bucket %s\n", file, bucketName)
	}

	return nil
}

// Fungsi tambahan untuk mengambil daftar file
func (s *supabaseStorage) ListObjects(bucketName, path string) ([]string, error) {
	// URL endpoint untuk list objek
	url := fmt.Sprintf("%s/object/public/%s", os.Getenv("SUPABASE_URL"), bucketName)

	// Membuat body request untuk list objek
	body := map[string]interface{}{
		"prefix":    path, // Path direktori, kosongkan jika ingin mengambil semua
		"delimiter": "/",  // Gunakan delimiter untuk memisahkan direktori
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat JSON body: %v", err)
	}

	// Membuat request HTTP POST
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat request HTTP: %v", err)
	}

	// Menambahkan header autentikasi
	req.Header.Set("apikey", os.Getenv("SUPABASE_APIKEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	// Eksekusi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gagal mengirim request ke Supabase: %v", err)
	}
	defer resp.Body.Close()

	// Membaca respons dari server
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca respons: %v", err)
	}

	// Mengecek status kode HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gagal mengambil daftar objek (status code: %d), response: %s", resp.StatusCode, string(bodyResp))
	}

	// Memproses respons JSON
	var result []map[string]interface{}
	if err := json.Unmarshal(bodyResp, &result); err != nil {
		return nil, fmt.Errorf("gagal memproses respons JSON: %v", err)
	}

	// Mengambil daftar file
	var files []string
	for _, obj := range result {
		if name, ok := obj["name"].(string); ok {
			files = append(files, name)
		}
	}

	return files, nil
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
