package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	AWS_REGION string = "us-east-1"
)

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GetEC2Client() (*ec2.EC2, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(AWS_REGION)}, nil)
	if err != nil {
		return nil, err
	}
	return ec2.New(sess), nil
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GenerateUniqueFilename(suffix string) (string, error) {
	dir, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	filename := fmt.Sprintf("%s%s%s", dir, strings.ToLower(suffix), strconv.Itoa(int(rand.Int63())))
	return filename, nil
}