package models

import (
	"xmas-list/db"
)

type Presente struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Link      string
}

func BuscaTodosOsPresentes() []Presente {
	db := db.ConectarComBancoDados()

	selectDeTodosOsPresentes, err := db.Query("select * from presente order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Presente{}
	presentes := []Presente{}

	for selectDeTodosOsPresentes.Next() {
		var id int
		var nome, descricao, link string
		var preco float64

		err = selectDeTodosOsPresentes.Scan(&id, &nome, &descricao, &preco, &link)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Link = link

		presentes = append(presentes, p)
	}
	defer db.Close()
	return presentes

}

func CriarNovoPresente(nome string, descricao string, preco float64, link string) {
	db := db.ConectarComBancoDados()

	insereDadosNoBanco, err := db.Prepare("insert into presente(nome, descricao, preco, link) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, link)
	defer db.Close()
}

func DeletePresente(id string) {
	db := db.ConectarComBancoDados()

	deletarOPresente, err := db.Prepare("delete from presente where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOPresente.Exec(id)
	defer db.Close()

}

func EditaPresente(id string) Presente {
	db := db.ConectarComBancoDados()

	presenteDoBanco, err := db.Query("select * from presente where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	presenteParaAtualizar := Presente{}

	for presenteDoBanco.Next() {
		var id int
		var nome, descricao, link string
		var preco float64

		err = presenteDoBanco.Scan(&id, &nome, &descricao, &preco, &link)
		if err != nil {
			panic(err.Error())
		}
		presenteParaAtualizar.Id = id
		presenteParaAtualizar.Nome = nome
		presenteParaAtualizar.Descricao = descricao
		presenteParaAtualizar.Preco = preco
		presenteParaAtualizar.Link = link
	}
	defer db.Close()
	return presenteParaAtualizar
}

func AtualizaPresente(id int, nome string, descricao string, preco float64, link string) {
	db := db.ConectarComBancoDados()

	AtualizaPresente, err := db.Prepare("update presente set nome=$1, descricao=$2, preco=$3, link=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaPresente.Exec(nome, descricao, preco, link, id)
	defer db.Close()
}
