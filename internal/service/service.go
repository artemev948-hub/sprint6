package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) string {
	data = strings.TrimSpace(data)

	if strings.ContainsAny(data, ".-") {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)
}
