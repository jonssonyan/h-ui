package bo

type AccountBo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Deleted  int64  `json:"deleted"`
}
