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
	{EntryURLQueryParam: "entry.1457946910", QuestionName: "Spol",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Muško", Percentage: 64},
			{Value: "Žensko", Percentage: 36},
		},
	}, {EntryURLQueryParam: "entry.1646221043", QuestionName: "Dob",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "18-25", Percentage: 28},
			{Value: "26-35", Percentage: 36},
			{Value: "36-45", Percentage: 12},
			{Value: "46-55", Percentage: 13},
			{Value: "56-65", Percentage: 7},
			{Value: "65+", Percentage: 3},
		},
	}, {EntryURLQueryParam: "entry.973358507", QuestionName: "Stručna sprema",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "NKV", Percentage: 3},
			{Value: "KV", Percentage: 53},
			{Value: "SSS", Percentage: 44},
			{Value: "VŠS", Percentage: 0},
			{Value: "VSS", Percentage: 0},
			{Value: "Doktorat", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.263059413", QuestionName: "Jeste li upoznati s pojmom Crowdfundinga odnosno skupnog financiranja?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Da", Percentage: 2},
			{Value: "Ne", Percentage: 63},
			{Value: "Nisam siguran/na", Percentage: 35},
		},
	}, {EntryURLQueryParam: "entry.1018478988", QuestionName: "Jeste li ikad sudjelovali u nekoj od Crowdfunding kampanji?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Da", Percentage: 30},
			{Value: "Ne", Percentage: 70},
		},
	}, {EntryURLQueryParam: "entry.99839994", QuestionName: "Jeste li čuli/koristili koje od navedenih Crowdfunding platformi?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Funderbeam SEE", Percentage: 0},
			{Value: "Croinvest.eu", Percentage: 0},
			{Value: "Capital.hr", Percentage: 0},
			{Value: "Croenergy.eu", Percentage: 0},
			{Value: "Kickstarter", Percentage: 0},
			{Value: "GoFundMe", Percentage: 0},
			{Value: "TravelStarter", Percentage: 0},
			{Value: "Indiegogo", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.310576124", QuestionName: "U kojoj mjeri smatrate da je investiranje u Crowdfunding rizično?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 0},
			{Value: "2", Percentage: 0},
			{Value: "3", Percentage: 0},
			{Value: "4", Percentage: 0},
			{Value: "5", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.555498744", QuestionName: "Slažete li se da je Crowdfunding dobro rješenje za projekte koji nemaju dovoljnu financijsku podršku, a važni su za dobrobit zajednice?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 0},
			{Value: "2", Percentage: 0},
			{Value: "3", Percentage: 0},
			{Value: "4", Percentage: 0},
			{Value: "5", Percentage: 0},
		},
	}, {QuestionName: "U kojoj mjeri bi navedenim projektima dali prioritet kada biste ulagali u Crowdfunding?",
		QuestionType: "grid",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{GridTypeAnswers: []SingleQuestion{
				{
					EntryURLQueryParam: "entry.1356236123",
					QuestionName:       "Znanost",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 79},
						{Value: "2", Percentage: 0},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 5},
						{Value: "5", Percentage: 16},
					},
				},
				{
					EntryURLQueryParam: "entry.973627734",
					QuestionName:       "Sport",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 89},
						{Value: "2", Percentage: 7},
						{Value: "3", Percentage: 4},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.1853526166",
					QuestionName:       "Kultura i umjetnost",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 93},
						{Value: "2", Percentage: 7},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.1197574119",
					QuestionName:       "Zdravstvo",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 98},
						{Value: "2", Percentage: 2},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.1650394690",
					QuestionName:       "Obnova i zaštita kulturne baštine",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 98},
						{Value: "2", Percentage: 2},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.1621881913",
					QuestionName:       "Humanitarni projekti",
					PossibleAnsweres: []PossibleAnswersAndStats{
						{Value: "1", Percentage: 98},
						{Value: "2", Percentage: 2},
						{Value: "3", Percentage: 0},
						{Value: "4", Percentage: 0},
						{Value: "5", Percentage: 0},
					},
				},
				{
					EntryURLQueryParam: "entry.739125279",
					QuestionName:       "Zaštita okoliša",
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
	}, {EntryURLQueryParam: "entry.106112672", QuestionName: "Kojim nositeljima projekata bi dali prednost kada biste ulagali u Crowdfunding?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Fizičkim osobama", Percentage: 0},
			{Value: "Neprofitnim organizacijama", Percentage: 0},
			{Value: "Volonterskim/dobrotvornim organizacijama", Percentage: 0},
			{Value: "Poslovni sektor", Percentage: 0},
			{Value: "Akademski sektor", Percentage: 0},
		},
	}, {EntryURLQueryParam: "entry.89483892", QuestionName: "Smatrate li da će Crowdfunding platforme u budućnosti biti veća konkurencija klasičnim modelima financiranja (bankovni krediti, leasing, strateški partneri i drugo) nego danas?",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "Da", Percentage: 0},
			{Value: "Ne", Percentage: 0},
			{Value: "Nisam siguran/na", Percentage: 0},
		},
	},
}
