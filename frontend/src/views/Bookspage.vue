<template>
	<actions-buttons
		:canEdit="singleSelected"
		:canDelete="hasSelected"
		@add-pressed="handleAdd"
		@edit-pressed="handleEdit"
		@delete-pressed="handleDelete"
	></actions-buttons>

	<ag-grid-vue
		style="width: 100%; height: 100%;"
		class="ag-theme-alpine"
		:gridOptions="gridOptions"
		:context="actionContext"
		:rowData="getBooksRows"
	></ag-grid-vue>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount, computed } from "vue";
import { useStore } from "../store";
import { storeToRefs } from "pinia";
import GridButtonsVue from "../components/GridButtons.vue";
import ActionsButtons from "../components/ActionsButtons.vue"
import { AgGridVue } from "ag-grid-vue3";
import {
	GridReadyEvent,
	GridApi,
	GridOptions,
	ColumnApi,
	SelectionChangedEvent,
} from "@ag-grid-community/all-modules";
import { BookRow } from "../types/grid";

const store = useStore();
const { isBooksLoaded, getBooksRows } = storeToRefs(store);

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const handleAdd = () => {
	console.log("add was pressed");
}

const handleEdit = () => {
	const selected = gridApi.value?.getSelectedRows()
	if (selected?.length !== 1) {
		console.log("Выбрано более одного элемента. Редактирование недоступно.");
		return;
	}

	console.log("edit was pressed");
}

const handleDelete = () => {
	const selected = gridApi.value?.getSelectedRows()

	if (confirm("Вы уверены, что хотите удалить выбранные книги?")) {
		selected?.forEach((book: BookRow) => {
			store.deleteBook(book.id);
		})
	}
}

const onSelectionChanged = (e: SelectionChangedEvent) => {
	const cnt = e.api.getSelectedRows().length
	switch (cnt) {
		case 0:
			hasSelected.value = false;
			singleSelected.value = false;
			break;
		case 1:
			hasSelected.value = true;
			singleSelected.value = true;
			break;
		default:
			hasSelected.value = true;
			singleSelected.value = false;
			break;

	}
}

// Save grid and column API on grid ready event
const onGridReady = (params: GridReadyEvent) => {
	gridApi.value = params.api;
	columnApi.value = params.columnApi;
};

const gridOptions = ref<GridOptions>({
	onSelectionChanged: onSelectionChanged,
	onGridReady: onGridReady,
	rowSelection: "multiple",
	suppressCellSelection: true,
	suppressRowClickSelection: true,
	defaultColDef: {
		wrapText: true,
		autoHeight: true,
	},
	columnDefs: [
		{
			field: "id",
			hide: true,
		},
		{
			field: "title",
			checkboxSelection: true,
			headerCheckboxSelection: true,
			headerName: "Название",
			flex: 1,
			minWidth: 250,
		},
		{
			field: "snippet",
			headerName: "Описание",
			flex: 2,
			minWidth: 450,
		},
		{
			field: "genre",
			headerName: "Жанр",
			minWidth: 120,
			maxWidth: 140,
		},
		{
			field: "pages_cnt",
			headerName: "Страниц",
			width: 100,
		},
		{
			field: "pub_year",
			headerName: "Год публикации",
			maxWidth: 155,
		},
		{
			field: "actions",
			headerName: "Действия",
			cellRenderer: "gridBtn",
			// cellClass: ["test", "test2"], // Use for grid action buttons
			flex: 0.5,
			minWidth: 110,
		}
	],
	frameworkComponents: {
		gridBtn: GridButtonsVue
	},
});

// Actions funcs which passed to ag-grid
// and called after button pressed
const deleteFunc = (id: string) => {
	if (confirm("Вы действительно хотите удалить выбранную книгу?")) {
		store.deleteBook(id);
	}
};
const editFunc = (id: string) => {
	alert(`from edit ${id}`);
};

// Before Mount Event handler
onBeforeMount(() => {
	// Setup actions context
	actionContext.value = {
		deleteFunc: deleteFunc,
		editFunc: editFunc,
	};
});
</script>

<style>
</style>