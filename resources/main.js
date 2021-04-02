import { createApp } from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';

import PhotoUploadComponent from './components/PhotoUploadComponent.vue'

const app = createApp({})
app.use(VueAxios, axios)
app.use(VueSweetalert2);
app.component('photo-upload', PhotoUploadComponent)
app.mount('#app')
