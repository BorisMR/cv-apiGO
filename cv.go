package main

type CVs []cv

type cv struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Background []exp
	Education  []edu
}

type background []exp

type exp struct {
	Position    string `json:"position"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

type education []edu

type edu struct {
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
	Year        string `json:"year"`
}

//mock object for testing purposes
var cvs = CVs{
	cv{
		"Boris Morales Rios",
		"+56 9 45296652",
		"anum.bmr.01@gmail.com",
		"Garibaldi 0935",
		background{
			exp{
				"Software developer",
				"CyM Contulmo",
				"Chile, Temuco",
				"2019 - present",
				"I Work here as a full-stack developer, using tracking technologies for public transport management",
			},
			exp{
				"Web developer",
				"Instituto de Informatica Educativa",
				"Chile, Temuco",
				"2017 - 2019",
				"I Worked as a frontend and backend on platforms from Endfid-2017 a national teachers knowledgment validation process, also I made Scripts development to automatize personal reports generation",
			},
			exp{
				"Software developer",
				"Lirmi Chile SPA",
				"Chile, Temuco",
				"2016 - 2017",
				"I did my main internship developing a module for a teaching planification assistant platform, this module does feedback improvements on planification tasks.",
			},
			exp{
				"Internship II",
				"Instituto de Informatica Educativa",
				"Chile, Temuco",
				"2016 - 2016",
				"I did my first internship doing tasks related to Information Security",
			},
			exp{
				"Internship I",
				"Instituto de Informatica Educativa",
				"Chile, Temuco",
				"2015 - 2016",
				"I did my first internship as developer in the project Academia PSU working in a webapp used to provide effective follow-up from class sessions",
			},
			exp{
				"Assistant",
				"Instituto de Informatica Educativa",
				"Chile, Temuco",
				"2014 - 2014",
				"TICEdu project assistant, developing scripts in JavaScript and performing analysis of audio using a voice recognition tool derived from LIUM.",
			},
			exp{
				"Assistant",
				"Instituto de Informatica Educativa",
				"Chile, Temuco",
				"2014 - 2014",
				"TEDI project assistant, developing tools for an digital interactive schoolar books platform.",
			},
		},
		education{
			edu{
				"Computer Science / Information Systems",
				"University of the Frontier",
				"2016",
			},
			edu{
				"Licentiate in Engineering Sciences",
				"University of the Frontier",
				"2016",
			},
		},
	},
}
