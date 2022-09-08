package main

type AllQuestions []SingleQuestion

type SingleQuestion struct {
	EntryURLQueryParam string
	QuestionName       string
	QuestionType       string
	PossibleAnsweres   []PossibleAnswersAndStats
}

type PossibleAnswersAndStats struct {
	Value            string
	Percentage       float64
	AlreadyGenerated float64
	GridTypeAnswers  []SingleQuestion
}

type Answer struct {
	QueryParam string
	Value      string
}

var Questions = AllQuestions{
	{EntryURLQueryParam: "entry.1339899397", QuestionName: "Radio button",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "abc", Percentage: 20},
			{Value: "cde", Percentage: 58},
			{Value: "efg", Percentage: 22},
		},
	},
	{EntryURLQueryParam: "entry.322260907", QuestionName: "Check Box", QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Option 1", Percentage: 30},
			{Value: "Option 2", Percentage: 60},
			{Value: "Option 3", Percentage: 10},
		},
	},
	{EntryURLQueryParam: "entry.1250162232", QuestionName: "Is true?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "True", Percentage: 33},
			{Value: "False", Percentage: 67},
		},
	},
	{EntryURLQueryParam: "entry.670726587", QuestionName: "Linear scale",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 10},
			{Value: "2", Percentage: 20},
			{Value: "3", Percentage: 30},
			{Value: "4", Percentage: 35},
			{Value: "5", Percentage: 5},
		},
	},
	{QuestionName: "mcgrid", QuestionType: "grid",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{GridTypeAnswers: []SingleQuestion{
				{
					EntryURLQueryParam: "entry.1168482335",
					QuestionName:       "abc",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 50},
						{Value: "2", Percentage: 40},
						{Value: "3", Percentage: 10},
					},
				},
				{
					EntryURLQueryParam: "entry.1242191203",
					QuestionName:       "cde",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 10},
						{Value: "2", Percentage: 50},
						{Value: "3", Percentage: 40},
					},
				},
				{
					EntryURLQueryParam: "entry.583794341",
					QuestionName:       "efg",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 10},
						{Value: "2", Percentage: 40},
						{Value: "3", Percentage: 50},
					},
				},
			}},
		},
	},
}

