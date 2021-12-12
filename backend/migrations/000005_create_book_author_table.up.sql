CREATE TABLE IF NOT EXISTS book_author (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  book_id integer NOT NULL,
  author_id integer NOT NULL,
  deleted boolean NOT NULL DEFAULT(false),
  FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE NO ACTION ON UPDATE NO ACTION
) COMMENT "Отношение книг и авторов";