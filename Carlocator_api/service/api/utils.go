package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"capregsoft.com/carlocator/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
)

func ConnectAws() *session.Session {
	MyRegion := config.Cfg.Region
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func ComputeSecretHash(clientSecret string, email string, clientId string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(email + clientId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func CreateUUID() string {
	id := uuid.New()
	return id.String()
}

func CheckRecordExists(newEmail, newUsername, existedEmail, existedUsername *string) error {
	if (len(*existedEmail) > 0 && *existedEmail == *newEmail) && (len(*existedUsername) > 0 && *existedUsername == *newUsername) {
		return fmt.Errorf("email and username already exist")
	} else if len(*existedEmail) > 0 && *existedEmail == *newEmail {
		return fmt.Errorf("email already exist")
	} else if len(*existedUsername) > 0 && *existedUsername == *newUsername {
		return fmt.Errorf("username already exist")
	}
	return nil
}
