# Лабораторная работа 1 по дисциплине "Программирование и администрирование в среде СУБД"

## Задание
Разработать веб-приложение, для управления базой данных MySQL.
Реализовать стандартные модули вывода информации на основе хранимых процедур (3 процедуры).

### Сущности
- Жанры
- Книги
- Авторы

### Отношения
- Жанры 1:M Книги
- Книги M:M Авторы

### Таблицы
1. Жанры
	- id
	- title
	- snippet
	- deleted
2. Книги
	- id
	- title
	- snippet
	- genre_id
	- pages_cnt
	- pub_year
	- book_lang_id
	- deleted
3. Авторы
	- id
	- firstname
	- lastname
	- surname
	- snippet
	- deleted
4. Языки книг
	- id
	- title
5. Книги и авторы
	- id
	- book_id
	- author_id
	- deleted
