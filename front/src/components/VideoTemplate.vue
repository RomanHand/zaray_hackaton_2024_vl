<script setup>
const props = defineProps({
  name: String,
  size: String,
  onDeleteHandler: Function,
});

function onDelete(e) {
  props.onDeleteHandler(e);
}
</script>

<template>
  <li class="block p-1 w-1/2 sm:w-1/3 md:w-1/4 lg:w-1/6 xl:w-1/8 h-24">
    <article
      tabindex="0"
      class="group w-full h-full rounded-md focus:outline-none focus:shadow-outline elative bg-gray-100 cursor-pointer relative shadow-sm"
    >
      <img
        alt="upload preview"
        class="img-preview hidden w-full h-full sticky object-cover rounded-md bg-fixed"
      />

      <section
        class="flex flex-col rounded-md text-xs break-words w-full h-full z-20 absolute top-0 py-2 px-3"
      >
        <h1 class="flex-1 group-hover:text-blue-800">{{ name }}</h1>
        <div class="flex">
          <span class="p-1 text-blue-800">
            <i>
              <svg
                class="fill-current w-4 h-4 ml-auto pt-1"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <path d="M15 2v5h5v15h-16v-20h11zm1-2h-14v24h20v-18l-6-6z" />
              </svg>
            </i>
          </span>
          <p class="p-1 size text-xs text-gray-700">{{ size }}</p>
          <button
            class="delete ml-auto focus:outline-none hover:bg-gray-300 p-1 rounded-md text-gray-800"
            @click="onDelete"
          >
            <svg
              class="pointer-events-none fill-current w-4 h-4 ml-auto"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
            >
              <path
                class="pointer-events-none"
                d="M3 6l3 18h12l3-18h-18zm19-4v2h-20v-2h5.711c.9 0 1.631-1.099 1.631-2h5.316c0 .901.73 2 1.631 2h5.711z"
              />
            </svg>
          </button>
        </div>
      </section>
    </article>
  </li>
</template>
