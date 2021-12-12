export type GenreRow = {
	id: string,
	title: string,
	snippet: string,
}

export type BookRow = {
	id: string,
	title: string,
	snippet: string,
	genre: string,
	book_lang: string,
	pages_cnt: number,
	pub_date: string,
}

export type AuthorRow = {
	id: string,
	firstname: string,
	lastname: string,
	surname: string,
	birth_date: string,
	snippet: string,
}