<script setup>
const props = defineProps({
  uploadHandler: Function,
  onCancelHandler: Function,
  files: Object
});

const onUpload = () => {
  props.uploadHandler(props.files)
}

const onCancel = () => {
  props.onCancelHandler();
}
</script>

<template>
  <footer class="flex justify-end px-8 pb-8 pt-4">
    <button
      id="submit"
      class="rounded-sm px-3 py-1 bg-blue-700 hover:bg-blue-500 text-white focus:shadow-outline focus:outline-none"
      @click="onUpload"
    >
      Загрузить
    </button>
    <button
      id="cancel"
      class="ml-3 rounded-sm px-3 py-1 hover:bg-gray-300 focus:shadow-outline focus:outline-none"
      @click="onCancel"
    >
      Отмена
    </button>
  </footer>
</template>
