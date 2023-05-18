import {createApp} from 'vue'
//import {Vue} from 'vue'
import App from './App.vue'
import VueApexCharts from 'vue3-apexcharts'

//createApp.use(VueApexCharts)
//createApp.component('apexchart', VueApexCharts)

//createApp(App).mount('#app')

const app = createApp(App)
//app.use(VueApexCharts)
app.component('apexchart', VueApexCharts)
app.mount('#app')
