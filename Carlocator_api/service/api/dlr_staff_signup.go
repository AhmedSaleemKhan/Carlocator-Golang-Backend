package api

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/lib"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

const (

	// EmailSender Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	EmailSender = "info@findr.click"
)

func (api *CarLocatorApiImpl) StaffSignupAPI(c *gin.Context, newStaff *models.Staff) (*models.Staff, error) {
	dealerRole := c.MustGet("cognito-group").(string)
	if dealerRole != lib.SalesManager {
		return nil, fmt.Errorf("you are not authorize to add staff")
	}
	existedStaff, err := api.db.GetStaffDB(c, newStaff)
	if err != nil {
		return nil, err
	}
	if existedStaff != nil {
		err = CheckRecordExists(newStaff.EmailWork, newStaff.Username, existedStaff.EmailWork, existedStaff.Username)
		if err != nil {
			return nil, err
		}
	}
	CognitoClient := cognito.New(api.awssession)
	resp, err := CognitoSignup(CognitoClient, newStaff)
	if err != nil {
		return nil, err
	}
	maxStaffNumber, err := api.db.GetLatestStaffNumber(c)
	if err != nil {
		return nil, err
	}
	dealerName, err := api.db.GetStaffDealerName(c)
	if err != nil {
		return nil, err
	}
	maxRoleKeys, err := api.db.GetMaximumRoleKeysDB(c, *newStaff.StaffRoleName)
	if err != nil {
		return nil, err
	}
	*newStaff.StaffId = *resp.UserSub
	*newStaff.MaximumRoleKeys = *maxRoleKeys
	*newStaff.DealerID = c.MustGet("user-id").(string)
	*newStaff.CreatedAt = api.clock.NowUnix()
	literals := "SN"
	var staffNumber string
	if dealerName != "" {
		literals = string(dealerName[0])
		for i := 0; i < len(dealerName); i++ {
			if dealerName[i] == ' ' {
				literals += string(dealerName[i+1])
			}
		}
		staffNumber = fmt.Sprintf("%v-%v", literals, *maxStaffNumber)
	} else {
		staffNumber = fmt.Sprintf("%v-%v", literals, *maxStaffNumber)
	}
	*newStaff.StaffNumber = staffNumber
	signupResp, err := api.db.StaffSignupDB(c, newStaff)
	if err != nil {
		return nil, err
	}
	err = sendUserNameAndPassword(*newStaff.EmailWork, *newStaff.Username, *newStaff.Password)
	if err != nil {
		log.Println(err)
	}
	return signupResp, nil
}
func sendUserNameAndPassword(emailWork, username, password string) error {
	// Subject The subject line for the email.
	Subject := "You are  Successfully registered with LotVector👏"
	// HtmlBody The HTML body for the email.

	HtmlBody := "<h1>Lot Vector Registration Credentials></h1><br/>" +
		"<h4>Your Email : " + emailWork + "</h4><br/>" +
		"<h4>Your Password : " + password + "</h4>"
	// TextBody The email body for recipients with non-HTML email clients.

	TextBody := fmt.Sprintf("Your account is created with following\nEmail: %v\nUsername: %v\n,Password: %v", emailWork, username, password)
	// CharSet The character encoding for the email.

	CharSet := "UTF-8"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Cfg.Region)},
	)
	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(emailWork),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(EmailSender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}
	// Attempt to send the email.
	_, err = svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if err, ok := err.(awserr.Error); ok {
			switch err.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, err.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, err.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, err.Error())
			default:
				return err
			}
		} else {
			// Print the error, cast err to aws err.Error to get the Code and
			// Message from an error.
			return err
		}
		return err
	}
	return nil
}

func CognitoSignup(CognitoClient *cognito.CognitoIdentityProvider, newStaff *models.Staff) (*cognito.SignUpOutput, error) {
	staff := &cognito.SignUpInput{
		Username: aws.String(*newStaff.EmailWork),
		Password: aws.String(*newStaff.Password),
		ClientId: aws.String(config.Cfg.AppClientId),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(*newStaff.EmailWork),
			},
		},
	}
	secretHash := ComputeSecretHash(config.Cfg.AppClientSecret, *newStaff.EmailWork, config.Cfg.AppClientId)
	staff.SecretHash = aws.String(secretHash)
	resp, err := CognitoClient.SignUp(staff)
	if err != nil {
		return nil, err
	}
	confirmSignup := &cognito.AdminConfirmSignUpInput{
		Username:   aws.String(*newStaff.EmailWork),
		UserPoolId: aws.String(config.Cfg.CognitoPoolId),
	}
	if _, err := CognitoClient.AdminConfirmSignUp(confirmSignup); err != nil {
		return nil, err
	}
	groupUser := &cognito.AdminAddUserToGroupInput{
		GroupName:  newStaff.StaffRoleName,
		UserPoolId: &config.Cfg.CognitoPoolId,
		Username:   newStaff.EmailWork,
	}
	if _, err := CognitoClient.AdminAddUserToGroup(groupUser); err != nil {
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
