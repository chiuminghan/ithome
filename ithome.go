package ithome

import (
	"github.com/chiuminghan/ithome/ironman"

	log "github.com/sirupsen/logrus"
)

func PrintItHome() {
	log.Info("hi chiuminghan")
	ironman.PrintIronMan()
}
