package models

type Alunos struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
	Rg   string `json:"rg"`
	Cpf  string `json:"cpf"`
}
