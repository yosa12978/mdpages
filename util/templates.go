package util

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/yosa12978/mdpages/session"
	"github.com/yosa12978/mdpages/types"
)

func RenderView(w io.Writer, r *http.Request, view string, payload any) error {
	templPath := fmt.Sprintf("templates/views/%s.html", view)
	templ, err := template.ParseFiles(
		templPath,
		"templates/top.html",
		"templates/bottom.html",
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	usr, _ := session.GetSession(r)
	data := types.Templ{
		User:    usr,
		Title:   "mdpages",
		Payload: payload,
	}
	return templ.Execute(w, data)
}

func RenderBlock(w io.Writer, name string, payload any) error {
	templ := template.Must(
		template.ParseFiles(
			"templates/blocks/articles.html",
			"templates/blocks/categories.html",
		),
	)
	return templ.ExecuteTemplate(w, name, payload)
}
