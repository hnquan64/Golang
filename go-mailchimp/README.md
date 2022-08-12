# Step 1: Get emails from Google Forms

e21 <-> ***
359 <-> ***
 * Needs:
 - File credentials.json is service account of google cloud
 - FormsID is id of google forms, get in url
 * Workflow:
 - Credentials google cloud, make sure to enable google forms api
 - Get list response_id with forms_id
 - Get emails from list response with response_id
 

 # Step 2: Add contacts on Mailchimp
 * Needs:
 - List email get from Step 1
 - API_KEY & AUDIENCE_ID from account Mailchimp
 * Workflow:
 - Get 