package convoso

// LeadResponse this is a response struct from convoso when requesting Leads via
// their API
type LeadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Offset  interface{} `json:"offset"`
		Limit   interface{} `json:"limit"`
		Total   string      `json:"total"`
		Entries []Lead      `json:"entries"`
	} `json:"data"`
}
