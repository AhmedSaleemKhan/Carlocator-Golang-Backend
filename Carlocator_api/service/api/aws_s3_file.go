package api

import (
	"bytes"
	"fmt"
	"math/rand"
	"mime/multipart"
	"time"

	config "capregsoft.com/carlocator/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func GenerateCode(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(max-min+1) + min
	return code
}

func (api *CarLocatorApiImpl) S3UploadFileAPI(c *gin.Context, file *multipart.FileHeader) (*string, error) {
	bucketName := config.Cfg.PrivateBucket
	session := api.awssession
	fileToUpload, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileToUpload.Close()
	fileBuffer := make([]byte, file.Size)
	_, err = fileToUpload.Read(fileBuffer)
	if err != nil {
		return nil, err
	}
	getRandomNumber := fmt.Sprint(GenerateCode(100000, 999999))
	key := getRandomNumber + file.Filename
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(fileBuffer),
	})

	if err != nil {
		return nil, err
	}
	filepath := fmt.Sprintf("https://%v.s3.%v.amazonaws.com/%v", bucketName, config.Cfg.Region, key)
	return &filepath, nil
}
