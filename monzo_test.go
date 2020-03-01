package myzo

import (
	"encoding/json"
	"testing"
)

func TestAuthMonzo_GetAllMerchants(t *testing.T) {
	auth := Myzo{ClientID: "", UserID: "", AccessToken: "", AccountID: "", Debug:true}
	body := auth.ResponseBody
	var m Merchant
	if err := json.Unmarshal(body, &m); err != nil {
		t.Error("Could not unmarshal response: " + err.Error())
	} else {
		if m.ID != "" {
			t.Log("Passed test, found result: " + m.ID)
		} else {
			t.Error("FAILED! Could not find the ID.")
		}
	}
}
