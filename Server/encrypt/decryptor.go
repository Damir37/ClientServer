package encrypt

import (
	"encoding/base64"
	"log"
)

func DecryptMessage(msg string) string {
	decoding, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		log.Fatal(err)
	}
	return string(decoding)
}
