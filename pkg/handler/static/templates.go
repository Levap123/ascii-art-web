package static

import (
	"html/template"
	"net/http"
)

const templatesDir = "templates/"

// data any: информация, которую ты закидываешь в html темплейт. Указать nil при вызове, если таковой
func ExecuteTemplate(templateName string, data any, w http.ResponseWriter) error {
	templatePath := templatesDir + templateName
	templ := template.Must(template.ParseFiles(templatePath))
	if err := templ.Execute(w, data); err != nil {
		return err
	}
	return nil
}
