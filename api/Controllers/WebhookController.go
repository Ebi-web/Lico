package Controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"lineBot/Services/Line"
	"lineBot/domain/usecase"
)

type Request struct {
	Destination string `json:"destination"`
	Events      []struct {
		Type    string `json:"type"`
		Message struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"message"`
		WebhookEventID  string `json:"webhookEventId"`
		DeliveryContext struct {
			IsRedelivery bool `json:"isRedelivery"`
		} `json:"deliveryContext"`
		Timestamp int64 `json:"timestamp"`
		Source    struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		ReplyToken string `json:"replyToken"`
		Mode       string `json:"mode"`
	} `json:"events"`
}

type reply struct {
	ReplyToken string `json:"replyToken"`
	Messages   []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Read the request body as a byte slice
	webhookJson, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Parse the request body into a Request struct
	var webhookStruct Request
	err = json.Unmarshal(webhookJson, &webhookStruct)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	replyStruct := reply{ReplyToken: webhookStruct.Events[0].ReplyToken, Messages: []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}{
		{
			Type: "text",
			Text: usecase.PickMember(),
		},
	}}
	body, err := json.Marshal(replyStruct)
	if err != nil {
		fmt.Println(err)
		return
	}

	reply, err := http.NewRequest("POST", Line.ReplyEndpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	reply.Header.Set("Content-Type", "application/json")
	reply.Header.Set("Authorization", "Bearer "+Line.ChannelAccessToken)

	client := &http.Client{}
	resp, err := client.Do(reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Set the appropriate HTTP status code
	w.WriteHeader(http.StatusOK)
}
