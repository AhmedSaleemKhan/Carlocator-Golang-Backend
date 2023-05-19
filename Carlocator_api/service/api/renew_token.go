package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) RenewTokenAPI(c *gin.Context, body *models.RefreshAccessTokenRequest) (*cognito.InitiateAuthOutput, error) {
	params := map[string]*string{
		"REFRESH_TOKEN": aws.String(*body.Token),
		"SECRET_HASH":   aws.String(config.Cfg.AppClientSecret),
	}
	authTry := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(*params["REFRESH_TOKEN"]),
			"SECRET_HASH":   aws.String(*params["SECRET_HASH"]),
		},
		ClientId: aws.String(config.Cfg.AppClientId),
	}
	CognitoClient := cognito.New(api.awssession)
	authResp, err := CognitoClient.InitiateAuth(authTry)
	if err != nil {
		return nil, err
	}
	return authResp, nil
}
