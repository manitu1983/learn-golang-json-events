package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type pocket struct {
	Name      string    `json:"name,omitempty"`
	Reference string    `json:"reference,omitempty"`
	LastUsage time.Time `json:"lastUsage,omitempty"`
}

type StartedEvent struct {
	Type      string    `json:"type,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Pocket    pocket    `json:"pocket,omitempty"`
}

type RejectedEvent struct {
	Type               string    `json:"type,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	RejectionCode      string    `json:"rejectionCode,omitempty"`
	CancelledBy        string    `json:"cancelledBy,omitempty"`
	ProviderFailReason string    `json:"providerFailReason,omitempty"`
}

type TypeEvent struct {
	Type string `json:"type,omitempty"`
}

func main() {
	inputEvents := [3]string{
		`{
			"type": "started",
			"createdAt": "2025-01-01T00:00:00Z",
			"pocket": {
        		"name": "411111******1111",
        		"reference": "1234sadf",
        		"lastUsage": "2024-01-01T00:00:00Z"
      		}
		}`,
		`{"type": "someOldEvent"}`,
		`{
			"type": "rejected",
			"createdAt": "2025-01-01T00:10:00Z",
			"rejectionCode": "nogEnoughMoney",
			"cancelledBy": "provider",
			"providerFailReason": "NO_MONEY"
		}`,
	}

	for _, inputEvent := range inputEvents {
		var typeEvent TypeEvent
		if err := json.Unmarshal([]byte(inputEvent), &typeEvent); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(typeEvent.Type)

		switch typeEvent.Type {
		case "started":
			var event StartedEvent
			if err := json.Unmarshal([]byte(inputEvent), &event); err != nil {
				log.Fatalln(err)
			}
			fmt.Println(event)
		case "rejected":
			var event RejectedEvent
			if err := json.Unmarshal([]byte(inputEvent), &event); err != nil {
				log.Fatalln(err)
			}
			fmt.Println(event)
		default:
			fmt.Println("unknown event type:", typeEvent.Type)
		}

		fmt.Println()
	}
}
