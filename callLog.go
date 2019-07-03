package convoso

// CallLog This struct represents a single log entry from Convoso
type CallLog struct {
	ID           *string `json:"id"`
	LeadID       *string `json:"lead_id"`
	ListID       *string `json:"list_id"`
	CampaignID   *string `json:"campaign_id"`
	Campaign     *string `json:"campaign"`
	User         *string `json:"user"`
	UserID       *string `json:"user_id"`
	PhoneNumber  *string `json:"phone_number"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Status       *string `json:"status"`
	StatusName   *string `json:"status_name"`
	CallLength   *string `json:"call_length"`
	CallDate     *string `json:"call_date"`
	AgentComment *string `json:"agent_comment"`
	QueueID      *string `json:"queue_id"`
}
