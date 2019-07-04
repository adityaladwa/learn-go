package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"http://some.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"https://google.com":      true,
		"http://some.com":         true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebSiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
