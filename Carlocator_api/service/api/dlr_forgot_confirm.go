package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DealerForgotConfirmAPI(c *gin.Context, body *models.DealerForgotConfirmRequest) error {
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *body.Email, config.Cfg.AppClientId)
	user := &cognito.ConfirmForgotPasswordInput{
		SecretHash:       aws.String(secretHash),
		ClientId:         aws.String(config.Cfg.AppClientId),
		Username:         aws.String(*body.Email),
		ConfirmationCode: aws.String(*body.OTP),
		Password:         aws.String(*body.NewPassword),
	}
	user.SecretHash = aws.String(secretHash)
	CognitoClient := cognito.New(api.awssession)
	_, err := CognitoClient.ConfirmForgotPassword(user)
	if err != nil {
		return err
	}
	return nil
}
