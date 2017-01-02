package main

import (
	"io/ioutil"
	"golang.org/x/net/context"
	"cloud.google.com/go/storage"
	
	"log"
	"time"
	"os"
	"io"

)


func UploadFileToBucket(ctx context.Context, pathOfImage string, bucket *storage.BucketHandle, finalName string) error {
	
	rawData, err := ioutil.ReadFile(pathOfImage)
	
	if err != nil {
		return err
	}
	data := bucket.Object(finalName)
	writer := data.NewWriter(ctx)
	n, err := writer.Write(rawData)
	if err != nil {
		return err
	}
	t1 := time.Now()
	writer.Close()
	t2 := time.Now();
	log.Printf("Uploaded %v bytes in %v seconds", n, (t2.Sub(t1)).Seconds())
	return nil
	
}

func DownloadFileFromBucket(ctx context.Context, pathOfImageInBucket string, bucket *storage.BucketHandle, pathForImage string) error {

	t1 := time.Now()
	t2 := time.Now()

	obj := bucket.Object(pathOfImageInBucket)
	reader, err := obj.NewReader(ctx)

	rawData := make([]byte, reader.Size())


	if err != nil {
		return err
	}
	defer func () {
		reader.Close()
		t2 = time.Now();
		log.Printf("Downloaded %v bytes in %v seconds", reader.Size(), (t2.Sub(t1)).Seconds())

	}()
	if _, err = io.ReadFull(reader, rawData); err != nil {
		return err
	}
	err = ioutil.WriteFile(pathForImage, rawData, os.ModeDir)
	if err != nil {
		return err

	}

	return nil
	
}


func UploadImage(ctx context.Context, pathOfImage, finalPathInStorage string) error {
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bucket := client.Bucket("detech-1e226.appspot.com")

	err = UploadFileToBucket(ctx, pathOfImage, bucket, finalPathInStorage)

	if err != nil {
		return err
	}

	return nil
}

func DownloadImage(ctx context.Context, pathOfImageInStorage, localPath string) error {

	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bucket := client.Bucket("detech-1e226.appspot.com")

	err = DownloadFileFromBucket(ctx, pathOfImageInStorage, bucket, localPath)

	if err != nil {
		return err
	}

	return nil

}


