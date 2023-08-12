import {createApp} from 'vue'

import App from './App.vue'

// Sidebar things
import VueSidebarMenu from 'vue-sidebar-menu'
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'
import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
    { path: '/', component: App },
]

const app = createApp(App)

const router = createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHashHistory(),
    routes,  // short for `routes: routes`
})

app.use(VueSidebarMenu)
app.use(router)

app.mount('#app')
