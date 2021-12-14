<template>
	<div class="genreeditor">
		<input v-model="currentGenre.snippet" type="text" />
		{{ currentGenre.snippet }}
	</div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, reactive } from 'vue';
import { useStore } from "../store"
import { Genre } from '../types';

const props = defineProps({
	genre_id: {
		type: String,
	}
})

const store = useStore()
const currentGenre = reactive<Genre>({
	id: "0",
	title: "",
	snippet: "",
	deleted: false,
})

onBeforeMount(() => {
	if (props.genre_id) {
		const genre = store.getGenreByID(props.genre_id);
		currentGenre.title = (genre && genre.title != undefined) ? genre?.title : ""
		currentGenre.snippet = (genre && genre.snippet != undefined) ? genre?.snippet : ""
	}
})


</script>

<style>
/* .genreeditor {
	position: absolute;
	top: 0;
	bottom: 0;
	left: 0;
	right: 0;
	background-color: rgba(var(--black), 0.3);
} */
</style>