<script setup lang="ts">
	// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
	// import HelloWorld from './components/HelloWorld.vue'
	import { ref, reactive, onBeforeMount, onMounted } from "vue";
	import { useStore } from "./store/index";
	import { Book } from "./types";
	import { GET } from "./HTTP";

	const store = useStore();

	onBeforeMount(async () => {
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
			<h1 class="header__title">Лабораторная работа №1</h1>
			<h3 class="header__subtitle">
				По дисциплине "Программирование и администрирование в
				среде"
			</h3>
		</div>
	</header>

	<nav class="app-nav">
		<router-link to="/">Главная</router-link>
		<router-link to="/books">Книги</router-link>
		<router-link to="/genres">Жанры</router-link>
		<router-link to="/authors">Авторы</router-link>
	</nav>

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
					>Репозиторий</a
				>
			</div>
		</div>
	</footer>
</template>

<style lang="postcss">
#app {
	min-height: 100vh;
	display: grid;
	grid-template-rows: max-content max-content 1fr max-content;
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

.app-nav {
	max-width: max-content;
	margin: 0 auto;
	display: flex;
	justify-content: space-evenly;

	background-color: rgb(var(--leanen));
	border-radius: var(--offset-quarter);
	padding: calc(var(--offset-half) * 1.5);
	box-shadow: 0px 2px 5px rgba(var(--black), 0.3);
}

.app-nav a {
	display: block;
	min-width: 100px;

	padding: var(--offset-half);
	margin-right: var(--offset);

	color: var(--sapphire);
	font-weight: bold;
	text-decoration: none;
	text-align: center;

	background-color: rgb(var(--white));
	transition: var(--main-transition);

	border-radius: var(--offset-quarter);
	box-shadow: 0px 4px 12px 2px rgba(var(--black), 0.1);
}

.app-nav a:last-of-type {
	margin-right: 0;
}

.app-nav a:hover {
	transform: scale(1.04);
	color: rgb(var(--greyblue));
}

.app-nav a.router-link-active {
	color: rgb(var(--white));
	background-color: rgb(var(--redsand));
}

.app-nav a.router-link-active:hover {
	transform: scale(1);
	color: rgb(var(--leanen));
}
</style>
