<template>
	<template v-if="isAuthorsLoaded">
		<div v-for="author in authors" :key="author.id">
			{{ author }}
		</div>
	</template>
</template>

<script lang="ts" setup>
	import { useStore } from "../store";
	import { storeToRefs } from "pinia";
	const store = useStore();
	const { isAuthorsLoaded, authors } = storeToRefs(store);
</script>

<style>
</style>