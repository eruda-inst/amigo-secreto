package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"xmas-list/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsPresentes := models.BuscaTodosOsPresentes()
	temp.ExecuteTemplate(w, "Index", todosOsPresentes)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		link := r.FormValue("link")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		models.CriarNovoPresente(nome, descricao, precoConvertidoParaFloat, link)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoPresente := r.URL.Query().Get("id")
	models.DeletePresente(idDoPresente)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoPresente := r.URL.Query().Get("id")
	presente := models.EditaPresente(idDoPresente)
	temp.ExecuteTemplate(w, "Edit", presente)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		link := r.FormValue("link")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		models.AtualizaPresente(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, link)
	}
	http.Redirect(w, r, "/", 301)
}
