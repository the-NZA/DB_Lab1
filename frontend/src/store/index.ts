import { defineStore } from "pinia";
import { Book, Genre, Author } from "../types";
import { GET, DELETE } from "../HTTP"
import { AuthorRow, BookRow, GenreRow } from "../types/grid";
import { formatDate } from "../utils/date";

export const useStore = defineStore("main", {
	state: () => ({
		books: [] as Array<Book>,
		genres: [] as Array<Genre>,
		authors: [] as Array<Author>,
		booksLoaded: false,
		genresLoaded: false,
		authorsLoaded: false,
	}),
	getters: {
		isBooksLoaded(): boolean {
			return this.booksLoaded
		},
		isGenresLoaded(): boolean {
			return this.genresLoaded
		},
		isAuthorsLoaded(): boolean {
			return this.authorsLoaded
		},
		getBooks(): Book[] {
			return this.books
		},
		getBookByID: (state) => {
			return (id: string): Book | undefined => {
				return state.books.find(item => item.id === id)
			}
		},
		getAuthors(): Author[] {
			return this.authors
		},
		getAuthorByID: (state) => {
			return (id: string): Author | undefined => {
				return state.authors.find(item => item.id === id)
			}
		},
		getGenres(): Genre[] {
			return this.genres
		},
		getGenresRows(): GenreRow[] {
			return this.genres.map((genre): GenreRow => {
				return {
					id: genre.id,
					title: genre.title,
					snippet: genre.snippet,
				}
			})
		},
		getBooksRows(): BookRow[] {
			return this.books.map((book): BookRow => {
				return {
					id: book.id,
					title: book.title,
					snippet: book.snippet,
					pages_cnt: book.pages_cnt,
					pub_date: formatDate(new Date), // ! FIX THIS
					book_lang: "Русский",
					genre: "Роман"
				}
			})
		},
		getAuthorsRows(): AuthorRow[] {
			return this.authors.map((author): AuthorRow => {
				return {
					id: author.id,
					firstname: author.firstname,
					lastname: author.lastname,
					surname: author.surname,
					snippet: author.snippet,
					birth_date: formatDate(new Date), // ! FIX THIS
				}
			})
		}
	},
	actions: {
		/* BOOKS */
		async loadBooks() {
			try {
				const allBooks = await GET<Book[]>("api/book/all");
				this.books = allBooks;
				this.booksLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
		async deleteBook(id: string) {
			try {
				const resp = await DELETE(`/api/book/${id}`)
				if (!resp.ok) {
					console.log(resp);

					throw new Error(resp.statusText)

				}

				const res = resp.json()
				// console.log(res);

				this.books= this.books.filter((book): boolean => {
					return book.id != id;
				})

			} catch (err) {
				console.error(err);
			}
		},

		/* GENRES */
		async loadGenres() {
			try {
				const allGenres = await GET<Genre[]>("api/genre/all");
				this.genres = allGenres;
				this.genresLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
		addGenre(): void {
			this.genres.push({
				id: "23",
				deleted: false,
				snippet: "werqwer",
				title: "test genre",
			})
		},
		async deleteGenre(id: string) {
			try {
				const resp = await DELETE(`/api/genre/${id}`)
				if (!resp.ok) {
					console.log(resp);

					throw new Error(resp.statusText)

				}

				const res = resp.json()
				// console.log(res);

				this.genres= this.genres.filter((genre): boolean => {
					return genre.id != id;
				})

			} catch (err) {
				console.error(err);
			}
		},

		/* AUTHORS */
		async loadAuthors() {
			try {
				const allAuthors = await GET<Author[]>("api/author/all");
				this.authors = allAuthors;
				this.authorsLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
		async deleteAuthor(id: string) {
			try {
				const resp = await DELETE(`/api/author/${id}`)
				if (!resp.ok) {
					console.log(resp);

					throw new Error(resp.statusText)

				}

				const res = resp.json()
				// console.log(res);

				this.authors = this.authors.filter((author): boolean => {
					return author.id != id;
				})

			} catch (err) {
				console.error(err);
			}
		},
	}
})