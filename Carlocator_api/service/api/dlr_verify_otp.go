package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/lib"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DealerVerifyOTPAPI(c *gin.Context, body *models.OTPRequest) (*models.VerifyOTPResponse, error) {
	user := &cognito.ConfirmSignUpInput{
		ConfirmationCode: aws.String(*body.OTP),
		Username:         aws.String(*body.Email),
		ClientId:         aws.String(config.Cfg.AppClientId),
	}
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *body.Email, config.Cfg.AppClientId)
	user.SecretHash = aws.String(secretHash)
	CognitoClient := cognito.New(api.awssession)
	_, err := CognitoClient.ConfirmSignUp(user)
	if err != nil {
		return nil, err
	}
	groupName := lib.SalesManager
	groupUser := &cognito.AdminAddUserToGroupInput{
		GroupName:  &groupName,
		UserPoolId: &config.Cfg.CognitoPoolId,
		Username:   body.Email,
	}
	_, err = CognitoClient.AdminAddUserToGroup(groupUser)
	if err != nil {
		return nil, err
	}
	dealer := models.NewDealer(body.Email, nil, nil)
	*dealer.EmailVerified = true
	resp, err := api.db.DealerVerifyOTPDB(c, dealer)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
