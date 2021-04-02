package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/thiduzz/lenslocked.com/models"
	"io"
	"os"
)

type awsSession struct {
	bucket  string
	region  string
	session *session.Session
	s3      *s3manager.Uploader
}

func NewAWSConnection() *awsSession {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKey,
				secret,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return &awsSession{
		bucket: bucket,
		region: region,
		session:   sess,
		s3: s3manager.NewUploader(sess),
	}
}

func (awss *awsSession) S3PutObject(file io.Reader, filename string) (string, error) {

	_, err := awss.s3.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awss.bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return "", models.ErrUploadFailed
	}
	filepath := "https://" + awss.bucket + "." + "s3-" + awss.region + ".amazonaws.com/" + filename
	return filepath, nil
}