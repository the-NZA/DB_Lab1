CREATE TABLE IF NOT EXISTS books (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  title text NOT NULL,
  snippet text,
  pages_cnt integer DEFAULT(0),
  pub_date timestamp,
  deleted boolean NOT NULL DEFAULT(false),
  genre_id integer NOT NULL,
  book_lang_id integer NOT NULL,
  FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY (book_lang_id) REFERENCES book_lang (id) ON DELETE NO ACTION ON UPDATE NO ACTION
) COMMENT 'Информация о книгах';