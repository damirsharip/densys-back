package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	ImageContentTypes = "image/jpeg,image/png,image/gif,image/webp,image/jpg,image/bmp"
)

func (s *Service) imageUpload(ctx context.Context, b64EncodedImg *string) (string, error) {
	if b64EncodedImg == nil {
		return "", errors.New("[Service.imageUpload] image is empty")
	}
	// get content type from strings split
	rawContentType := strings.Split(*b64EncodedImg, ";")[0]
	// get content type from strings split

	contentType := strings.Split(rawContentType, ":")[1]

	if !strings.Contains(ImageContentTypes, contentType) {
		fmt.Println("rawContentType", contentType)
		return "", errors.New("[Service.imageUpload] invalid image type")
	}

	encodedImage := strings.Split(*b64EncodedImg, "base64,")[1]

	decodedImage, errD := base64.StdEncoding.DecodeString(encodedImage)

	if errD != nil {
		fmt.Println("errD", errD)
		return "", errors.Wrap(errD, "[Service.imageUpload] failed to decode base64 image")
	}

	fileName := uuid.New().String() + "." + strings.Split(contentType, "/")[1]

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")))

	if err != nil {
		fmt.Println("unable to load SDK config, " + err.Error())
		return "", errors.Wrap(err, "[Service.imageUpload] failed to load aws config")
	}

	uploader := manager.NewUploader(s3.NewFromConfig(cfg))

	awsFile := s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(decodedImage),
		ContentType: aws.String(contentType),
	}

	object, err := uploader.Upload(ctx, &awsFile)

	if err != nil {
		fmt.Println(err, "failed to upload file")
		return "", errors.Wrap(err, "[Service.imageUpload] failed to upload file")
	}

	if object == nil {
		return "", errors.New("unexpected error after file uploading")
	}

	return object.Location, nil
}
