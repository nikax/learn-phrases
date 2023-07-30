package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/save-phrase", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		dec := json.NewDecoder(r.Body)
		var payload struct {
			SourcePhrase     string              `json:"source_phrase"`
			SourceLang       string              `json:"source_lang"`
			TranslatedPhrase string              `json:"translated_phrase"`
			TranslatedLang   string              `json:"translated_lang"`
			NewWords         map[string][]string `json:"new_words"`
		}
		err := dec.Decode(&payload)
		if err == nil {
			dir, _ := os.Getwd()
			phrasesFilePath := dir + "/../../data/phrases.txt"
			wordsFilePath := dir + "/../../data/words.txt"

			phrasesFile, _ := os.OpenFile(phrasesFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
			defer phrasesFile.Close()

			_, err = phrasesFile.WriteString(payload.SourcePhrase + "|" + payload.TranslatedPhrase + "\n")
			if err != nil {
				fmt.Println("Can not save phrase to a file")
				fmt.Println(err)
			}

			if len(payload.NewWords) > 0 {
				wordsFile, _ := os.OpenFile(wordsFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
				defer wordsFile.Close()

				for word, translations := range payload.NewWords {
					_, err := wordsFile.WriteString(word + "|" + strings.Join(translations, "|") + "\n")
					if err != nil {
						fmt.Println("Can not save word to a file")
						fmt.Println(err)
					}
				}
			}
		} else {
			fmt.Println(err)
		}
	})

	http.ListenAndServe("127.0.0.1:7777", nil)
}
