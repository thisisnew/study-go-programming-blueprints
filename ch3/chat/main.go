package main

import (
	"flag"
	"fmt"
	"github.com/matryer/goblueprints/chapter1/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("C:/Users/dhkim/GolandProjects/study-go-programming-blueprints/ch3/chat/templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	fmt.Println("data[\"UserData\"] :", data["UserData"])

	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	gomniauth.SetSecurityKey("PUT YOUR AUTH KEY HERE")
	gomniauth.WithProviders(
		facebook.New("key", "secret", "http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret", "http://localhost:8080/auth/callback/github"),
		//google.New("key", "secret", "http://localhost:8080/auth/callback/google"),
		google.New("499113554757-9mevutrp9htdoinok3n60sik5gnr46qd.apps.googleusercontent.com", "jKNsrk8gIPhV_wn-XRd3z78z", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		http.SetCookie(writer, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		writer.Header().Set("Location", "/chat")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	})

	go r.run()

	log.Println("Starting web server on", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
