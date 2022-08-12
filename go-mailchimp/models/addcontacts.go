package models

import (
	"gomailchimp/getmailform"
	"os"
	"time"

	"github.com/hanzoai/gochimp3"
)

func AddContacts() {
	// CREATED BY MAILCHIMP ACCOUNT
	// https://mailchimp.com/help/about-api-keys/
	api_key := os.Getenv("API_KEY1")

	// https://mailchimp.com/help/find-audience-id/
	audience_id := os.Getenv("AUDIENCE_ID")
	client := gochimp3.New(api_key)
	client.Timeout = (5 * time.Second)

	emails := getmailform.GetMails()

	// Fetch list
	list, err := client.GetList(audience_id, nil)
	getmailform.CheckError(err)

	// Add subscriber in mailchimp
	for i := len(emails) - 1; i >= 0; i-- {
		req := &gochimp3.MemberRequest{
			EmailAddress: emails[i],
			Status:       "subscribed",
		}
		if _, err := list.CreateMember(req); err != nil {
			break
		}
	}
}
