package convoso

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

// Lead this is a convoso lead
type Lead struct {
	ID                   string `json:"id"`
	CreatedAt            string `json:"created_at"`
	ModifiedAt           string `json:"modified_at"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Email                string `json:"email"`
	Status               string `json:"status"`
	SourceID             string `json:"source_id"`
	ListID               string `json:"list_id"`
	LocaleGmt            string `json:"locale_gmt"`
	CalledSinceLastReset string `json:"called_since_last_reset"`
	PhoneCode            string `json:"phone_code"`
	PhoneNumber          string `json:"phone_number"`
	AltPhone2            string `json:"alt_phone_2"`
	AltPhone1            string `json:"alt_phone_1"`
	CalledCount          string `json:"called_count"`
	CalledCountInbound   string `json:"called_count_inbound"`
	LastCalled           string `json:"last_called"`
	ObjectType           string `json:"object_type"`
	CreatedBy            string `json:"created_by"`
	CarrierName          string `json:"carrier_name"`
	LeadID               string `json:"lead_id"`
	Field1               string `json:"field_1"`
	Field2               string `json:"field_2"`
	Field3               string `json:"field_3"`
	Field11              string `json:"field_11"`
	Field13              string `json:"field_13"`
	Field84              string `json:"field_84"`
	UserUID              string `json:"user_uid"`
	UserName             string `json:"user_name"`
	DirectoryUID         string `json:"directory_uid"`
	DirectoryName        string `json:"directory_name"`
	SourceUID            string `json:"source_uid"`
	SourceName           string `json:"source_name"`
	CampaignUID          string `json:"campaign_uid"`
	CampaignName         string `json:"campaign_name"`
	StatusName           string `json:"status_name"`
}

// GetLead pass in a convoso LeadID and we can return all the details for a
// specific lead
func GetLead(leadID string) (Lead, error) {
	if leadID == "" {
		return Lead{}, errors.New("Lead ID can't be empty")
	}

	data := url.Values{}
	data.Add("lead_id", leadID)
	r, err := postFormRequest("https://api.convoso.com/v1/leads/search", data)
	if err != nil {
		return Lead{}, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Lead{}, err
	}

	result := LeadResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Errorf("RAW JSON: " + string(body))
		return Lead{}, err
	}

	if len(result.Data.Entries) != 1 {
		return Lead{}, errors.New("Only 1 result should be returned by the Lead Lookup route")
	}

	return result.Data.Entries[0], nil
}

// CreateLeadAndAddToList send the information to convoso to create a new lead
func CreateLeadAndAddToList(listName string, formData map[string]string) (Lead, error) {
	leadID, ok := listMapper[listName]
	if !ok {
		return Lead{}, errors.New("Please initialize this package with the correct mapping of lists and their IDs. The list name you provided is " + listName)
	}

	data := url.Values{}
	data.Add("list_id", strconv.Itoa(leadID))
	for key, val := range formData {
		data.Add(key, val)
	}

	// check for the required fields
	if data.Get("list_id") == "" || data.Get("phone_number") == "" {
		return Lead{}, errors.New("list_id and phone_number are both required fields to create a lead")
	}
	r, err := postFormRequest("https://api.convoso.com/v1/leads/insert", data)
	if err != nil {
		return Lead{}, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Lead{}, err
	}

	log.Info(string(body))

	response := CreateLeadResponse{}
	err = json.Unmarshal(body, &response)

	lead, err := GetLead(strconv.Itoa(response.Data.LeadID))
	if err != nil {
		return lead, err
	}

	return lead, nil
}

// UpdateLastCallDate set the field "date_notes_last_updated" for this convoso lead
func (l *Lead) UpdateLastCallDate() error {
	data := url.Values{}
	data.Add("lead_id", l.LeadID)
	data.Add("date_notes_last_updated", time.Now().Format("2006-01-02"))

	_, err := postFormRequest("https://api.convoso.com/v1/leads/update", data)
	if err != nil {
		return err
	}
	return nil
}
