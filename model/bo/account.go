package bo

type AccountBo struct {
	Id       int64    `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Deleted  int64    `json:"deleted"`
}
