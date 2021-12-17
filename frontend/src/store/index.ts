import { defineStore } from "pinia";
import { Book, Genre, Author, BookAuthor } from "../types";
import { GET, POST, PUT, DELETE } from "../HTTP"
import { AuthorRow, BookRow, GenreRow } from "../types/grid";

export const useStore = defineStore("main", {
	state: () => ({
		books: [] as Array<Book>,
		genres: [] as Array<Genre>,
		authors: [] as Array<Author>,
		booksAuthors: [] as Array<BookAuthor>,
		booksLoaded: false,
		genresLoaded: false,
		authorsLoaded: false,
		booksAuthorsLoaded: false,
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
			return (id?: string): Author | undefined => {
				if (!id) {
					return undefined
				}

				return state.authors.find(item => item.id === id)
			}
		},
		getGenres(): Genre[] {
			return this.genres
		},
		getGenreByID: (state) => {
			return (id?: string): Genre | undefined => {
				if (!id) {
					return undefined
				}

				return state.genres.find(item => item.id === id)
			}
		},
		getAuthorsByBookID: (state) => {
			return (book_id: string) => {
				const booksAuthorsIDs = state.booksAuthors.filter(ba => ba.book_id === book_id)
				const authors: string[] = []

				for (const ba of booksAuthorsIDs) {
					const author = state.authors.find(item => item.id === ba.author_id)
					if (author) {
						let formatedAuthor = author.surname.length === 0 ? `${author.lastname} ${author.firstname[0].toUpperCase()}.` : `${author.lastname} ${author.firstname[0].toUpperCase()}. ${author.surname[0].toUpperCase()}.`

						authors.push(formatedAuthor.charAt(0).toUpperCase() + formatedAuthor.slice(1))
					}
				}

				// authors.push("Kozlov R.K.")
				// authors.push("Kozlov R.K.")

				return authors.join(", ")
			}
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
					pub_year: book.pub_year,
					genre: this.getGenreByID(book.genre_id)!.title,
					authors: this.getAuthorsByBookID(book.id),
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
		async addBook(book: Book) {
			try {
				const res = await POST<Book>(book, "/api/book")
				this.books.push(res)
			}
			catch (err) {
				console.error(err);
			}
		},
		async updateBook(updated_book: Book) {
			try {
				const res = await PUT<Book>(updated_book, "/api/book")
				const idx = this.books.findIndex(book => book.id === res.id)

				// Update store through slice with spreads
				this.books = [...this.books.slice(0, idx), res, ...this.books.slice(idx + 1)]
			}
			catch (err) {
				console.error(err);
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

				this.books = this.books.filter((book): boolean => book.id != id)

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
		async addGenre(genre: Genre) {
			try {
				const res = await POST<Genre>(genre, "/api/genre")
				this.genres.push(res)
			}
			catch (err) {
				console.error(err);
			}
		},
		async updateGenre(updated_genre: Genre) {
			try {
				const res = await PUT<Genre>(updated_genre, "/api/genre")
				const idx = this.genres.findIndex(genre => genre.id === res.id)

				// Update store through slice with spreads
				this.genres = [...this.genres.slice(0, idx), res, ...this.genres.slice(idx + 1)]
			}
			catch (err) {
				console.error(err);
			}
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

				this.genres = this.genres.filter((genre): boolean => genre.id != id)

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
		async addAuthor(author: Author) {
			try {
				const res = await POST<Author>(author, "/api/author")
				this.authors.push(res)
			}
			catch (err) {
				console.error(err);
			}
		},
		async updateAuthor(updated_author: Author) {
			try {
				const res = await PUT<Author>(updated_author, "/api/author")
				const idx = this.authors.findIndex(author => author.id === res.id)

				// Update store through slice with spreads
				this.authors = [...this.authors.slice(0, idx), res, ...this.authors.slice(idx + 1)]
			}
			catch (err) {
				console.error(err);
			}
		},
		async deleteAuthor(id: string) {
			try {
				const resp = await DELETE(`/api/author/${id}`)
				console.log(resp);

				if (!resp.ok) {
					console.log(resp);

					throw new Error(resp.statusText)

				}

				const res = resp.json()
				// console.log(res);

				this.authors = this.authors.filter((author): boolean => author.id != id)

			} catch (err) {
				console.error(err);
			}
		},

		/* BOOKS AUTHORS */
		async loadBooksAuthors() {
			try {
				const allBooksAuthors = await GET<BookAuthor[]>("api/book-author");

				this.booksAuthors = allBooksAuthors;
				this.booksAuthorsLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
	}
})