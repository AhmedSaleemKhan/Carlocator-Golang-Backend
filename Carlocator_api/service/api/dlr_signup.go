package api

import (
	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/lib"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DealerSignupAPI(c *gin.Context, dealer *models.Dealer) (*models.DealerResponse, error) {
	existedDealer, err := api.db.GetDealerDB(c, dealer)
	if err != nil {
		return nil, err
	}
	if existedDealer != nil {
		err = CheckRecordExists(dealer.Email, dealer.Username, existedDealer.Email, existedDealer.Username)
		if err != nil {
			return nil, err
		}
	}
	CognitoClient := cognito.New(api.awssession)
	user := &cognito.SignUpInput{
		Username: aws.String(*dealer.Email),
		Password: aws.String(*dealer.Password),
		ClientId: aws.String(config.Cfg.AppClientId),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(*dealer.Email),
			},
		},
	}
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *dealer.Email, config.Cfg.AppClientId)
	user.SecretHash = aws.String(secretHash)
	resp, err := CognitoClient.SignUp(user)
	if err != nil {
		return nil, err
	}
	*dealer.DealerId = *resp.UserSub
	*dealer.CreatedAt = api.clock.NowUnix()
	*dealer.EmailVerified = *resp.UserConfirmed
	*dealer.RoleName = lib.SalesManager
	signupResp, err := api.db.DealerSignupDB(c, dealer)
	if err != nil {
		return nil, err
	}

	return signupResp, nil
}

