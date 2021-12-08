<script setup lang="ts">
import { ref, reactive, onMounted } from "vue"
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
// import HelloWorld from './components/HelloWorld.vue'

import FirstTab from "./components/FirstTab.vue"
import SecondTab from "./components/SecondTab.vue"

const tabview = ref()
const tabs = reactive([0, 1, 2])
const tabsHeaders = ref<HTMLButtonElement[]>([])
const tabPannels = ref<HTMLDivElement[]>([])
let selectedTab = ref(0)

onMounted(() => {
	console.log(tabview.value);
})

const selectTab = (e: MouseEvent) => {
	selectedTab.value = parseInt((e.target as HTMLButtonElement).dataset["tabindex"] || "")

	tabsHeaders.value.forEach((p, i) => {
		if (i === selectedTab.value) {
			p.classList.add("activetab")
		} else {
			p.classList.remove("activetab")
		}
	})

	tabPannels.value.forEach((p, i) => {
		if (i === selectedTab.value) {
			p.classList.add("activepanel")
		} else {
			p.classList.remove("activepanel")
		}
	})
}


</script>

<template>
	<header class="app-header">
		<div class="header wrapper">
			<h1 class="header__title">Лабораторная работа №1</h1>
			<h3 class="header__subtitle">По дисциплине "Программирование и администрирование в среде"</h3>
		</div>
	</header>
	<main class="app-main">
		<div class="tabview" ref="tabview">
			<div class="tabview__header">
				<button
					class="tabview_btn"
					@click="selectTab"
					v-for="t in tabs"
					:data-tabindex="t"
					:ref="(el: any) => {
						if (el) {
							tabsHeaders[t] = el;
						}
					}"
				>Tab{{ t }}</button>
			</div>

			<div class="tabview__pannels">
				<div
					class="tabview__panel"
					v-for="t in tabs"
					v-show="t == selectedTab ? true : false"
					:ref="(el: any) => { if (el) tabPannels[t] = el }"
				>
					<first-tab />
				</div>
			</div>
		</div>
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
