package model

type MeetingLayout struct {
	LogoUrl      string   `json:"logo_url,omitempty" bson:"logo_url,omitempty"`
	LogoSmallUrl string   `json:"logo_small_url,omitempty" bson:"logo_small_url,omitempty"`
	BannerUrl    string   `json:"banner_url,omitempty" bson:"banner_url,omitempty"`
	ColorSet     ColorSet `json:"color_set,omitempty" bson:"color_set,omitempty"`
}
