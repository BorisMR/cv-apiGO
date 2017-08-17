package main

type Cv struct {
	Nombre      string  `json:"nombre"`
	Fono        string  `json:"fono"`
	Email       string  `json:"email"`
	Direccion   string  `json:"direccion"`
}

type CVs []Cv

var cvs = CVs{

	Cv{
		"Boris Morales Rios",
		"+56 9 45296652",
		"anum.bmr.01@gmail.com",
		"Garibaldi 0935",
	},
}