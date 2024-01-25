package post;




var Post struct {
	ID          models.PrimitiveObjectID   `json:_id,omitempty`
	Title       string                     `json:_id,omitempty`
	Likes       int64                      `json:likes,omitempty`
	dislikes    int64                      `json:dislikes,omitempty`
}