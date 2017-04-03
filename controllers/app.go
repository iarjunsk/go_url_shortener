package controllers

import (
	"hash/fnv"
	"net/http"
	"io"
	"fmt"
	"net/url"
)

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}


func RedirectTo(w http.ResponseWriter, r *http.Request, urlStr string){
	http.Redirect(w, r, urlStr, http.StatusFound)
}

func ValidateURL(long_url string) (error){
	_,err := url.ParseRequestURI(long_url)
	return err
}

func ShowNotifications(w io.Writer,notify_type int,msg string)  {
	var notify_script string = ""
	switch notify_type {
	case 1:
		notify_script = "<script>showSuccessMessage(\""+msg+"\")</script>"
	case 2:
		notify_script = "<script>showInfoMessage(\""+msg+"\")</script>"
	case 3:
		notify_script = "<script>showWarningMessage(\""+msg+"\")</script>"
	case 4:
		notify_script = "<script>showErrorMessage(\""+msg+"\")</script>"
	default:
		return
	}

	fmt.Fprint(w,notify_script)
}


