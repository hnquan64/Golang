package campaigns

import (
	"io/ioutil"
	"time"

	"gomailchimp/models"

	"github.com/hanzoai/gochimp3"
)

const (
	//https://mailchimp.com/help/about-api-keys/
	API_KEY = "15ce0149d9cd8f2a649a6d08d483f972-us8"
	//https://mailchimp.com/help/find-audience-id/
	listID = "a636b93906"
)

func InitCampaign() {
	// CREATE CLIENT THEN CONNECT MAILCHIMP WITH API_KEY
	client := gochimp3.New(API_KEY)
	client.Timeout = (5 * time.Second)

	// CREATE TEMPLATE FOLDER THEN GET ID
	templateFolder_id := models.CreateTemplateFolder(client, "Be01")
	// INPUT OF CREATE TEMPLATE FUNCTION
	html, err := ioutil.ReadFile("template.html")
	models.CheckError(err)
	htmlString := string(html)
	// CREATE TEMPLATE
	template_id := models.CreateTemplate(client, "demo", htmlString, templateFolder_id)
	// CREATE CAMPAIGN FOLDER THEN GET ID
	campaignFolder_id := models.CreateCampaignFolder(client, "Be01")
	// CREATE CAMPAIGN THEN GET ID
	campaign_id := models.CreateCampaign(client, listID, campaignFolder_id, template_id)

	// PROCESS SEND CAMPAIGN
	models.SendCampaign(client, campaign_id)
}
