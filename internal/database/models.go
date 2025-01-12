package database

type Profile struct {
	ID     string  `db:"id"`
	UserID string  `db:"user_id"`
	Gender string  `db:"gender"`
	Age    int     `db:"age"`
	Weight float32 `db:"weight"`
	Height float32 `db:"height"`
	Goal   string  `db:"goal"`
}
