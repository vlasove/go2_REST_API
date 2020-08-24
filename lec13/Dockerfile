# sudo apt-get docker docker-compose 

# Базовый образ го
FROM golang:alpine as builder 

#ENV GO111MODULE=on

LABEL maintainer="GoCourse <restapi@specialist.ru>"

RUN apk update && apk add --no-cache git 

# Создаем рабочую директорию
WORKDIR /app

# Скопируем все наши зависимости в контейнер (только для версий go1.13 >)
COPY go.mod go.sum ./

RUN go mod download 

COPY . .

# Сборка приложения (без учета рантайма GoLang)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . 



#Настройка внутреннего порта контейнера
EXPOSE 8080 

# Запуск скомпилированного проекта
CMD [ "./main" ]