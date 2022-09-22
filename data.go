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

var RealData = AllQuestions{
	{EntryURLQueryParam: "entry.1047804639", QuestionName: "Spol",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Žensko", Percentage: 64},
			{Value: "Muško", Percentage: 36},
		},
	}, {EntryURLQueryParam: "entry.2112220968", QuestionName: "Dob",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "18-24", Percentage: 28},
			{Value: "25-29", Percentage: 36},
			{Value: "30-34", Percentage: 12},
			{Value: "35-39", Percentage: 13},
			{Value: "40-44", Percentage: 7},
			{Value: "45-50", Percentage: 3},
			{Value: "50 i više", Percentage: 1},
		},
	}, {EntryURLQueryParam: "entry.2111576696", QuestionName: "Stručna sprema",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Završena osnovna škola", Percentage: 3},
			{Value: "Završena srednja škola", Percentage: 53},
			{Value: "Viša/visoka stručna sprema", Percentage: 44},
			{Value: "Doktorat", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.1294115916", QuestionName: "Koliko vremena dnevno provodite koristeći društvene medije?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "do 1 sat", Percentage: 2},
			{Value: "od 2 do 4 sata", Percentage: 63},
			{Value: "5 i više sati", Percentage: 35},
		},
	}, {EntryURLQueryParam: "entry.1299024644", QuestionName: "U koju svrhu koristite društvene medije?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Posao", Percentage: 3},
			{Value: "Zabava", Percentage: 45},
			{Value: "Upoznavanje", Percentage: 8},
			{Value: "Kupovina", Percentage: 3},
			{Value: "Video igre", Percentage: 38},
			{Value: "Učenje", Percentage: 2},
			{Value: "Povezane aktivnosti (e-banking, e-građani, e-porezna...)", Percentage: 1},
		},
	}, {EntryURLQueryParam: "entry.1665770947", QuestionName: "Koliko često kupujete online?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "svaki dan", Percentage: 0},
			{Value: "jednom tjedno", Percentage: 7},
			{Value: "jednom mjesečno", Percentage: 46},
			{Value: "jednom godišnje", Percentage: 34},
			{Value: "nikada", Percentage: 13},
		},
	}, {EntryURLQueryParam: "entry.1082389360", QuestionName: "Ako trgovina u kojoj često kupujete ima webshop (pristup preko web preglednika) i mobilnu aplikaciju koja se instalira na vaš mobitel, koji opciju ćete radije odabrati?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Webshop putem pregledinka", Percentage: 68},
			{Value: "Mobilna aplikacija", Percentage: 32},
		},
	}, {EntryURLQueryParam: "entry.1367234296", QuestionName: "Što vam je najbitnije kod odabira webshopa na kojem kupujete?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Sigurnost i zaštita podataka prilikom kupnje", Percentage: 36},
			{Value: "Pozitivne recenzije potrošača", Percentage: 18},
			{Value: "Mogućnost chat pomoći kod odabira proizvoda", Percentage: 5},
			{Value: "Mogućnost preuzimanja robe uposlovnici", Percentage: 0},
			{Value: "Mogućnost plaćanja koje preferiram", Percentage: 23},
			{Value: "Mogućnost kupnje mobitelom", Percentage: 16},
			{Value: "Mogućnost pristupa virtualnom tehnologijom", Percentage: 2},
		},
	}, {EntryURLQueryParam: "entry.1296490983", QuestionName: "Da li prema recenzijama na društvenim medijima formirate svoje mišljenje, stavove i vrednovanje proizvoda, koje tada utječu na odluku o kupovini?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Uvijek", Percentage: 41},
			{Value: "Rijetko", Percentage: 57},
			{Value: "Nikada", Percentage: 2},
		},
	}, {EntryURLQueryParam: "entry.1967503838", QuestionName: "Što Vas je potaknulo na korištenje društvenih medija za kupovinu?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Praćenje omiljenih influencera", Percentage: 14},
			{Value: "Praćenje omiljnih celebrity-a", Percentage: 2},
			{Value: "Traženje pogodnosti i nižih cijena", Percentage: 41},
			{Value: "Ekskluzivni i limitirani paketi proizvoda", Percentage: 7},
			{Value: "Pronalazak novih proizvoda", Percentage: 6},
			{Value: "Dobivanje više informacija o proizvodu", Percentage: 19},
			{Value: "Razmjena iskustava i mišljenja o proizvodu s drugim potrošačima", Percentage: 11},
			{Value: "Ne koristim društvene medije", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.1405491769", QuestionName: "Proizvodi koji su promovirani na društvenim medijima imaju višu reputaciju i prodavaniji su od ostalih proizvoda.",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 0},
			{Value: "2", Percentage: 3},
			{Value: "3", Percentage: 18},
			{Value: "4", Percentage: 56},
			{Value: "5", Percentage: 23},
		},
	}, {EntryURLQueryParam: "entry.1490922471", QuestionName: "Koristite li neke od navedenih komponenata online kupovine?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Kupovina preporučenih proizvoda na video platformama (Tik Tok, Youtube)", Percentage: 21},
			{Value: "Kupovina preporučenih proizvoda na društvenim mrežama (Facebook,  Pinterest, Instagram, LinkedIn)", Percentage: 55},
			{Value: "Kupovina korištenjem glasovnih asistenata (Google, Siri, Alexa)", Percentage: 0},
			{Value: "Kupovina uz pomoć chat aplikacija (Messenger, WhatsApp, Viber)", Percentage: 5},
			{Value: "Kupovina digitalnih proizvoda u sklopu video igara", Percentage: 16},
			{Value: "Kupovina korištenjem VR pomagala", Percentage: 2},
			{Value: "Ne koristim", Percentage: 1},
		}, // Last checkpoint
	}, {EntryURLQueryParam: "entry.841651687", QuestionName: "Jeste li upoznati s terminom virtualne stvarnosti i Metaverse-a?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Da", Percentage: 38},
			{Value: "Ne", Percentage: 21},
			{Value: "Nisam siguran/na", Percentage: 41},
		},
	}, {EntryURLQueryParam: "entry.1896937933", QuestionName: "U kojoj mjeri koristite virtualnu stvarnost?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 43},
			{Value: "2", Percentage: 33},
			{Value: "3", Percentage: 19},
			{Value: "4", Percentage: 5},
			{Value: "5", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.890103158", QuestionName: "Jeste li čuli/koristili neke od navedenih VR headset-ova da bi pristupili virtualnoj stvarnosti?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Digitalne naočele (npr. Sony PlayStation VR)", Percentage: 48},
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
					EntryURLQueryParam: "entry.470228456",
					QuestionName:       "Igranje video igara",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 79},
						{Value: "2", Percentage: 0},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 5},
						{Value: "5", Percentage: 16},
					},
				},
				{
					EntryURLQueryParam: "entry.387018731",
					QuestionName:       "Druženje i zabava",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 89},
						{Value: "2", Percentage: 7},
						{Value: "3", Percentage: 4},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.727932485",
					QuestionName:       "Trgovina",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 93},
						{Value: "2", Percentage: 7},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.432155184",
					QuestionName:       "Praćenje i prisustvovanje virtualnim događajima",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 98},
						{Value: "2", Percentage: 2},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
			}},
		},
	}, {EntryURLQueryParam: "entry.669798995", QuestionName: "Koje od sljedećih aktivnosti u virtualnom 3D svijetu bi Vas zanimale u budućnosti?",
		QuestionType: "multiple-choice",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Isprobavanje odjeće", Percentage: 16},
			{Value: "Isprobavanje kozmetičkih proizvoda", Percentage: 12},
			{Value: "Isprobavanje naočala/sunčanih naočala", Percentage: 11},
			{Value: "Kupovina s prijateljima", Percentage: 10},
			{Value: "\"Hod\" kroz virtualnu trgovinu (virtualna stvarnost kao pomoć u pregledu i odabiru proizvoda)", Percentage: 9},
			{Value: "Kupovina namještaja", Percentage: 9},
			{Value: "Grupni treninzi", Percentage: 8},
			{Value: "Posjećivanje virtualnih koncerata/sportskih događaja", Percentage: 18},
			{Value: "Kupovina stvari za svog avatara", Percentage: 4},
			{Value: "Ništa od navedenog", Percentage: 2},
		},
	},
}
