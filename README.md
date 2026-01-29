# GreenAPI клиент для отправки сообщений

## Технологии

**Backend** 

- Golang 1.25
- chi-router
- Task - для запуска, сборки, линтера/форматтера 

**Frontend**

- NodeJS 24
- ReactJS (без TypeScript)
- axios
- Tailwind

Монорепозиторий для работы с GREEN-API с функционалом:

- Получение информации об инстансе
- Отправка текстового сообщения
- Отправка сообщения с файлом по ссылке

## Структура проекта

Монорепозиторий фронтенда и бэкенда

## Запуск для разработки

### Backend

Предварительно заполнить .env файл в соответствии с примером `backend/example.env`

```bash
cd backend
task run
```


### Frontend

Предварительно заполнить .env-файл в соответствии с примером `frontend/example.env`

```bash
cd frontend
npm install
npm run start
```
