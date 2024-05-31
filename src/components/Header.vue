<script setup></script>

<template>
  <header
    class="border-dashed border-2 border-gray-400 py-12 flex flex-col justify-center items-center"
  >
    <p class="mb-3 font-semibold text-gray-900 flex flex-wrap justify-center">
      <span>Перетащите ваше видео</span>&nbsp;<span>в любое место чтобы добавить</span>
    </p>
    <input id="hidden-input" type="file" multiple class="hidden" />
    <button
      id="button"
      class="mt-2 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none"
    >
      Добавить видео
    </button>
  </header>
</template>
