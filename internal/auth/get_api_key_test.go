package auth

import (
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey myapikey")
	got, err := GetAPIKey(header)
	if err != nil {
		log.Printf("Something went wrong try again")
	}
	want := "myapikey"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected %v, got: %v", want, got)
	}
}
