<template>
	<div class="editor">
		<div class="editor__header">
			<h2>{{ title }}</h2>
			<button @click="emit('closePressed')"></button>
		</div>
		<div class="editor__body edfields">
			<div class="edfields__field">
				<label class="edfields__label" for="edfirstname">Имя</label>
				<input
					class="edfields__input"
					id="edfirstname"
					v-model="currentAuthor.firstname"
					type="text"
					placeholder="Введите имя автора"
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edlastname">Фамилия</label>
				<input
					class="edfields__input"
					id="edlastname"
					v-model="currentAuthor.firstname"
					type="text"
					placeholder="Введите фамилию автора"
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsurname">Отчество</label>
				<input
					class="edfields__input"
					id="edsurname"
					v-model="currentAuthor.firstname"
					type="text"
					placeholder="Введите отчество автора"
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsnippet">Описание</label>
				<textarea
					class="edfields__input edfields__textarea"
					id="edsnippet"
					v-model="currentAuthor.snippet"
					placeholder="Введите описание автора"
				></textarea>
			</div>
		</div>
		<div class="editor__footer">
			<button @click="saveAuthor">{{ buttonText }}</button>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, reactive } from 'vue';
import { useStore } from "../store"
import { Author } from '../types';
import { AuthorEditorTitle, SaveButtonValue } from '../types/enums';

const props = defineProps({
	author_id: {
		type: String,
	},
})

const emit = defineEmits<{
	(e: "savePressed"): void
	(e: "closePressed"): void
}>()

const store = useStore()
const currentAuthor = reactive<Author>({
	id: "0",
	firstname: "",
	lastname: "",
	surname: "",
	snippet: "",
	deleted: false,
})

const title = ref<AuthorEditorTitle>(AuthorEditorTitle.Create)
const buttonText = ref<SaveButtonValue>(SaveButtonValue.Save)

onBeforeMount(() => {
	if (props.author_id) {
		const author = store.getAuthorByID(props.author_id);
		if (author) {
			title.value = AuthorEditorTitle.Edit
			buttonText.value = SaveButtonValue.Update
		}

		console.log(author);


		currentAuthor.firstname = (author && author.firstname != undefined) ? author?.firstname : ""
		currentAuthor.lastname = (author && author.lastname != undefined) ? author?.lastname : ""
		currentAuthor.surname = (author && author.surname != undefined) ? author?.surname : ""
		currentAuthor.snippet = (author && author.snippet != undefined) ? author?.snippet : ""
	}
})

const saveAuthor = () => {
	setTimeout(() => {
		emit('savePressed')
	}, 500)
}

</script>

<style>
</style>