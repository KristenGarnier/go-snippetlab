package main

import (
	"fmt"
	"net/http"
	"strconv"

	"go-snippetlab/pkg/forms"
	"go-snippetlab/pkg/models"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.Database.GetLatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "home.page.html", &HTMLData{
		Snippets: snippets,
	})
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Database.GetSnippet(id)
	if err != nil {
		app.NotFound(w)
		return
	}

	session := app.Sessions.Load(r)

	flash, err := session.PopString(w, "flash")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Flash:   flash,
		Snippet: snippet,
	})
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "new.page.html", nil)
}

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.NewSnippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: r.PostForm.Get("expires"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "new.page.html", &HTMLData{Form: form})
		return
	}

	id, err := app.Database.InsertSnippet(form.Title, form.Content, form.Expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	session := app.Sessions.Load(r)

	err = session.PutString(w, "flash", "Your snippet wad saved successfully!")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}

func (app *App) SignupUser(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "signup.page.html", &HTMLData{
		Form: &forms.SignupUser{},
	})
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.SignupUser{
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "signup.page.html", &HTMLData{
			Form: form,
		})

		return
	}

	err = app.Database.InsertUser(form.Name, form.Email, form.Password)
	if err == models.ErrDuplicateEmail {
		form.Failures["Email"] = "Address is already in use"
		app.RenderHTML(w, r, "signup.page.html", &HTMLData{
			Form: form,
		})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	msg := "Your signup was successful. Please log in using your credentails"
	session := app.Sessions.Load(r)
	err = session.PutString(w, "flash", msg)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	session := app.Sessions.Load(r)

	flash, err := session.PopString(w, "flash")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "login.page.html", &HTMLData{
		Flash: flash,
		Form:  &forms.LoginUser{},
	})
}

func (app *App) VerifyUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.LoginUser{
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "login.page.html", &HTMLData{
			Form: form,
		})
		return
	}

	currentUserId, err := app.Database.VerifyUser(form.Email, form.Password)
	if err == models.ErrInvalidCredentials {
		form.Failures["Generic"] = "Email or Password is incorrect"
		app.RenderHTML(w, r, "login.page.html", &HTMLData{
			Form: form,
		})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	session := app.Sessions.Load(r)
	err = session.PutInt(w, "currentUserID", currentUserId)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/snippet/new", http.StatusSeeOther)
}

func (app *App) LogoutUser(w http.ResponseWriter, r *http.Request) {
	session := app.Sessions.Load(r)
	err := session.Remove(w, "currentUserId")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/", 303)
}
