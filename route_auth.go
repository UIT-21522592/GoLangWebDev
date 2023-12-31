package main
import (
	"net/http"
)

func login(writer http.ResponseWriter, request *http.Request){
	t := parseTemplateFiles("login.layout","public.navbar","login")
	t.Execute(writer, nil)
}

func signup(writer http.ResponseWriter, request *http.Request){
	generateHTML(writer, nil, "login.layout","public.navbar","signup")
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
	session := user.CreateSession()
	cookie := http.Cookie{
	Name: "_cookie",
	Value: session.Uuid,
	HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
	} else {
	http.Redirect(w, r, "/login", 302)
	}
}
