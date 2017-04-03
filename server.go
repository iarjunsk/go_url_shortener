package main

import(
	"html/template"
	"net/http"
	"os"
	"go_shortify_web_app_heroku/models"
	"go_shortify_web_app_heroku/controllers"
)

type pageData struct {
	Title string
	Short_url string
	Long_url string
}

var tpl *template.Template
var page_data pageData
var host_name string = "https://app_id.herokuapp.com"
var notify_type int
var notify_msg string


func init() {
	tpl = template.Must(template.ParseGlob("views/*.html"))
	models.Redis_db_init()
}

func main(){
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/about",AboutHandler)
	http.HandleFunc("/404",ErrorHandler)
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	page_data = pageData{Title:"404"}
	tpl.ExecuteTemplate(w, "error.html",page_data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page_data = pageData{Title:"About"}
	tpl.ExecuteTemplate(w, "about.html",page_data)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {


	shortCode := r.URL.Path[1:]

	// The variables are static. So we need to re/initialize it every time
	page_data = pageData{Title:"Shortify",Short_url:"",Long_url:""}
	notify_type = 0

	if len(shortCode) != 0 { // GET from DB

		redirect_url,err := models.Redis_db_get(shortCode)
		if err != nil {
			redirect_url = host_name + "/404"
		}

		controllers.RedirectTo(w,r,redirect_url) // redirect to long url

		return

	}else if r.Method == "POST" { // SAVE to DB

		long_url := r.PostFormValue("long_url") //get form data by id

		err := controllers.ValidateURL(long_url)// validate url

		if err != nil {
			notify_type,notify_msg  = 4, "Invalid URL."
		}else{
			short_url := host_name + "/" + models.Redis_db_save(long_url)
			page_data = pageData{Title:"Shortify", Short_url:short_url ,Long_url:long_url}

			notify_type,notify_msg  = 1, "URL shortified."
		}
	}



	tpl.ExecuteTemplate(w, "app.html",page_data)
	controllers.ShowNotifications(w,notify_type,notify_msg) // run this after loading the page

}


