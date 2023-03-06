package encrypt

import "encoding/base64"

func EcnryptMessage(msg string) string {
	encoding := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoding
}
