package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	request, err := http.NewRequest("GET", "/articles", nil) //Формируем запрос к нашему API
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()         // Куда пишем
	handler := http.HandleFunc(GetAllArticles) //Какой функционал тестриуем
	handler.ServeHTTP(recorder, request)       // Сопоставляем запрос и функционал

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("HandleFunc returned wrong status code: has %v expected %v", status, http.StatusOK)
	}
}
