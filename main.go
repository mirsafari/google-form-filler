package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/* Get Query params
const formPrefilledURL string =
"https://docs.google.com/forms/d/e/1FAIpQLSc9QKh4qcopFSqAztrvvYyVqsi06IyMWOS9I_OPXCw58kBVJQ/viewform?usp=pp_url&entry.1339899397=cde&entry.322260907=Option+2&entry.322260907=Option+3"
queryParams := parseUrlEntry(formPrefilledURL)
q := req.URL.Query()

for k,v := range queryParams {
	if strings.HasPrefix(k, "entry.") {
		fmt.Println(k, " -> ", v)
	}
}
*/

func main() {

	var responses float64 = 100
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < int(responses)+1; i++ {
		fromAnswers := []Answer{}

		// For each question, generate answer
		for _, question := range Questions {
			// If the quesiton is grid type, we need to iterate over all sub questions in the grid
			if question.QuestionType == "grid" {
				for _, subquestion := range question.PossibleAnsweres[0].GridTypeAnswers {
					// And prepare answer for each sub question
					singleAnswer := Answer{
						QueryParam: subquestion.EntryURLQueryParam,
						Value:      getAnswer(&subquestion, responses, i),
					}
					// Define single response containing answers to all questions
					fromAnswers = append(fromAnswers, singleAnswer)
				}
			} else if question.QuestionType == "multiple-choice" {
				fmt.Println("a")
			} else {
				singleAnswer := Answer{
					QueryParam: question.EntryURLQueryParam,
					Value:      getAnswer(&question, responses, i),
				}
				// Define single response containing answers to all questions
				fromAnswers = append(fromAnswers, singleAnswer)
			}
		}

		// Now that all questions have been answered it is time to send them to Google Forms
		sendRequests(fromAnswers)

		// Delay execution by 2 seconds not to get rate limited by Google
		//time.Sleep(2 * time.Second)
	}
}

func getAnswer(question *SingleQuestion, numberOfResponses float64, currentResponseLoop int) string {
	randomNumber := rand.Intn(len(question.PossibleAnsweres))

	for {
		// Increase counter
		question.PossibleAnsweres[randomNumber].AlreadyGenerated += 1

		// Check if answer percentage has not been reached
		if ((question.PossibleAnsweres[randomNumber].AlreadyGenerated / numberOfResponses) * 100) <= question.PossibleAnsweres[randomNumber].Percentage {
			// If the percantage is lower than desired, generated answer was ok
			if numberOfResponses+1 != float64(currentResponseLoop) {
				break
			}
		}
		// Decrease the counter as it is invalid answer
		question.PossibleAnsweres[randomNumber].AlreadyGenerated -= 1

		// Regenerate number
		randomNumber = rand.Intn(len(question.PossibleAnsweres))
	}

	return question.PossibleAnsweres[randomNumber].Value
}

func sendRequests(answers []Answer) {
	formURL := "https://docs.google.com/forms/d/e/1FAIpQLSc9QKh4qcopFSqAztrvvYyVqsi06IyMWOS9I_OPXCw58kBVJQ/formResponse"

	req, err := http.NewRequest("POST", formURL, nil)

	if err != nil {
		fmt.Println("Error reading request.", err)
	}

	q := req.URL.Query()

	// Get all answers to POST request values
	for _, answer := range answers {
		q.Add(answer.QueryParam, answer.Value)
	}
	req.URL.RawQuery = q.Encode()

	fmt.Print("Sending: ", formatRequest(req))

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println(" - Got HTTP ", resp.StatusCode)

}

func parseUrlEntry(inputURL string) url.Values {

	u, err := url.Parse(inputURL)
	if err != nil {
		log.Fatal(err)
	}
	queryParams := u.Query()

	return queryParams
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// Return the request as a string
	return strings.Join(request, "\n")
}
