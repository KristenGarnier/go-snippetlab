package main

import "go-snippetlab/pkg/models"

type App struct {
	Database  *models.Database
	HTMLDir   string
	StaticDir string
}
