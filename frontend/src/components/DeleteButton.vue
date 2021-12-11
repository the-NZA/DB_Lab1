<template>
	<button @click="clicked" class="deleteBtn">Delete</button>
</template>

<script lang="ts">
	import { ICellRendererParams } from "@ag-grid-community/all-modules";

	export default {
		name: "DeleteButton",
		setup(props: any) {
			const clicked = () => {
				const params: ICellRendererParams = props.params;
				params.context.deleteFunc(params.data.id);
			};

			return { clicked };
		},
	};
</script>

<style>
</style>