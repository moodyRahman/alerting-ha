package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)

		var alert map[string]any
		json.Unmarshal(body, &alert)

		testout, _ := json.MarshalIndent(alert, "", "	")

		fmt.Println(string(testout))

		var alert_arr = alert["alerts"].([]any)

		var email_body string

		for _, x := range alert_arr {
			x := x.(map[string]any)
			ann := x["annotations"].(map[string]any)
			status := x["status"].(string)

			email_body += status + " " + ann["summary"].(string) + ", "
		}

		req, err := http.NewRequest("POST", "http://ntfy:7070/urgent", strings.NewReader(email_body + "end of message"))
		// req, err := http.NewRequest("POST", "http://localhost:5000", strings.NewReader(email_body + "end of message"))


		if (err != nil) {
			fmt.Println(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Email", "moududur921@gmail.com")

		http.DefaultClient.Do(req)
		fmt.Fprint(w, "request processed")

	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "hello world")
	})

	fmt.Println("booting server...")
	http.ListenAndServe(":9001", nil)
}


