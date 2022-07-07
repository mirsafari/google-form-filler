package main

type AllQuestions []SingleQuestion

type SingleQuestion struct {
	EntryURLQueryParam string
	QuestionName       string
	PossibleAnsweres   []PossibleAnswersAndStats
}

type PossibleAnswersAndStats struct {
	Value            string
	Percentage       float64
	AlreadyGenerated float64
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
	{EntryURLQueryParam: "entry.322260907", QuestionName: "Check Box",
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
	{EntryURLQueryParam: "entry.670726587", QuestionName: "mcgrid",
		PossibleAnsweres: []PossibleAnswersAndStats{
			{Value: "1", Percentage: 10},
			{Value: "2", Percentage: 20},
			{Value: "3", Percentage: 30},
			{Value: "4", Percentage: 35},
			{Value: "5", Percentage: 5},
		},
	},
}
