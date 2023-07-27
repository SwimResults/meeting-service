package model

type MeetingData struct {
	LivetimingUrl     string `json:"livetiming_url,omitempty" bson:"livetiming_url,omitempty"`
	WebsiteUrl        string `json:"website_url,omitempty" bson:"website_url,omitempty"`
	StreamUrl         string `json:"stream_url,omitempty" bson:"stream_url,omitempty"`
	InstagramUrl      string `json:"instagram_url,omitempty" bson:"instagram_url,omitempty"`
	FacebookUrl       string `json:"facebook_url,omitempty" bson:"facebook_url,omitempty"`
	HasLivetiming     bool   `json:"has_livetiming,omitempty" bson:"has_livetiming,omitempty"`
	HasFtpUpload      bool   `json:"has_ftp_upload,omitempty" bson:"has_ftp_upload,omitempty"`
	FtpStartListMask  string `json:"ftp_start_list_mask,omitempty" bson:"ftp_start_list_mask,omitempty"`
	FtpResultListMask string `json:"ftp_result_list_mask,omitempty" bson:"ftp_result_list_mask,omitempty"`
}
