<script setup>
import { inject } from "vue";
import Empty from "./Empty.vue";
import VideoTemplate from './VideoTemplate.vue'

const props = defineProps({
  files: Object
});

const files = props.files;
const hasFiles = (f) =>  Object.keys(f).length > 0;
const keys = Object.keys(files);

for (const key in keys) {
  console.log(files[key]);
}
</script>

<template>
  <h1 class="pt-8 pb-3 font-semibold sm:text-lg text-gray-900">Видео к загрузке </h1>

  <ul id="gallery" class="flex flex-1 flex-wrap -m-1">
    <Empty :class="{hidden: hasFiles(files)}"/>
    <VideoTemplate v-for="key in keys" :key="key" :name="files[key].name" :size="files[key].size"/>
  </ul>
</template>
