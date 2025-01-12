package model

type Profile struct {
	UserID string  `json:"user_id"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
	Goal   string  `json:"goal"`
}
