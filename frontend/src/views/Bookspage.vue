<template>
	<template v-if="isBooksLoaded">
		<div v-for="book in books" :key="book.id">
			{{ book }}
		</div>
	</template>
</template>

<script lang="ts" setup>
	import { ref, reactive } from "vue";
	import { useStore } from "../store";
	import { storeToRefs } from "pinia";

	const store = useStore();
	const { isBooksLoaded, books } = storeToRefs(store);
</script>

<style>
</style>