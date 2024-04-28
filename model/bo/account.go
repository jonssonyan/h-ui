package bo

type AccountBo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	IsAdmin  int64  `json:"isAdmin"`
	Deleted  int64  `json:"deleted"`
}
