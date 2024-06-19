package qrcode

import (
	"github.com/skip2/go-qrcode"
)

func Generate(content string, size int) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, size)
}
