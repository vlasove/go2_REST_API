package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//Тесты для непараметризованного GET запроса
func TestGetAllArticles(t *testing.T) {
	request, err := http.NewRequest("GET", "/articles", nil) //Формируем запрос к нашему API
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()          // Куда пишем
	handler := http.HandlerFunc(GetAllArticles) //Какой функционал тестриуем
	handler.ServeHTTP(recorder, request)        // Сопоставляем запрос и функционал
	//Тест сопоставления статус кода
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("HandlerFunc returned wrong status code: has %v expected %v", status, http.StatusOK)
	}
	//Тест сопоставления ответа запроса
	expectedAnswer := `[{"Id":"1","Title":"First title","Author":"First author","Content":"First content"},{"Id":"3","Title":"Updated title PUT","Author":"Updated author PUT","Content":"Updated content PUT"}]`
	if recorder.Body.String() != expectedAnswer {
		t.Errorf("HandlerFunc returned wrong answer: has %v expected %v", recorder.Body.String(), expectedAnswer)
	}
}

//Тест для параметризованного GET запроса
func TestGetArticleWithId(t *testing.T) {
	request, err := http.NewRequest("GET", "/article", nil)
	if err != nil {
		t.Fatal(err)
	}
	query := request.URL.Query()          // Доп очередь параметров к базовому запросу
	query.Add("id", "1")                  // Добавляю в очередь параметр id со значением 1
	request.URL.RawQuery = query.Encode() // Приклеиваем параметры следом за запросом /article + /1

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetArticleWithId)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("HandlerFunc returned wrong status code: has %v expected %v", status, http.StatusOK)
	}
	expectedAnswer := `[{"Id":"1","Title":"First title","Author":"First author","Content":"First content"}]`
	if recorder.Body.String() != expectedAnswer {
		t.Errorf("HandlerFunc returned wrong answer: has %v expected %v", recorder.Body.String(), expectedAnswer)
	}
}
