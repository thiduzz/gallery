<template>
  <div>
    <div v-show="file" >
      <img ref="filePreview" src="#" alt="Selected image" />
      <div class="my-6 pt-3 rounded bg-gray-200">
        <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="subtitle">Subtitle</label>
        <input v-model="subtitle" type="text" id="subtitle" name="subtitle" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
      </div>
    </div>
    <div v-show="!file"  class="flex w-full items-center justify-center bg-grey-lighter my-5">
      <label class="w-64 flex flex-col items-center px-4 py-6 bg-white text-blue rounded-lg shadow-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue-400 hover:text-white">
        <svg class="w-8 h-8" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
          <path d="M16.88 9.1A4 4 0 0 1 16 17H5a5 5 0 0 1-1-9.9V7a3 3 0 0 1 4.52-2.59A4.98 4.98 0 0 1 17 8c0 .38-.04.74-.12 1.1zM11 11h3l-4-4-4 4h3v3h2v-3z" />
        </svg>
        <span class="mt-2 text-base leading-normal">Pick a photo</span>
        <input ref="fileInput" @change="showPreview" type="file" class="hidden" accept="image/png, image/jpeg" />
      </label>
    </div>
    <div v-if="file" class="w-full flex justify-between">
      <button @click="removeUpload" class="bg-red-600 px-3 hover:bg-red-700 text-white font-bold py-2 rounded shadow-lg hover:shadow-xl transition duration-200" type="button">Delete</button>
      <button @click="upload" class="bg-purple-600 px-3 hover:bg-purple-700 text-white font-bold py-2 rounded shadow-lg hover:shadow-xl transition duration-200" type="button">Save</button>
    </div>
  </div>
  <div v-if="photos.length > 0" class="">
    <PhotoItem v-for="photo in photos" :photo="photo" @delete="deleteUpload"></PhotoItem>
  </div>
</template>

<script>

import PhotoItem from "./PhotoItem.vue";
export default {
  components: {PhotoItem},
  props:{
    url: String,
    currentPhotos: Array
  },
  mounted() {
    this.photos = this.currentPhotos || []
  },
  data() {
    return {
      file: null,
      subtitle: '',
      photos: []
    }
  },
  methods: {
    upload() {
      if(this.file) {
        const formData = new FormData();
        formData.append("image", this.file);
        formData.append("subtitle", this.subtitle)
        this.axios.post(this.url, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }).then((response) => {
          this.photos.push(response.data)
          this.removeUpload()
        }).catch((e) => {
          this.$swal({icon:"error",title: "Oops!", text: e.response.data.error});
        });
      }
    },
    deleteUpload(photoId)
    {
      alert(photoId)
    },
    removeUpload() {
      this.file = null;
      this.subtitle = '';
      this.$refs.fileInput.value = '';
    },
    showPreview(){
      this.subtitle = '';
      if (this.$refs.fileInput.files && this.$refs.fileInput.files[0]) {
        this.file = this.$refs.fileInput.files[0]
        const reader = new FileReader();
        reader.onload = (e) => {
          this.$refs.filePreview.setAttribute("src",e.target.result)
        }
        reader.readAsDataURL(this.file);
      }else{
        this.file = null;
      }
    }
  }
}
</script>
