CREATE TABLE IF NOT EXISTS authors (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  firstname text NOT NULL,
  lastname text NOT NULL,
  surname text,
  birth_date timestamp,
  snippet text,
  deleted boolean NOT NULL DEFAULT(false)
) COMMENT "Информация об авторах";