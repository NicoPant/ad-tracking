package tracker

type Tracker struct {
	Id    string `json:"_key,omitempty" bson:"_key,omitempty"`
	AdId  string `json:"ad_id,omitempty" bson:"ad_id" required:"true"`
	Count int    `json:"count" bson:"count" required:"true"`
}
