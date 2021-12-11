<template>
	<button @click="clicked" class="editBtn">Edit</button>
</template>

<script lang="ts">
	import { ICellRendererParams } from "@ag-grid-community/all-modules";

	export default {
		name: "EditButton",
		setup(props: any) {
			const clicked = () => {
				const params: ICellRendererParams = props.params;
				params.context.editFunc(params.data.id);
			};

			return { clicked };
		},
	};
</script>

<style>
</style>