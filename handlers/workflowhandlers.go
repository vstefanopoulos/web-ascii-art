package handlers

import (
	"ascii-art-web-stylize/asciiart"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Handler για την αρχική σελίδα
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil || r.URL.Path != "/" {
		error404(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

// Handler για την επεξεργασία του POST request και την εμφάνιση του ASCII art
func asciiArt(w http.ResponseWriter, r *http.Request) {
	// Έλεγχος αν το method είναι POST
	if r.Method != http.MethodPost {
		error405(w, r) // 405 Method Not Allowed
		return
	}

	// Λήψη δεδομένων από το HTML form
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	alignment := r.FormValue("alignment")

	// Έλεγχος για κενά πεδία
	if text == "" || banner == "" || alignment == "" {
		error400(w, r) // 400 Bad Request
		return
	}

	// Καθορισμός της διαδρομής του banner
	bannerPath := filepath.Join("banners", banner+".txt")

	// Έλεγχος αν το αρχείο banner υπάρχει
	if _, err := os.Stat(bannerPath); os.IsNotExist(err) {
		error404(w, r) // 404 Not Found αν δεν βρεθεί το αρχείο
		return
	}

	// Εκτέλεση της συνάρτησης για ASCII art
	var result string
	result, err := asciiart.GenerateAsciiArt(text, bannerPath, alignment)
	if err != nil {
		fmt.Println(err)
		errorMessage := fmt.Sprint(err)
		if errorMessage == "Bad request" {
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				error404(w, r) // 404 Not Found αν δεν βρεθεί το template
				return
			}
			result = "Error: Only ASCII characters allowed\n\n\tTry something else"
			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, map[string]interface{}{
				"Result":    result,
				"Text":      text,
				"Banner":    banner,
				"Alignment": alignment,
			})
			// error400(w, r)
		} else {
			// 500 Internal Server Error αν υπάρχει σφάλμα κατά τη δημιουργία του ASCII art
			error500(w, r)
		}
		return
	}

	// Φορτώνουμε το template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		error404(w, r) // 404 Not Found αν δεν βρεθεί το template
		return
	}

	// Επιστρέφουμε το αποτέλεσμα και τα δεδομένα πίσω στην HTML σελίδα
	w.WriteHeader(http.StatusOK) // 200 OK
	tmpl.Execute(w, map[string]interface{}{
		"Result":    result,
		"Text":      text,
		"Banner":    banner,
		"Alignment": alignment,
	})
}

func StartServer() {
	// Serve static files from the "templates" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ascii-art", asciiArt)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
