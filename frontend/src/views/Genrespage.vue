<template>
	<template v-if="isGenresLoaded">
		<div v-for="genre in genres" :key="genre.id">
			{{ genre }}
		</div>
	</template>
</template>

<script lang="ts" setup>
	import { useStore } from "../store";
	import { storeToRefs } from "pinia";
	const store = useStore();
	const { isGenresLoaded, genres } = storeToRefs(store);
</script>

<style>
</style>