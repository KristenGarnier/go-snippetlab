package main

import (
	"go-snippetlab/pkg/models"

	"github.com/alexedwards/scs"
)

type App struct {
	Addr      string
	Database  *models.Database
	HTMLDir   string
	Sessions  *scs.Manager
	StaticDir string
	TLSCert   string
	TLSKey    string
}
