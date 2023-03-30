package attachedto

import "time"

type AttacheResponse struct {
	Title    string    `json:"name" form:"name"`
	Image    string    `json:"image" form:"image" `
	EndDate  time.Time `json:"endDate"`
	Donation int       `json:"donation" form:"donation" `
	Desc     string    `json:"desc" form:"desc" `
}
