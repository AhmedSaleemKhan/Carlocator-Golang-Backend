package api

import (
	"fmt"

	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) StaffResetPasswordAPI(c *gin.Context, staff *models.Staff) error {
	*staff.StaffId = c.MustGet("user-id").(string)
	staffDB, err := api.db.GetStaffByIDDB(c, staff)
	if err != nil {
		return err
	}
	if staffDB == nil {
		return fmt.Errorf("staff doesn't exist")
	}
	cognitoClient := cognito.New(api.awssession)
	user := &cognito.AdminSetUserPasswordInput{
		Username:   aws.String(*staffDB.EmailWork),
		Password:   aws.String(*staff.Password),
		Permanent:  aws.Bool(true),
		UserPoolId: aws.String(config.Cfg.CognitoPoolId),
	}
	_, err = cognitoClient.AdminSetUserPassword(user)
	if err != nil {
		return nil
	}
	return nil
}
