<template>
	<span class="grid-buttons">
		<button @click="editFunc" class="grid-button editBtn">
			<img src="../assets/edit.svg" alt="Edit" />
		</button>
		<button @click="deleteFunc" class="grid-button deleteBtn">
			<img src="../assets/delete.svg" alt="Delete" />
		</button>
	</span>
</template>

<script lang="ts">
import { ICellRendererParams } from "@ag-grid-community/all-modules";

export default {
	name: "DeleteButton",
	setup(props: any) {
		const params: ICellRendererParams = props.params;

		const deleteFunc = () => {
			params.context.deleteFunc(params.data.id);
		};

		const editFunc = () => {
			params.context.editFunc(params.data.id);
		}

		return { editFunc, deleteFunc };
	},
};
</script>

<style>
.grid-buttons {
	display: flex;
	min-height: 40px;
	align-items: center;
	justify-content: space-between;
}

.grid-button {
	margin-right: var(--offset-half);

	appearance: none;
	cursor: pointer;

	background-color: rgb(var(--leanen));
	border: 1px solid rgb(var(--redsand));
	border-radius: var(--offset-quarter);

	transition: var(--main-transition);
}

.editBtn:hover {
	background-color: rgb(var(--white));
}

.deleteBtn:hover {
	background-color: rgb(var(--redsand));
}

.grid-button img {
	display: block;
	box-sizing: border-box;

	width: var(--offset-half);
	height: calc(var(--offset-half) * 1.5);
}

.grid-button:last-child {
	margin-right: 0;
}
</style>