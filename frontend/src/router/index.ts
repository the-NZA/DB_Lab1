import { RouteRecordRaw, createRouter, createWebHistory} from 'vue-router'

import Homepage from "../views/Homepage.vue"
import FirstTab from "../components/FirstTab.vue"
import SecondTab from "../components/SecondTab.vue"

const routes: RouteRecordRaw[] = [
	{
		path: '/',
		name: 'Home',
		component: Homepage,
	},
	{
		path: '/second',
		name: 'Second view',
		component: SecondTab,
	},
	{
		path: '/third',
		name: 'Third view',
		component: FirstTab,
	},
]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

export default router