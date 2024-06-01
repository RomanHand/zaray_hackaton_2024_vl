<script setup>
const props = defineProps({
  text: String
});
</script>

<template>
  <li
    id="empty"
    class="h-full w-full text-center flex flex-col items-center justify-center items-center"
  >
    <img
      class="mx-auto w-32"
      src="https://user-images.githubusercontent.com/507615/54591670-ac0a0180-4a65-11e9-846c-e55ffce0fe7b.png"
      alt="no data"
    />
    <span class="text-small text-gray-500">{{ text }}</span>
  </li>
</template>
