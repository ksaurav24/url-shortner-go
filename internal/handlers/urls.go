package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type URL struct {
	ID        string  `json:"id"`
	LongURL   string  `json:"longUrl"`
	CreatedAt string  `json:"createdAt"`
	ExpiresAt *string `json:"expiresAt"`
	Status    string  `json:"status"`
}

func GetUrls(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(`
			SELECT id, long_url, created_at, expires_at, status 
			FROM urls
			ORDER BY created_at DESC
			LIMIT 100
		`)

		if err != nil {
			log.Printf("DB Query Error: %v", err)
			http.Error(w, "Failed to retrieve urls", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		urls := []URL{}

		for rows.Next() {
			var u URL
			err := rows.Scan(
				&u.ID,
				&u.LongURL,
				&u.CreatedAt,
				&u.ExpiresAt,
				&u.Status,
			)
			if err != nil {
				log.Printf("Failed to scan the rows: %v", err)
				http.Error(w, "Failed to retrieve the urls", http.StatusInternalServerError)
				return
			}

			urls = append(urls, u)
		}

		log.Println(urls)

		if err := rows.Err(); err != nil {
			log.Printf("rows.err: %v", err)
			http.Error(w, "Failed to retrieve urls", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(urls)
	}
}
