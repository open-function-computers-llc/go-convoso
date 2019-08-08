package convoso

// CreateLeadResponse this is the response returned from convoso when creating a
// lead via their API
type CreateLeadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		LeadID int `json:"lead_id"`
	} `json:"data"`
}
