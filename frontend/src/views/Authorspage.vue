<template>
	<!-- <template v-if="isAuthorsLoaded">
		<div v-for="author in authors" :key="author.id">
			{{ author }}
		</div>
	</template>-->
	<button @click="getSelectedRows">Get Selected Rows</button>
	<button v-show="hasSelected" @click="clearSelection">Clear selection</button>
	<button @click="addNewRow">add new row</button>
	<button @click="clearGrid">Clear grid</button>

	<!-- 
		Buttons to add: 
			1. Add new item – Always
			2. Edit – If selected only 1 item 
			3. Delete – If selected at least 1 item
	-->

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
import { AgGridVue } from "ag-grid-vue3";
import {
	GridReadyEvent,
	GridApi,
	GridOptions,
	ColumnApi,
	RowSelectedEvent,
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

// Row Selected event
const onRowSelected = (e: RowSelectedEvent) => {
	// console.log(e);

	// Set have or haven't selected rows var
	if (gridApi.value && gridApi.value.getSelectedRows().length > 0) {
		hasSelected.value = true;
	} else {
		hasSelected.value = false;
	}
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
			flex: 0.5,
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
:root {
	--ag-header-foreground-color: rgb(var(--sapphire));
	--ag-header-background-color: rgb(var(--leanen));
	--ag-row-hover-color: rgba(var(--redsand), 0.25);
}

.ag-theme-alpine {
	box-shadow: var(--main-shadow);
	border-radius: var(--offset-quarter);
}

.ag-theme-alpine .ag-root-wrapper {
	/* border: 1px solid rgb(var(--sapphire)) !important; */
	border-radius: var(--offset-quarter);
}

.ag-cell-wrapper .ag-selection-checkbox {
	margin-right: 18px !important;
}

.ag-cell-wrap-text {
	word-break: normal !important;
}
</style>