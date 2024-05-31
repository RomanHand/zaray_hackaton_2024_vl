<script setup></script>

<template>
  <h1 class="pt-8 pb-3 font-semibold sm:text-lg text-gray-900">Видео к загрузке </h1>

  <ul id="gallery" class="flex flex-1 flex-wrap -m-1">
    <li
      id="empty"
      class="h-full w-full text-center flex flex-col items-center justify-center items-center"
    >
      <img
        class="mx-auto w-32"
        src="https://user-images.githubusercontent.com/507615/54591670-ac0a0180-4a65-11e9-846c-e55ffce0fe7b.png"
        alt="no data"
      />
      <span class="text-small text-gray-500">Не выбрано видео</span>
    </li>
  </ul>
</template>
