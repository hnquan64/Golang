package models

import (
	"fmt"

	"github.com/hanzoai/gochimp3"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateTemplateFolder(client *gochimp3.API, name string) string {
	request := &gochimp3.TemplateFolderCreationRequest{
		Name: name,
	}

	folder, err := client.CreateTemplateFolder(request)
	CheckError(err)
	return folder.ID
}

func CreateTemplate(client *gochimp3.API, name string, htmlString string, folder_id string) uint {
	request := &gochimp3.TemplateCreationRequest{
		Name:     name,
		Html:     htmlString,
		FolderId: folder_id,
	}
	templ, err := client.CreateTemplate(request)
	CheckError(err)
	return templ.ID
}

func CreateCampaignFolder(client *gochimp3.API, name string) string {
	request := &gochimp3.CampaignFolderCreationRequest{
		Name: name,
	}
	campaign_folder, err := client.CreateCampaignFolder(request)
	CheckError(err)
	return campaign_folder.ID
}

func CreateCampaign(client *gochimp3.API, listID string, campaign_folder string, template_id uint) string {

	type condition struct {
		Field string
		Op    string
		Value interface{}
	}
	conditions := make([]condition, 0)
	cond := condition{"Email Addres", "contains", "@gmail"}
	conditions = append(conditions, cond)
	segments := &gochimp3.CampaignCreationSegmentOptions{
		SavedSegmentId: 12141618,
		Match:          gochimp3.CONDITION_MATCH_ANY,
		Conditions:     conditions,
	}
	recipients := &gochimp3.CampaignCreationRecipients{
		ListId:         listID,
		SegmentOptions: *segments,
	}

	settings := &gochimp3.CampaignCreationSettings{
		SubjectLine: "Chien dich 01",
		PreviewText: "click me!",
		Title:       "Khoi dong chien dich 1",
		FromName:    "Phan Chuan",
		ReplyTo:     "trondoibenem19xx@gmail.com",
		FolderId:    campaign_folder,
		TemplateId:  template_id,
	}
	request := &gochimp3.CampaignCreationRequest{
		Type:       gochimp3.CAMPAIGN_TYPE_REGULAR,
		Recipients: *recipients,
		Settings:   *settings,
	}

	campaign, err := client.CreateCampaign(request)
	CheckError(err)
	return campaign.ID
}

func SendCampaign(client *gochimp3.API, campaign_id string) {
	request := gochimp3.SendCampaignRequest{
		CampaignId: campaign_id,
	}
	_, err := client.SendCampaign(campaign_id, &request)
	CheckError(err)
}
