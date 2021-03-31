import { createApp } from 'vue'
import PhotoUploadComponent from './components/PhotoUploadComponent.vue'

const app = createApp({})
app.component('photo-upload', PhotoUploadComponent)
app.mount('#app')
