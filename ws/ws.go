package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"vote.app/m/db"
	"vote.app/m/graph/model"
)

// upgrader ใช้สำหรับอัปเกรด HTTP connection เป็น WebSocket connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// อนุญาตให้ทุกการเชื่อมต่อ
		return true
	},
}

// // formatMessageType แปลง message type เป็นสตริง
// func formatMessageType(mt int) string {
// 	switch mt {
// 	case websocket.TextMessage:
// 		return "TextMessage"
// 	case websocket.BinaryMessage:
// 		return "BinaryMessage"
// 	case websocket.CloseMessage:
// 		return "CloseMessage"
// 	case websocket.PingMessage:
// 		return "PingMessage"
// 	case websocket.PongMessage:
// 		return "PongMessage"
// 	default:
// 		return "UnknownMessageType"
// 	}
// }

// sendMessages ส่งข้อความไปยัง WebSocket client ทุกวินาที
func VoteList(ctx context.Context) ([]*model.VoteList, error) {
	dsn := db.GetGormDB()

	// Fetch all db.Vote records and handle potential errors
	var dbVotes []db.Vote
	err := dsn.Find(&dbVotes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch votes: %w", err) // Wrap error with context
	}

	// Convert db.Vote to model.VoteList efficiently using slices
	modelVoteList := make([]*model.VoteList, 0, len(dbVotes)) // Pre-allocate for efficiency
	for _, dbVote := range dbVotes {
		modelVote := &model.VoteList{
			ID:      int(dbVote.ID),
			Name:    dbVote.Name,
			Number:  int(dbVote.Number),
			Details: dbVote.Details,
			LogoURL: dbVote.LogoUrl,
			Score:   int(dbVote.Score),
		}
		modelVoteList = append(modelVoteList, modelVote)
	}

	return modelVoteList, nil
}

// sendMessages function to send messages to WebSocket client
func sendMessages(c *websocket.Conn, voteList []*model.VoteList) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for _ = range ticker.C {
		// Convert voteList to JSON format
		voteListData, err := json.Marshal(voteList)
		if err != nil {
			log.Println("error marshaling voteList:", err)
			return
		}

		// Send voteList data as JSON message
		message := []byte(voteListData)
		err = c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			return
		}
		log.Printf("sent vote data: %s", message)
	}
}

// echo function to handle WebSocket connection and send vote data
func echo(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to WebSocket
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// Fetch voteList data
	voteList, err := VoteList(context.TODO()) // Assuming VoteList takes context
	if err != nil {
		log.Println("error fetching voteList:", err)
		return
	}

	// Send voteList data to WebSocket client every second
	sendMessages(c, voteList)
}
func main() {
	http.HandleFunc("/ws", echo)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
