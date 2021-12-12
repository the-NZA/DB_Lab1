CREATE TABLE IF NOT EXISTS book_lang (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  title text UNIQUE NOT NULL
) COMMENT "Словарь с языками книг";