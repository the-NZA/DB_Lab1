import { defineStore } from "pinia";
import { Book, Genre, Author } from "../types";


export const useStore = defineStore("main", {
	state: () => ({
		books: [] as Array<Book>,
		genres: [] as Array<Genre>,
		authors: [] as Array<Author>,
	}),
	getters: {
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
			const res = null;
		}
	}
})