package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) SignInAPI(c *gin.Context, body *models.SignInRequest) (*cognito.InitiateAuthOutput, error) {
	params := map[string]*string{
		"EMAIL":    aws.String(*body.Email),
		"PASSWORD": aws.String(*body.Password),
	}
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *body.Email, config.Cfg.AppClientId)
	params["SECRET_HASH"] = aws.String(secretHash)

	authTry := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(*params["EMAIL"]),
			"PASSWORD":    aws.String(*params["PASSWORD"]),
			"SECRET_HASH": aws.String(*params["SECRET_HASH"]),
		},
		ClientId: aws.String(config.Cfg.AppClientId),
	}
	cognitoClient := cognito.New(api.awssession)
	authResp, err := cognitoClient.InitiateAuth(authTry)
	if err != nil {
		return nil, err
	}
	return authResp, nil
}
