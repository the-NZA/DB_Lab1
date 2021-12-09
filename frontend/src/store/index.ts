import { defineStore } from "pinia";
import { Book, Genre, Author } from "../types";
import { GET } from "../HTTP"

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
		}
	},
	actions: {
		async loadBooks() {
			try {
				const allBooks = await GET<Book[]>("api/book/all");
				this.books = allBooks;
				this.booksLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
		async loadGenres() {
			try {
				const allGenres = await GET<Genre[]>("api/genre/all");
				this.genres = allGenres;
				this.genresLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
		async loadAuthors() {
			try {
				const allAuthors = await GET<Author[]>("api/author/all");
				this.authors = allAuthors;
				this.authorsLoaded = true
			} catch (error) {
				console.error(error);
			}
		},
	}
})