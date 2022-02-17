package sexpr

import "fmt"

func ExampleMarshal() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     0,
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	data, _ := Marshal(strangelove)
	fmt.Println(string(data))

	// Output:
	// ((Title "Dr. Strangelove")
	//  (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
	//  (Oscars ("Best Actor (Nomin.)"
	//           "Best Adapted Screenplay (Nomin.)"
	//           "Best Director (Nomin.)"
	//           "Best Picture (Nomin.)")))
}
