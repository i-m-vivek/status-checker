package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostWebsiteHandler(t *testing.T) {
	reqBody := []byte(`{
  "websites": ["https://www.youtube.com", "https://www.google.com", "https://www.facebook.com", "https:www.fakewebsite123.com"]
}`)

	req := httptest.NewRequest(http.MethodPost, "/POST/websites", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	PostWebsiteHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
	data := make(map[string][]string)
	err = json.Unmarshal(dataByte, &data)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	wList := []string{"https://www.youtube.com", "https://www.google.com", "https://www.facebook.com", "https:www.fakewebsite123.com"}
	expectedRes := map[string][]string{"websites": wList}

	if !compareMapOfList(data, expectedRes) {
		t.Errorf("expected %v, got %v", expectedRes, data)
	}
}

func TestGetWebsiteHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/GET/websites", nil)
	w := httptest.NewRecorder()
	GetWebsiteHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
	data := make(map[string]string)
	err = json.Unmarshal(dataByte, &data)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	expectedRes := map[string]string{
		"https://www.facebook.com":     "UP",
		"https://www.google.com":       "UP",
		"https://www.youtube.com":      "UP",
		"https:www.fakewebsite123.com": "DOWN",
	}

	if !compareMapOfString(expectedRes, data) {
		t.Errorf("expected %v, got %v", expectedRes, data)
	}
}
func compareMapOfList(m1, m2 map[string][]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, _ := range m1 {
		if len(m1[key]) != len(m2[key]) {
			return false
		}
		for i, _ := range m1[key] {
			if m1[key][i] != m2[key][i] {
				return false
			}
		}
	}
	return true
}

func compareMapOfString(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, _ := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}
	return true
}
