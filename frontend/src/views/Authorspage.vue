<template>
	<!-- <button @click="getSelectedRows">Get Selected Rows</button>
	<button v-show="hasSelected" @click="clearSelection">Clear selection</button>
	-->

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
		:rowData="getAuthorsRows"
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

// Reactive vars from store
const store = useStore();
const { isAuthorsLoaded, getAuthorsRows } = storeToRefs(store);

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const getSelectedRows = () => {
	const selected = gridApi.value?.getSelectedRows();
	console.log(selected);
};

const clearSelection = () => {
	gridApi.value?.deselectAll();
};

const handleAdd = () => {
	console.log("add was pressed");
}

const handleEdit = () => {
	console.log("edit was pressed");
}

const handleDelete = () => {
	console.log("delete was pressed");
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
			field: "firstname",
			checkboxSelection: true,
			headerCheckboxSelection: true,
			headerName: "Имя",
			flex: 1,
		},
		{
			field: "lastname",
			headerName: "Фамилия",
			flex: 1,
		},
		{
			field: "surname",
			headerName: "Отчество",
			flex: 1,
		},
		{
			field: "birth_date",
			headerName: "Дата рождения",
			maxWidth: 140,
		},
		{
			field: "snippet",
			headerName: "Описание",
			flex: 2,
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
	alert(`from delete ${id}`);
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