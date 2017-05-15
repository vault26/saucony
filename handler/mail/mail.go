package mail

import (
	"os"
	"strings"
)

var (
	DOMAIN         = os.Getenv("MG_DOMAIN")
	API_KEY        = os.Getenv("MG_API_KEY")
	PUBLIC_API_KEY = os.Getenv("MG_PUBLIC_API_KEY")
	ADMIN_EMAILS   = strings.Split(os.Getenv("ADMIN_EMAILS"), ",")
)
