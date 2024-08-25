package handlers

import (
    "encoding/json"
    "net/http"
	"github.com/mahdi-eth/history-storage-service/internal/config/db"
)

func GetHistory(w http.ResponseWriter, r *http.Request) {
    messages, err := db.RedisClient.LRange(db.Ctx, "chat_history", 0, 99).Result()
    if err != nil {
        http.Error(w, "Failed to retrieve chat history", http.StatusInternalServerError)
        return
    }

    for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
        messages[i], messages[j] = messages[j], messages[i]
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(messages); err != nil {
        http.Error(w, "Failed to encode chat history", http.StatusInternalServerError)
    }
}
