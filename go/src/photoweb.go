/**
 * Created with IntelliJ IDEA.
 * User: sunlzx
 * Date: 13-7-20
 * Time: 上午8:48
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	UPLOAD_DIR   = "uploads"
	TEMPLATE_DIR = "views"
)

var templates map[string]*template.Template

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.Method == "GET" {
			//			io.WriteString(w, "<html><body><form method=\"POST\" action=\"/upload\" " +
			//						" enctype=\"multipart/form-data\">" +
			//						"Choose an image to upload: <input name=\"image\" type=\"file\" />" +
			//						"<input type=\"submit\" value=\"Upload\" />" +
			//						"</form></body></html>")

			// t, err := template.ParseFiles("upload.html")
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }
			// t.Execute(w, nil)
			// return

			if err := renderHtml(w, "upload.html", nil); err != nil {
				http.Error(w, err.Error(),
					http.StatusInternalServerError)
				return
			}
		}

	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename,
			http.StatusFound)
	}

}

func viewHandler_old(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

//判断文件是否存在
//判断文件是否存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	//	var listHtml string
	//	for _, fileInfo := range fileInfoArr {
	//		imgid := fileInfo.Name()
	//		listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
	//	}
	//
	//	io.WriteString(w, "<html><body><ol>" + listHtml + "</ol></body></html>")

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	// t, err := template.ParseFiles("list.html")
	// if err != nil {
	// 	http.Error(w, err.Error(),
	// 		http.StatusInternalServerError)
	// 	return
	// }
	// t.Execute(w, locals)

	if err := renderHtml(w, "list.html", locals); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	// t, err := template.ParseFiles(tmpl + ".html")
	// if err != nil {
	// 	return err
	// }
	// err = t.Execute(w, locals)

	err := templates[tmpl].Execute(w, locals)
	return err
}

func init2() {
	templates = make(map[string]*template.Template)
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()

		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))

		templates[templateName] = t
	}
}

func main_bak() {
	fmt.Println("xx")
	init2()
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/", listHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
