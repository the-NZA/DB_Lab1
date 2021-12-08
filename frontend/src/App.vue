<script setup lang="ts">
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
	<header class="app-header">
		<div class="header wrapper">
			<h1 class="header__title">Лабораторная работа №1</h1>
			<h3 class="header__subtitle">По дисциплине "Программирование и администрирование в среде"</h3>
		</div>
	</header>
	<main class="app-main">
		<img alt="Vue logo" src="./assets/logo.png" />
		<hello-world msg="Hello Vue 3 + TypeScript + Vite" />
	</main>
	<footer class="app-footer">
		<div class="footer wrapper">
			<div class="footer__credits">
				<p>Выполнил студент группы 15.11Д-МО12/19б факультета ИМИСиЦЭ Козлов Роман</p>
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
}

.app-footer {
	background-color: rgb(var(--greyblue));
}
</style>
