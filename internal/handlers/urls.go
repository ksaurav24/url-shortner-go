package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type URL struct {
	ID        string     `json:"id"`
	LongURL   string     `json:"longUrl"`
	CreatedAt time.Time  `json:"createdAt"`
	ExpiresAt *time.Time `json:"expiresAt"`
	Status    string     `json:"status"`
}

// urlsHandler struct
type UrlHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewUrlHandler(db *sql.DB, logger *slog.Logger) *UrlHandler {
	return &UrlHandler{
		db:     db,
		logger: logger,
	}
}

// GetUrls retrieves the latest URLs.
//
// The returned handler queries the database and returns
// the results as JSON.
func (uh UrlHandler) GetUrls(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rows, err := uh.db.QueryContext(ctx,
		`
			SELECT id, long_url, created_at, expires_at, status 
			FROM urls
			ORDER BY created_at DESC
			LIMIT 100
		`)

	if err != nil {
		uh.logger.Error("Database query failure", "err", err)
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
			uh.logger.Error("rows scan failure", "err", err)
			http.Error(w, "Failed to retrieve the urls", http.StatusInternalServerError)
			return
		}

		urls = append(urls, u)
	}

	log.Println(urls)

	if err := rows.Err(); err != nil {
		uh.logger.Error("rows.err", "err", err)
		http.Error(w, "Failed to retrieve urls", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(urls)
}
