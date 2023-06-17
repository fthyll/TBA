package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getToken(teks string, j *int) string {
	k := len(teks)
	kata := ""

	// abaikan spasi dan pindah baris
	for jVal := *j; jVal < k && (teks[jVal] == ' ' || teks[jVal] == '\r' || teks[jVal] == '\n'); jVal++ {
		*j++
	}

	// ambil 1 kata/token
	for jVal := *j; jVal < k && (teks[jVal] != ' ' && teks[jVal] != '\r' && teks[jVal] != '\n'); jVal++ {
		if teks[jVal] == '>' {
			if kata != "" {
				return kata
			} else {
				*j++
				if teks[*j] == '=' {
					*j++
					return ">="
				} else {
					return ">"
				}
			}
		} else if teks[jVal] == '<' {
			if kata != "" {
				return kata
			} else {
				*j++
				if teks[*j] == '=' {
					*j++
					return "<="
				} else {
					return "<"
				}
			}
		} else if teks[jVal] == '=' {
			if kata != "" {
				return kata
			} else {
				*j++
				return "="
			}
		} else if teks[jVal] == '+' {
			if kata != "" {
				return kata
			} else {
				*j++
				return "+"
			}
		} else if teks[jVal] == '-' {
			if kata != "" {
				return kata
			} else {
				*j++
				return "-"
			}
		}
		kata += string(teks[jVal])
		*j++
	}
	return kata
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	teks := r.FormValue("teks")

	i := 0
	var tokens []string
	k := len(teks)
	for i < k {
		token := getToken(teks, &i)
		tokens = append(tokens, token)
	}

	response := strings.Join(tokens, "\n")
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/scan", scanHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
