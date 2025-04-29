// main.js
import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Toast) // NOT createToastInterface

app.mount('#app')