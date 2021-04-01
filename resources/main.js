import { createApp } from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import PhotoUploadComponent from './components/PhotoUploadComponent.vue'

const app = createApp({})
app.use(VueAxios, axios)
app.component('photo-upload', PhotoUploadComponent)
app.mount('#app')
