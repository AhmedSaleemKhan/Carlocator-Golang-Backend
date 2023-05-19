package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) ResendOTPAPI(c *gin.Context, body *models.ResendOTPRequest) (*string, error) {
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *body.Email, config.Cfg.AppClientId)

	user := &cognito.ResendConfirmationCodeInput{
		SecretHash: aws.String(secretHash),
		Username:   aws.String(*body.Email),
		ClientId:   aws.String(config.Cfg.AppClientId),
	}
	user.SecretHash = aws.String(secretHash)
	CognitoClient := cognito.New(api.awssession)
	_, err := CognitoClient.ResendConfirmationCode(user)
	if err != nil {
		return nil, err
	}

	return user.Username, nil
}
