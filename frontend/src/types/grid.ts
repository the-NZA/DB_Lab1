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
	pages_cnt: number,
	pub_year: number,
}

export type AuthorRow = {
	id: string,
	firstname: string,
	lastname: string,
	surname: string,
	snippet: string,
}