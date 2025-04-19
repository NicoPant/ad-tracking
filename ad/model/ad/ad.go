package ad

type Ad struct {
	Id          string `json:"_key,omitempty" bson:"_key,omitempty"`
	Title       string `json:"title" bson:"title" required:"true"`
	Description string `json:"description" bson:"description" required:"true"`
	Url         string `json:"url" bson:"url" required:"true"`
}
