<template>
	<!-- <template v-if="isAuthorsLoaded">
		<div v-for="author in authors" :key="author.id">
			{{ author }}
		</div>
	</template> -->
	<button @click="getSelectedRows">Get Selected Rows</button>
	<button @click="clearSelection">Clear selection</button>
	<button @click="addNewRow">add new row</button>

	<ag-grid-vue
		style="width: 100%; height: 100%"
		class="ag-theme-alpine"
		:gridOptions="gridOptions"
		:context="actionContext"
	></ag-grid-vue>
</template>

<script lang="ts" setup>
	import { ref, reactive, onBeforeMount, getCurrentInstance } from "vue";
	import { useStore } from "../store";
	import { storeToRefs } from "pinia";
	import DeleteButton from "../components/DeleteButton.vue";
	import EditButton from "../components/EditButton.vue";
	import { AgGridVue } from "ag-grid-vue3";
	import {
		GridReadyEvent,
		GridApi,
		GridOptions,
		ColumnApi,
		RowSelectedEvent,
	} from "@ag-grid-community/all-modules";

	const store = useStore();
	// Reactive vars from store
	const { isAuthorsLoaded, authors } = storeToRefs(store);

	// Context for action buttons
	const actionContext = ref(null);
	const gridApi = ref<GridApi>();
	const columnApi = ref<ColumnApi>();

	const rowData = reactive([
		{
			id: "1",
			firstname: "Ivan",
			lastname: "Ivanov",
			surname: "Ivanovich",
			birth_date: new Date().toDateString(),
			snippet: "This is Ivan Ivanov aka Triple I",
		},
		{
			id: "2",
			firstname: "Petr",
			lastname: "Petrov",
			surname: "Petrovich",
			birth_date: new Date(1999, 8, 31).toDateString(),
			snippet: "This is Petr Petrov aka Piper",
		},
	]);

	const addNewRow = () => {
		rowData.push({
			id: Math.random().toString(),
			firstname: "New firstnaem",
			lastname: "New lastnaem",
			surname: "New surnaem",
			birth_date: new Date().toDateString(),
			snippet: "default snippet",
		});
		gridApi.value.setRowData(rowData);
	};

	const getSelectedRows = () => {
		const selected = gridApi.value.getSelectedRows();
		console.log(selected);
	};

	const clearSelection = () => {
		gridApi.value.deselectAll();
	};

	const onRowSelected = (e: RowSelectedEvent) => {
		// console.log(e);
	};

	// Save grid and column API on grid ready event
	const onGridReady = (params: GridReadyEvent) => {
		gridApi.value = params.api;
		columnApi.value = params.columnApi;
	};

	const gridOptions = ref<GridOptions>({
		onRowSelected: onRowSelected,
		onGridReady: onGridReady,
		rowSelection: "multiple",
		suppressCellSelection: true,
		suppressRowClickSelection: true,
		columnDefs: [
			{
				field: "id",
				hide: true,
			},
			{
				field: "firstname",
				checkboxSelection: true,
				headerCheckboxSelection: true,
				headerName: "Имя",
			},
			{
				field: "lastname",
				headerName: "Фамилия",
			},
			{
				field: "surname",
				headerName: "Отчество",
			},
			{
				field: "birth_date",
				headerName: "Дата рождения",
			},
			{
				field: "snippet",
				headerName: "Описание",
			},
			{
				field: "edit",
				headerName: "Редактировать",
				cellRenderer: "editBtn",
			},
			{
				field: "delete",
				headerName: "Удалить",
				cellRenderer: "deleteBtn",
			},
		],
		rowData: rowData,
		frameworkComponents: {
			deleteBtn: DeleteButton,
			editBtn: EditButton,
		},
	});

	// Actions funcs which passed to ag-grid
	// and called after button pressed
	const deleteFunc = (id) => {
		alert(`from delete ${id}`);
	};
	const editFunc = (id) => {
		alert(`from edit ${id}`);
	};

	onBeforeMount(() => {
		// Setup actions context
		actionContext.value = {
			deleteFunc: deleteFunc,
			editFunc: editFunc,
		};
	});
</script>

<style lang="postcss">
.ag-cell-wrapper .ag-selection-checkbox {
	margin-right: 18px !important;
}
</style>