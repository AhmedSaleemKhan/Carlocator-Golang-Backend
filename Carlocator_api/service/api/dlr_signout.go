package api

import (
	"capregsoft.com/carlocator/service/models"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DealerSignOutAPI(c *gin.Context, body *models.DealerSignOutRequest) error {
	user := &cognito.GlobalSignOutInput{
		AccessToken: body.Token,
	}
	cognitoClient := cognito.New(api.awssession)
	_, err := cognitoClient.GlobalSignOut(user)
	if err != nil {
		return err
	}
	return nil
}
