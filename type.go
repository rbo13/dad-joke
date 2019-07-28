package joke

import "fmt"

// DadJoke represents a good 'ol dad joke.
type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func (dj DadJoke) String() string {
	json := `{ 
	 id: %s, 
	 joke: %s, 
	 status: %d,
}`

	return fmt.Sprintf(json, dj.ID, dj.Joke, dj.Status)
}
