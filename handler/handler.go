package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"progate/entity"
	"strconv"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl,err:=template.ParseFiles(path.Join("views","index.html"), path.Join("views","layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w,"Error is happening ,keep calm",http.StatusInternalServerError)
		return
	}

	//data:= map[string]interface{}{
		//"title":"I'm learning Golang web",
		//"content":"I;m learning web nathan ",
	//}

	//data :=entity.Product{ID: 1, Name:"Mobilio", Price:2200000, Stock: 3}
	data:=[]entity.Product{
		{ID: 1,Name:"Mobilio",Price:2200000, Stock:15},
		{ID: 2,Name:"Nathan",Price:300000, Stock:8},
		{ID: 3,Name:"Pajer",Price:1300000, Stock:1},
	}

	err=tmpl.Execute(w, data)
	if err != nil{
		log.Println(err)
		http.Error(w,"Error is happening ,keep calm",http.StatusInternalServerError)
		return
	}
}

func Hellohandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world,saya belajar golang"))
}

func Mariohandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from nitendo"))
}

func Producthandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Product page :%d", idNumb)
	data:=map[string]interface{}{
		"content":idNumb,
	}

	tmpl,err :=	template.ParseFiles(path.Join("views","product.html"), path.Join("views","layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w,"Error is happening ,keep calm",http.StatusInternalServerError)
		return
	}

	err= tmpl.Execute(w, data)
	if err != nil{
		log.Println(err)
		http.Error(w,"Error is happening ,keep calm",http.StatusInternalServerError)
		return
	}

}
func PostGet(w http.ResponseWriter, r *http.Request){
	method := r.Method

	switch method{
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "error is happening,keep calm",http.StatusBadRequest)
	}
} 

func Form(w http.ResponseWriter, r *http.Request){
	if r.Method =="GET"{
		tmpl, err :=template.ParseFiles(path.Join("views","form.html"),path.Join("views","layout.html"))
		if err !=nil{
			log.Println(err)
			http.Error(w,"Error is happening,keep calm",http.StatusInternalServerError)
			return 
		}

		err=tmpl.Execute(w,nil)
		if err != nil{
			log.Println(err)
			http.Error(w,"Error is happening,keep calm",http.StatusInternalServerError)
			return
		}

		return
	}
	
	http.Error(w, "error is happening,keep calm",http.StatusBadRequest)
}

func Process(w http.ResponseWriter,r *http.Request){
	if r.Method =="POST"{
		err:=r.ParseForm()
		if err != nil{
			log.Println(err)
			http.Error(w,"Error is happening,keep calm",http.StatusInternalServerError)
			return
		}

		name :=r.Form.Get("name")
		message:=r.Form.Get("message")

		data:=map[string]interface{}{
			"name":name,
			"message":message,
		}

		tmpl,err :=template.ParseFiles(path.Join("views","result.html"),path.Join("views","layout.html"))
		if err !=nil{
			log.Println(err)
			http.Error(w,"Error is happening,keep calm",http.StatusInternalServerError)
			return 
		}

		err=tmpl.Execute(w,data)
		if err != nil{
			log.Println(err)
			http.Error(w,"Error is happening,keep calm",http.StatusInternalServerError)
			return
		}
		return

		
	}
		
	http.Error(w, "error is happening,keep calm",http.StatusBadRequest)
}