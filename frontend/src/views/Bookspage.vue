<template>
	<!-- <template v-if="isBooksLoaded">
		<div v-for="book in books" :key="book.id">
			{{ book }}
		</div>
	</template>-->

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
	SelectionChangedEvent,
} from "@ag-grid-community/all-modules";
import { BookRow } from "../types/grid";

const store = useStore();
const { isBooksLoaded, books } = storeToRefs(store);

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const rowData = reactive<BookRow[]>([
	{
		id: "1",
		title: "Первая книга",
		genre: "Роман",
		pages_cnt: 123,
		book_lang: "Русский",
		pub_date: formatDate(new Date),
		snippet: "Первая книга написана ради написания",
	},
	{
		id: "2",
		title: "Книга для второго игрока",
		genre: "Фантастика",
		pages_cnt: 883,
		book_lang: "Русский",
		pub_date: formatDate(new Date(2019, 3, 21)),
		snippet: "Игроку номер два нужно не только быть готовым, но и в готовности",
	},
]);

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
			field: "book_lang",
			headerName: "Язык",
			minWidth: 120,
			maxWidth: 140,
		},
		{
			field: "pages_cnt",
			headerName: "Страниц",
			width: 100,
		},
		{
			field: "pub_date",
			headerName: "Дата публикации",
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