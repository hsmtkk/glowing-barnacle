package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const EnableLogKey = "AZURE_SDK_GO_LOGGING"
const EnableLogAll = "all"

func main() {
	os.Setenv(EnableLogKey, EnableLogAll)
	accountName := requiredEnv("ACCOUNT_NAME")
	accountKey := requiredEnv("ACCOUNT_KEY")
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatalf("failed to init credential; %s", err)
	}
	url := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)
	serviceClient, err := azblob.NewServiceClientWithSharedKey(url, credential, nil)
	if err != nil {
		log.Fatalf("failed to init client; %s", err)
	}
	container := serviceClient.NewContainerClient("test")
	blockBlob := container.NewBlockBlobClient("test.txt")
	ctx := context.Background()
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("failed to open file; %s", err)
	}
	defer file.Close()
	_, err = blockBlob.Upload(ctx, file, nil) // discard response
	if err != nil {
		log.Fatalf("failed to upload; %s", err)
	}
}

func requiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s environment variable must be defined", key)
	}
	return val
}
