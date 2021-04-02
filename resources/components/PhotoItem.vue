<template>
  <div class="h-24 w-24 block relative">
    <div @click="removePhoto" class="bg-white cursor-pointer absolute h-5 w-5 text-xs text-center rounded-full shadow-lg p-3 -right-3 -top-3 flex justify-center items-center">❌</div>
    <div class="w-full h-full bg-cover bg-center  rounded-lg shadow-lg" v-bind:style="{ backgroundImage: 'url(' + photo.url + ')' }"></div>
  </div>
</template>

<script>

export default {
  props:{
    photo: Object
  },
  methods: {
    removePhoto() {
      this.$emit("delete",this.photo.id)
    },
  }
}
</script>
