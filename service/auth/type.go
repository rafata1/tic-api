package auth

import (
	"encoding/json"
	"io"
)

type UserInfo struct {
	Email string `json:"email"`
}

func unmarshalUserInfo(buf io.Reader) (UserInfo, error) {
	var info UserInfo
	err := json.NewDecoder(buf).Decode(&info)
	return info, err
}
