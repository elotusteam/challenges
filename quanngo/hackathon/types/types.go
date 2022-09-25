package types

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type Register struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type FileInfo struct {
	FileName string `json:"file_name"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}
