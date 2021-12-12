<template>
	<!-- <template v-if="isAuthorsLoaded">
		<div v-for="author in authors" :key="author.id">
			{{ author }}
		</div>
	</template>-->
	<!-- <button @click="getSelectedRows">Get Selected Rows</button>
	<button v-show="hasSelected" @click="clearSelection">Clear selection</button>
	<button @click="addNewRow">add new row</button>
	<button @click="clearGrid">Clear grid</button>-->

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
	></ag-grid-vue>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount, computed } from "vue";
import { useStore } from "../store";
import { storeToRefs } from "pinia";
import GridButtonsVue from "../components/GridButtons.vue";
import ActionsButtons from "../components/ActionsButtons.vue"
import { AgGridVue } from "ag-grid-vue3";
import { formatDate } from "../utils/date";
import {
	GridReadyEvent,
	GridApi,
	GridOptions,
	ColumnApi,
	RowSelectedEvent,
	SelectionChangedEvent,
} from "@ag-grid-community/all-modules";

// Reactive vars from store
const store = useStore();
const { isAuthorsLoaded, authors } = storeToRefs(store);

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const rowData = reactive([
	{
		id: "1",
		firstname: "Ivan",
		lastname: "Ivanov",
		surname: "Ivanovich",
		birth_date: formatDate(new Date),
		snippet: "This is Ivan Ivanov aka Triple I",
	},
	{
		id: "2",
		firstname: "Petr",
		lastname: "Petrov",
		surname: "Petrovich",
		birth_date: new Date(1999, 8, 31).toLocaleDateString("ru", {
			day: "2-digit",
			month: "2-digit",
			year: "numeric"
		}),
		snippet: "This is Petr Petrov aka Piper This is Petr Petrov aka Piper pThis is Petr Petrov aka Piper",
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

	gridApi.value?.setRowData(rowData);
};

const clearGrid = () => {
	rowData.splice(0, rowData.length);
	gridApi.value?.setRowData(rowData);
};

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
	rowData: rowData,
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