## Project_for_GO


# Запуск:
make app_run


# Тестирование
Для того чтобы работать с сервисом обязательно нужно быть залогиненым, время жизни сессионного токена 1 минута
Сервис принимает curl запросы, его можно протестировать следующими запросаи:

1) Регистрация и залогинивание:
curl -X POST -H "Content-Type: application/json" -d '{"login": "admin", "password": "secret"}' -c cookies.txt http://localhost:8080/register

curl -X POST -H "Content-Type: application/json" -d '{"login": "admin", "password": "secret"}' -c cookies.txt http://localhost:8080/login

2) Добавление и проверка добавленной записи:
curl -X POST -H "Content-Type: application/json" -d '{"title":"Понедельник","text":"Сегодня был чудестный день, на улице теплая погода"}' -b cookies.txt http://localhost:8080/add

curl -X GET -H "Content-Type: application/json" -d '{"title": "Понедельник"}' -b cookies.txt http://localhost:8080/my_note

3) Изменение заголовка/текста:
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Понедельник","text":"Сегодня было холодно, пошел снег..."}' -b cookies.txt http://localhost:8080/change_text

curl -X PUT -H "Content-Type: application/json" -d '{"old_title":"Понедельник","new_title":"Вторник"}' -b cookies.txt http://localhost:8080/change_title

4) Удаление заметки:
curl -X DELETE -H "Content-Type: application/json" -d '{"title": "Понедельник"}' -b cookies.txt http://localhost:8080/delete 

5) Изменение пароля:
curl -X PUT -H "Content-Type: application/json" -d '{"login": "admin", "old_password": "secret", "new_password": "bobr"}' -b cookies.txt http://localhost:8080/change_password