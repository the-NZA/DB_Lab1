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
		:rowData="getGenresRows"
	></ag-grid-vue>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount, onMounted, computed, watch } from "vue";
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

const store = useStore();
const { isGenresLoaded, getGenresRows } = storeToRefs(store);

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables for buttons
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const handleAdd = () => {
	console.log("add was pressed");
	store.addGenre()
}

const handleEdit = () => {
	console.log("edit was pressed");
}

const handleDelete = () => {
	console.log("delete was pressed");
}

// Handle selection changed event
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

// Current grid options
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
			flex: 2,
			minWidth: 250,
		},
		{
			field: "snippet",
			headerName: "Описание",
			flex: 3,
			minWidth: 450,
		},
		{
			field: "actions",
			headerName: "Действия",
			cellRenderer: "gridBtn",
			// cellClass: ["test", "test2"], // Use for grid action buttons
			width: 105,
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