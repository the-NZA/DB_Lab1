<template>
	<div class="editor">
		<div class="editor__header">
			<h2>{{ title }}</h2>
			<button @click="emit('closePressed')"></button>
		</div>
		<div class="editor__body edfields">
			<div class="edfields__field">
				<label class="edfields__label" for="edtitle">Название</label>
				<input
					class="edfields__input"
					id="edtitle"
					v-model="currentGenre.title"
					type="text"
					placeholder="Введите название жанра"
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsnippet">Описание</label>
				<textarea
					class="edfields__input edfields__textarea"
					id="edsnippet"
					v-model="currentGenre.snippet"
					placeholder="Введите описание жанра"
				></textarea>
			</div>
		</div>
		<div class="editor__footer">
			<button @click="saveGenre">{{ buttonText }}</button>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, reactive } from 'vue';
import { useStore } from "../store"
import { Genre } from '../types';
import { GenreEditorTitle, SaveButtonValue } from '../types/enums';

const props = defineProps({
	genre_id: {
		type: String,
	},
})

const emit = defineEmits<{
	(e: "savePressed"): void
	(e: "closePressed"): void
}>()

const store = useStore()
const currentGenre = reactive<Genre>({
	id: "0",
	title: "",
	snippet: "",
	deleted: false,
})

const title = ref<GenreEditorTitle>(GenreEditorTitle.Create)
const buttonText = ref<SaveButtonValue>(SaveButtonValue.Save)

onBeforeMount(() => {
	if (props.genre_id) {
		const genre = store.getGenreByID(props.genre_id);
		if (genre) {
			title.value = GenreEditorTitle.Edit
			buttonText.value = SaveButtonValue.Update
		}

		currentGenre.title = (genre && genre.title != undefined) ? genre?.title : ""
		currentGenre.snippet = (genre && genre.snippet != undefined) ? genre?.snippet : ""

	}
})

const saveGenre = () => {
	setTimeout(() => {
		emit('savePressed')
	}, 500)
}

</script>

<style>
</style>