<template>
	<div v-if="isBooksLoaded">
		<div v-for="book in books" :key="book.id">
			{{ book }}
		</div>
	</div>
</template>

<script lang="ts" setup>
	import { useStore } from "../store";
	import { storeToRefs } from "pinia";
	const store = useStore();
	const { isBooksLoaded, books } = storeToRefs(store);
</script>

<style>
</style>