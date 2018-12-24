package obsws

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[obsws] ", log.LstdFlags)
