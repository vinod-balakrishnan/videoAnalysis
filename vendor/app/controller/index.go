package controller

import (
	"net/http"

	"app/shared/session"
	"app/shared/view"

	"github.com/josephspurrier/csrfbanana"
)

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	// if session.Values["id"] != nil {
	// Display the view
	v := view.New(r)
	v.Name = "index/auth"
	v.Vars["first_name"] = session.Values["first_name"]
	v.Render(w)
	// } else {
	// 	// Display the view
	// 	v := view.New(r)
	// 	v.Name = "index/anon"
	// 	v.Render(w)
	// 	return
	// }
}

// IndexAnalyse displays the home page
func IndexAnalyse(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	v := view.New(r)
	v.Name = "analysis/analysis"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)

	v.Render(w)

}
