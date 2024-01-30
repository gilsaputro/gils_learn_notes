package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("test")
	client, err := storage.NewClient(context.TODO(), option.WithEndpoint("http://localhost:4443/storage/v1/"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	const (
		bucketName = "doitpay-staging-private"
		fileKey    = "some_file.txt"
	)
	buckets, err := list(client, bucketName)
	if err != nil {
		log.Fatalf("failed to list: %v", err)
	}
	fmt.Printf("buckets: %+v\n", buckets)

	data, err := downloadFile(client, bucketName, fileKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contents of %s/%s: %s\n", bucketName, fileKey, data)

	// err = deleteFile(client, bucketName, fileKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = updateConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fileKeyNew := "/file/new-file.txt"
	newdata := []byte("This is the content of the new file.")

	err = createFile(client, bucketName, fileKeyNew, newdata)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	fmt.Printf("File created successfully: %s/%s\n", bucketName, fileKey)
}

func list(client *storage.Client, bucketName string) ([]string, error) {
	var objects []string
	it := client.Bucket(bucketName).Objects(context.Background(), &storage.Query{})
	for {
		oattrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, oattrs.Name)
	}
	return objects, nil
}

func downloadFile(client *storage.Client, bucketName, fileKey string) ([]byte, error) {
	reader, err := client.Bucket(bucketName).Object(fileKey).NewReader(context.TODO())
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func createFile(client *storage.Client, bucketName, fileKey string, data []byte) error {
	ctx := context.TODO()

	obj := client.Bucket(bucketName).Object(fileKey)
	wc := obj.NewWriter(ctx)

	_, err := wc.Write(data)
	if err != nil {
		return fmt.Errorf("createFile: unable to write data to bucket %q, file %q: %v", bucketName, fileKey, err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("createFile: unable to close bucket %q, file %q: %v", bucketName, fileKey, err)
	}

	return nil
}

func deleteFile(client *storage.Client, bucketName, fileKey string) error {
	return client.Bucket(bucketName).Object(fileKey).Delete(context.TODO())
}

// func updateConfig() error {
// 	changeExternalUrl := "http://localhost:8080/_internal/config"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodPut, changeExternalUrl, strings.NewReader("{\"externalUrl\": \"http://localhost:8080\"}"))
// 	req.Header.Add("Content-Type", "application/json")
// 	if err != nil {
// 		return err
// 	}
// 	_, err = client.Do(req)

// 	return err
// }
