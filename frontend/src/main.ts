import { createApp } from 'vue'
import {createPinia} from 'pinia'
import router from "./router"
import App from './App.vue'

// Import base css
import "normalize.css"
import "./css/variables.css"
import "./css/core.css"

// Import header css
import "./css/header/header.css"

// Import footer css
import "./css/footer/footer.css"

createApp(App)
.use(createPinia())
.use(router)
.mount('#app')
