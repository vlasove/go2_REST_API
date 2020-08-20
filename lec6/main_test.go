package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//3 . Создал ВСЕ МОДУЛЬНЫЕ ТЕСТЫ
//4. Прописал тестовый сценарий в Postman
func TestGetItems(t *testing.T) {
	request, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetItems)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v, expected %v", status, http.StatusOK)
	}

	//Возвращает не нулевуой ответ
	if len(recorder.Body.String()) <= 1 {
		t.Errorf("Handler returned %v", recorder.Body.String())
	}
}

func TestGetItemID(t *testing.T) {
	request, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetItemID)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v, expected %v", status, http.StatusOK)
	}

	if len(recorder.Body.String()) <= 1 {
		t.Errorf("Handler returned %v", recorder.Body.String())
	}
}

func TestGetItemIDDoesntExists(t *testing.T) {
	request, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("id", "123")
	request.URL.RawQuery = q.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetItemID)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned %v, expected %v", status, http.StatusNotFound)
	}

}

func TestPostItem(t *testing.T) {
	var jsonFIle = []byte(`{"id":"2","content":"dfge"}`)
	request, err := http.NewRequest("POST", "/item", bytes.NewBuffer(jsonFIle))

	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PostItem)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong %v code", status)
	}
}

func TestDeleteItemIDDoesntExist(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/item", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("id", "200")
	request.URL.RawQuery = q.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteItemID)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestPutItemID(t *testing.T) {
	var jsonFIle = []byte(`{"id":"1","content":"dfge"}`)

	request, err := http.NewRequest("PUT", "/item", bytes.NewBuffer(jsonFIle))

	q := request.URL.Query()

	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PutItemID)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusAccepted {
		t.Errorf("Handler returned wrong %v code", status)
	}

}

func TestPutItemIDdOESNTeXIST(t *testing.T) {
	var jsonFIle = []byte(`{"id":"200","content":"dfge"}`)

	request, err := http.NewRequest("PUT", "/item", bytes.NewBuffer(jsonFIle))

	q := request.URL.Query()
	q.Add("id", "200")
	request.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PutItemID)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong %v code", status)
	}

}

func TestDeleteItemID(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/item", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteItemID)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}
