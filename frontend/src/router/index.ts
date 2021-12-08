import { RouteRecordRaw, createRouter, createWebHistory} from 'vue-router'

import FirstTab from "../components/FirstTab.vue"
import SecondTab from "../components/SecondTab.vue"

const routes: RouteRecordRaw[] = [
	{
		path: '/',
		name: 'Home',
		component: FirstTab,
	},
	{
		path: '/second',
		name: 'Second view',
		component: SecondTab,
	},
]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

export default router