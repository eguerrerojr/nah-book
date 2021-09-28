package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/eguerrerojr/nah-book/internal/models"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProd        bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