var RealData = AllQuestions{
	{EntryURLQueryParam: "", QuestionName: "Spol",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Žensko", Percentage: 64},
			{Value: "Muško", Percentage: 36},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Dob",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "18-24", Percentage: 28},
			{Value: "25-29", Percentage: 36},
			{Value: "30-34 ", Percentage: 12},
			{Value: "35-39", Percentage: 13},
			{Value: "40-44 ", Percentage: 7},
			{Value: "45-50 ", Percentage: 3},
			{Value: "50 i više", Percentage: 1},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Stručna sprema",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Završena osnovna škola", Percentage: 3},
			{Value: "Završena srednja škola", Percentage: 53},
			{Value: "Viša/visoka stručna sprema", Percentage: 44},
			{Value: "Doktorat", Percentage: 0},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Koliko vremena dnevno provodite koristeći društvene medije?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "do 1 sat", Percentage: 2},
			{Value: "od 2 do 4 sata", Percentage: 63},
			{Value: "5 i više sati", Percentage: 35},
		},
	}, {EntryURLQueryParam: "", QuestionName: "U koju svrhu koristite društvene medije?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Posao", Percentage: 1},
			{Value: "Zabava", Percentage: 67},
			{Value: "Upoznavanje", Percentage: 1},
			{Value: "Kupovina", Percentage: 1},
			{Value: "Video igre", Percentage: 18},
			{Value: "Učenje", Percentage: 4},
			{Value: "Povezane aktivnosti (e-bankinge, e-građani, e-porezna...) ", Percentage: 1},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Koliko često kupujete online?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "svaki dan", Percentage: 0},
			{Value: "jednom tjedno", Percentage: 7},
			{Value: "jednom mjesečno", Percentage: 46},
			{Value: "jednom godišnje", Percentage: 34},
			{Value: "nikada", Percentage: 13},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Ako trgovina u kojoj često kupujete ima webshop (pristup preko web preglednika) i mobilnu aplikaciju koja se instalira na vaš mobitel, koji opciju ćete radije odabrati?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Webshop putem pregledinka", Percentage: 68},
			{Value: "Mobilna aplikacija", Percentage: 32},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Što vam je najbitnije kod odabira webshopa na kojem kupujete?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Sigurnost i zaštita podataka prilikom kupnje", Percentage: 36},
			{Value: "Pozitivne recenzije potrošača", Percentage: 18},
			{Value: "Mogućnost chat pomoći kod odabira proizvoda", Percentage: 3},
			{Value: "Mogućnost preuzimanja robe uposlovnici", Percentage: 0},
			{Value: "Mogućnost plaćanja koje preferiram", Percentage: 23},
			{Value: "Mogućnost kupnje mobitelom", Percentage: 16},
			{Value: "Mogućnost pristupa virtualnom tehnologijom", Percentage: 4},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Da li prema recenzijama na društvenim medijima formirate svoje mišljenje, stavove i vrednovanje proizvoda, koje tada utječu na odluku o kupovini?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Uvijek", Percentage: 41},
			{Value: "Rijetko", Percentage: 57},
			{Value: "Nikada", Percentage: 2},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Što Vas je potaknulo na korištenje društvenih medija za kupovinu?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Praćenje omiljenih influencera", Percentage: 14},
			{Value: "Praćenje omiljnih celebrity-a", Percentage: 2},
			{Value: "Traženje pogodnosti i nižih cijena", Percentage: 41},
			{Value: "Ekskluzivni i limitirani paketi proizvoda", Percentage: 7},
			{Value: "Pronalazak novih proizvoda", Percentage: 6},
			{Value: "Dobivanje više informacija o proizvodima", Percentage: 19},
			{Value: "Razmjena iskustava i mišljenja o proizvodu s drugim potrošačima", Percentage: 11},
			{Value: "Ne koristim društvene medije", Percentage: 0},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Proizvodi koji su promovirani na društvenim medijima imaju višu reputaciju i prodavaniji su od ostalih proizvoda.",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 0},
			{Value: "2", Percentage: 3},
			{Value: "3", Percentage: 18},
			{Value: "4", Percentage: 56},
			{Value: "5", Percentage: 23},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Koristite li neke od navedenih komponenata online kupovine?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Kupovina preporučenih proizvoda na video platformama (Tik Tok, Youtube) ", Percentage: 21},
			{Value: "Kupovina preporučenih proizvoda na društvenim mrežama (Facebook, Pinterest, Instagram, LinkedIn) ", Percentage: 57},
			{Value: "Kupovina korištenjem glasovnih asistenata (Google, Siri, Alexa) ", Percentage: 0},
			{Value: "Kupovina uz pomoć chat aplikacija (Messenger, WhatsApp, Viber) ", Percentage: 5},
			{Value: "Kupovina digitalnih proizvoda u sklopu video igara", Percentage: 16},
			{Value: "Kupovina korištenjem VR pomagala", Percentage: 0},
			{Value: "Ne koristim", Percentage: 1},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Jeste li upoznati s terminom virtualne stvarnosti i Metaverse-a?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Da", Percentage: 45},
			{Value: "Ne", Percentage: 39},
			{Value: "Nisam siguran/na", Percentage: 16},
		},
	}, {EntryURLQueryParam: "", QuestionName: "U kojoj mjeri koristite virtualnu stvarnost?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 43},
			{Value: "2", Percentage: 38},
			{Value: "3", Percentage: 19},
			{Value: "4", Percentage: 0},
			{Value: "5", Percentage: 0},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Jeste li čuli/koristili neke od navedenih VR headset-ova da bi pristupili virtualnoj stvarnosti?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Digitalne naočele (npr. Sony PlayStation VR) ", Percentage: 48},
			{Value: "Oculus Go uređaj", Percentage: 12},
			{Value: "Nintendo Labo VR Kit uređaj", Percentage: 0},
			{Value: "HTC Vive Cosmos uređaj", Percentage: 0},
			{Value: "Lenovo Mirage Solo With Daydream uređaj", Percentage: 0},
			{Value: "Valve Index (Steam) uređaj", Percentage: 0},
			{Value: "Winows Mixed Reality uređaj", Percentage: 0},
			{Value: "VR odijela sa senzorima", Percentage: 23},
			{Value: "Nisam čuo/la, niti sam koristio/la", Percentage: 17},
		},
	}, {QuestionName: " Koliko često koristite virtualnu stvarnost za navedene aktivnosti?",
		QuestionType: "grid",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{GridTypeAnswers: []SingleQuestion{
				{
					EntryURLQueryParam: "",
					QuestionName:       "Igranje video igara",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 50},
						{Value: "2", Percentage: 40},
						{Value: "3", Percentage: 10},
						{Value: "4", Percentage: 50},
						{Value: "5", Percentage: 40},
					},
				},
				{
					EntryURLQueryParam: "",
					QuestionName:       "Druženje i zabava",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 50},
						{Value: "2", Percentage: 40},
						{Value: "3", Percentage: 10},
						{Value: "4", Percentage: 50},
						{Value: "5", Percentage: 40},
					},
				},
				{
					EntryURLQueryParam: "",
					QuestionName:       "Trgovina",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 50},
						{Value: "2", Percentage: 40},
						{Value: "3", Percentage: 10},
						{Value: "4", Percentage: 50},
						{Value: "5", Percentage: 40},
					},
				},
				{
					EntryURLQueryParam: "",
					QuestionName:       "Praćenje i prisustvovanje virtualnim događajima",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 50},
						{Value: "2", Percentage: 40},
						{Value: "3", Percentage: 10},
						{Value: "4", Percentage: 50},
						{Value: "5", Percentage: 40},
					},
				},
			}},
		},
	}, {EntryURLQueryParam: "", QuestionName: "Koje od sljedećih aktivnosti u virtualnom 3D svijetu bi Vas zanimale u budućnosti?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Isprobavanje odjeće", Percentage: 32},
			{Value: "Isprobavanje kozmetičkih proizvoda", Percentage: 27},
			{Value: "Isprobavanje naočala/sunčanih naočala", Percentage: 16},
			{Value: "Kupovina s prijateljima", Percentage: 16},
			{Value: "\"Hod\" kroz virtualnu trgovinu (virtualna stvarnost kao pomoć u pregledu i odabiru proizvoda)", Percentage: 34},
			{Value: "Kupovina namještaja", Percentage: 14},
			{Value: "Grupni treninzi", Percentage: 23},
			{Value: "Posjećivanje virtualnih koncerata/sportskih događaja", Percentage: 34},
			{Value: "Kupovina stvari za svog avatara", Percentage: 7},
			{Value: "Ništa od navedenog", Percentage: 6},
		},
	},
}
