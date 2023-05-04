package model

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestOffer(t *testing.T) {
	offer := &Offer{
		RewardAmount:       100,
		CustomerId:         1,
		CategoryName:       "外卖",
		PickupLocationId:   1,
		DeliveryLocationId: 2,
		TimeLimit:          time.Now(),
	}
	fmt.Println(json.Marshal(offer))
}
