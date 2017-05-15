package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/objx"
	"github.com/stretchr/signature"
)

var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	var callback = os.Getenv("CALLBACK_URL")
	if callback == "" {
		callback = "http://localhost:8080/auth/callback"
	}
	// Parse the flags
	flag.Parse()

	// Auth providers
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New(os.Getenv("FACEBOOK_ID"), os.Getenv("FACEBOOK_KEY"), fmt.Sprintf("%s/facebook", callback)),
		github.New(os.Getenv("GITHUB_ID"), os.Getenv("GITHUB_KEY"), fmt.Sprintf("%s/github", callback)),
		// google.New("key", "secret", fmt.Sprintf("%s/google", callback)),
	)

	// Initialize the room
	r := newRoom()
	// r.tracer = trace.New(os.Stdout)

	// Serve pages
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})

	http.HandleFunc("/uploader", uploaderHandler)

	// Authenticate users
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	// Serve websocket upgrader
	http.Handle("/room", r)

	// Serve static assets
	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("./assets"))))
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

	go r.run()

	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
