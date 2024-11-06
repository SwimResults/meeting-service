package model

type Sponsor struct {
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	ImageUrl   string `json:"image_url,omitempty" bson:"image_url,omitempty"`
	WebsiteUrl string `json:"website_url,omitempty" bson:"website_url,omitempty"`
}
