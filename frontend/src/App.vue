<script setup lang="ts">
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
// import HelloWorld from './components/HelloWorld.vue'
import { ref, reactive, onBeforeMount, onMounted } from "vue";
import { useStore } from "./store/index";
import { Book } from "./types";
import { GET } from "./HTTP";

const store = useStore();

onBeforeMount(async () => {
	// Load all data from API
	await Promise.all([
		store.loadAuthors(),
		store.loadGenres(),
		store.loadBooks(),
	]);
});
</script>

<template>
	<header class="app-header">
		<div class="header wrapper">
			<div class="header__intro">
				<h1 class="header__title">Лабораторная работа №1</h1>
				<h3 class="header__subtitle">
					По дисциплине "Программирование и администрирование в
					среде"
				</h3>
			</div>

			<nav class="header__nav">
				<router-link to="/">Главная</router-link>
				<router-link to="/books">Книги</router-link>
				<router-link to="/genres">Жанры</router-link>
				<router-link to="/authors">Авторы</router-link>
			</nav>
		</div>
	</header>

	<main class="app-main wrapper">
		<router-view></router-view>
	</main>
	<footer class="app-footer">
		<div class="footer wrapper">
			<div class="footer__credits">
				<p>
					Выполнил студент группы 15.11Д-МО12/19б факультета
					ИМИСиЦЭ Козлов Роман
				</p>
			</div>
			<div class="footer__link">
				<a
					href="https://github.com/the-NZA/DB_Lab1"
					target="_blank"
					title="Откроется в новой вкладке"
				>Репозиторий</a>
			</div>
		</div>
	</footer>
</template>

<style>
#app {
	min-height: 100vh;
	display: grid;
	grid-template-rows: max-content 1fr max-content;
	gap: var(--offset);
}

.app-header {
	background-color: rgb(var(--leanen));
}

.app-main {
	background-color: rgb(var(--white));
	width: 100%;
	padding: 0 var(--offset);
}

.app-footer {
	background-color: rgb(var(--greyblue));
}
</style>
