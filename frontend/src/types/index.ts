export type Book = {
	id: string,
	title: string,
	snippet: string,
	genre_id: string,
	book_lang_id: string,
	pages_cnt: number,
	pub_date: Date,
	deleted: boolean,
}

export type Genre = {
	id: string,
	title: string,
	snippet: string,
	deleted: boolean,
}

export type Author = {
	id: string,
	firstname: string,
	lastname: string,
	surname: string,
	birth_date: Date,
	snippet: string,
	deleted: boolean,
}