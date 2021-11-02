import { createApp } from 'vue'
import App from './App.vue'
import 'windi.css'
import './assets/style/global.scss'
import VCalendar from 'v-calendar'

const app = createApp(App)

app.use(VCalendar, {}).mount('#app')
