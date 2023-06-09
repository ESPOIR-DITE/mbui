package controller

import (
	"github.com/go-chi/chi"
	"html/template"
	"mbui/config"
	"net/http"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	//mux.Use(middleware.RequestID)
	//mux.Use(middleware.RealIP)
	//mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	//mux.Handle("/", homeHandler(env))
	mux.Handle("/", homeHandler(env))

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux

}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "index.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
