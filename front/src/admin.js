import { createApp } from 'vue'
import Admin from './app/Admin.vue'
import { createPinia } from 'pinia'

const app = createApp(Admin)

app.use(createPinia())
app.mount("#app")



