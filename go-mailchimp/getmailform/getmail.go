package getmailform

import (
	"context"
	"fmt"

	"google.golang.org/api/forms/v1"
	"google.golang.org/api/option"
)

const forms_id = "1HxzCHDlHbAwWE3WYFjbrVz2ZiDANSKnyrs_1Bplk870"

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetMails() []string {

	// CREDENTIALS GOOGLE CLOUD WITH SERVICE ACCOUNT (file .json)
	ctx := context.Background()
	service, err := forms.NewService(ctx, option.WithCredentialsFile("credentials.json"))
	CheckError(err)

	// GET RESPONSE FORMS WITH FORMS_ID
	response, err := service.Forms.Responses.List(forms_id).Do()
	CheckError(err)

	var response_id []string // slice response id get from forms
	var emails []string      // slice emails get from forms

	// GET LIST RESPONSE ID FROM FORMS
	for _, resp := range response.Responses {
		response_id = append(response_id, resp.ResponseId)
	}

	// GET EMAIL WITH FORM_ID & RESPONSE_ID
	for _, each := range response_id {
		result, err := service.Forms.Responses.Get(forms_id, each).Do()
		CheckError(err)
		for _, answer := range result.Answers {
			for _, email := range answer.TextAnswers.Answers {
				emails = append(emails, email.Value)
			}
		}

	}
	return emails
}
