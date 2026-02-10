package vo

type JwtVo struct {
	TokenType       string `json:"tokenType"`
	AccessToken     string `json:"accessToken"`
	TelegramWarning bool   `json:"telegram_warning,omitempty"`
}
