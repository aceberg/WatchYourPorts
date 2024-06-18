package web

import (
	"embed"

	"github.com/aceberg/WatchYourPorts/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	allAddrs map[string]models.AddrToScan

	// quitScan - send stop signal to routineScan
	quitScan chan bool
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
