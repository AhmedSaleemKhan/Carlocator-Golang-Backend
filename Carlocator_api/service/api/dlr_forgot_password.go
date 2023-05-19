package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DealerForgotPasswordAPI(c *gin.Context, body *models.DealerForgotPasswordRequest) (*cognito.ForgotPasswordOutput, error) {
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *body.Email, config.Cfg.AppClientId)
	user := &cognito.ForgotPasswordInput{
		SecretHash: aws.String(secretHash),
		ClientId:   aws.String(config.Cfg.AppClientId),
		Username:   aws.String(*body.Email),
	}
	user.SecretHash = aws.String(secretHash)
	CognitoClient := cognito.New(api.awssession)
	resp, err := CognitoClient.ForgotPassword(user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
