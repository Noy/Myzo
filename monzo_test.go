package myzo

import (
	"encoding/json"
	"strings"
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

func TestMyzo_FeedHandler(t *testing.T) {
	auth := Myzo{ClientID:"", UserID:"", AccessToken:"", AccountID:"", Debug:true, ResponseBody: nil}
	resp, err := auth.FeedHandler("https://google.com", [6]string{
		"This will be the notification on your app!",
		"Link to url",
		"#000000",
		"#000000",
		"#000000",
		"Body text",
	})
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(resp), "bad") {
		t.Error("Did not do it correctly, check all is in order: " + string(resp))
	}
}