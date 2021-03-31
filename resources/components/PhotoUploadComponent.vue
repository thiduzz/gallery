<template>
  <div>
    <div class="flex w-full items-center justify-center bg-grey-lighter my-5">
      <label class="w-64 flex flex-col items-center px-4 py-6 bg-white text-blue rounded-lg shadow-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue-400 hover:text-white">
        <svg class="w-8 h-8" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
          <path d="M16.88 9.1A4 4 0 0 1 16 17H5a5 5 0 0 1-1-9.9V7a3 3 0 0 1 4.52-2.59A4.98 4.98 0 0 1 17 8c0 .38-.04.74-.12 1.1zM11 11h3l-4-4-4 4h3v3h2v-3z" />
        </svg>
        <span class="mt-2 text-base leading-normal">Select a file</span>
        <input type='file' class="hidden" />
      </label>
    </div>
    <div class="mb-6 pt-3 rounded bg-gray-200">
      <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="title">Title</label>
      <input type="text" id="title" name="title" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
    </div>
    <button @click="upload" class="bg-purple-600 hover:bg-purple-700 text-white font-bold py-2 rounded shadow-lg hover:shadow-xl transition duration-200" type="button">Save</button>
  </div>
</template>

<script>

export default {
  props:{
    galleryId: Number
  },
  data() {
    return {
    }
  },

  methods: {
    upload() {
      debugger;
      console.log("upload!")
    },
  }
}
</script>
