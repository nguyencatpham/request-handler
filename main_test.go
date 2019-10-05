package requestHandler_test

import (
	"encoding/json"
	"fmt"
	"testing"

	requestHandler "github.com/nguyencatpham/request-handler"
)

type TestData struct {
	ID   int
	Name string
}

func TestRequestHandle(t *testing.T) {

	RootURL := "http://5916858b5e17bd0011f6b941.mockapi.io/api/request-handler"

	model := TestData{
		ID:   1,
		Name: "Test",
	}
	b, _ := json.Marshal(model)

	requestModel := requestHandler.RequestModel{
		URL:       RootURL,
		TokenType: requestHandler.Basic,
		Username:  "",
		Password:  "",
		Body:      string(b),
	}

	testGet(requestModel, t)
	testPost(requestModel, t)
	testPut(requestModel, t)
	testPatch(requestModel, t)
	testDelete(requestModel, t)
}
func testGet(model requestHandler.RequestModel, t *testing.T) {
	var result interface{}

	if err := requestHandler.Get(model, &result); err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil response")
	}
}
func testPost(model requestHandler.RequestModel, t *testing.T) {
	var result interface{}

	if err := requestHandler.Post(model, &result); err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil response")
	}
}
func testPut(model requestHandler.RequestModel, t *testing.T) {
	var result interface{}

	model.URL = fmt.Sprintf("%s/%s", model.URL, "1")
	if err := requestHandler.Put(model, &result); err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil response")
	}
}
func testPatch(model requestHandler.RequestModel, t *testing.T) {
	var result interface{}

	model.URL = fmt.Sprintf("%s/%s", model.URL, "1")
	if err := requestHandler.Patch(model, &result); err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil response")
	}
}
func testDelete(model requestHandler.RequestModel, t *testing.T) {
	var result interface{}

	model.URL = fmt.Sprintf("%s/%s", model.URL, "1")
	if err := requestHandler.Delete(model, &result); err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil response")
	}
}
