<script setup>
import Empty from "./components/Empty.vue";
import Gallery from "./components/Gallery.vue";
import VideoTemplate from "./components/VideoTemplate.vue";
import Footer from "./components/Footer.vue";
import Overlay from "./components/Overlay.vue";
import { provide, ref } from "vue";
import axios from "axios";

const filesExists = ref(false);
const FILES = ref({});
const showSuccessToast = ref(false);
const results = ref([]);
const keys = Object.keys(FILES.value);
function addFile(file) {
  const objectURL = URL.createObjectURL(file);
  FILES.value[objectURL] = file;
}

const fileSizeToReadable = (file) => {
  const size =
    file.size > 1024
      ? file.size > 1048576
        ? Math.round(file.size / 1048576) + "mb"
        : Math.round(file.size / 1024) + "kb"
      : file.size + "b";

  return size;
};
// use to check if a file is being dragged
const hasFiles = ({ dataTransfer: { types = [] } }) =>
  types.indexOf("Files") > -1;

let counter = 0;
function dropHandler(ev) {
  ev.preventDefault();
  for (const file of ev.dataTransfer.files) {
    addFile(file);
    overlay.classList.remove("draggedover");
    counter = 0;
  }
}
function dragEnterHandler(e) {
  e.preventDefault();
  if (!hasFiles(e)) {
    return;
  }
  ++counter && overlay.classList.add("draggedover");
}

function dragLeaveHandler(e) {
  1 > --counter && overlay.classList.remove("draggedover");
}

function dragOverHandler(e) {
  if (hasFiles(e)) {
    e.preventDefault();
  }
}

const uploadOnclick = () => {
  const hidden = document.getElementById("hidden-input");
  hidden.click();
};
const inputOnChange = (e) => {
  for (const file of e.target.files) {
    addFile(file);
  }
};

const onDeleteHandler = (key) => {
  delete FILES.value[key];
};

const onCancelHandler = () => {
  FILES.value = {};
};

const uploadHandler = async (f) => {
  try {
    const formData = new FormData();
    for (const k in f) {
      formData.append("file", f[k]);
    }
    // if (!formData.getAll("file").length) {
    //   console.log("Нет данных");
    //   return;
    // }
    
    // const r = await axios.post("http://192.168.87.66:12013/upload", formData);
    const r = await axios.get("https://d9d5bc62a4d01ca2.mokky.dev/videos");
    if (r) {
      showSuccessToast.value = true;
    }
  } catch (error) {
    console.error(error);
  }
};

provide("files", {
  FILES,
});

const closeToast = () => {
  showSuccessToast.value = false;
}
const clickToggler = (e, d, f) => {
  e.preventDefault();
  let tabName = e.target.hash;

  let tabsContainer = document.querySelector("#tabs");
  let tabContents = document.querySelector("#tab-contents");
  let tabTogglers = tabsContainer.querySelectorAll("a");

  for (let i = 0; i < tabContents.children.length; i++) {
    tabTogglers[i].parentElement.classList.remove(
      "border-blue-400",
      "border-b",
      "-mb-px",
      "opacity-100"
    );
    tabContents.children[i].classList.remove("hidden");
    if ("#" + tabContents.children[i].id === tabName) {
      continue;
    }
    tabContents.children[i].classList.add("hidden");
  }
  e.target.parentElement.classList.add(
    "border-blue-400",
    "border-b-4",
    "-mb-px",
    "opacity-100"
  );
};

const watchVideo = (url) => {
  
}
</script>

<template>
  <!-- component -->
  <div class="bg-gray-500 h-screen w-screen sm:px-8 md:px-16 sm:py-8">
    <main class="container mx-auto max-w-screen-lg h-full">
      <!-- file upload modal -->
      <article
        aria-label="File Upload Modal"
        class="relative h-full flex flex-col bg-white shadow-xl rounded-md"
        @drop="dropHandler"
        @dragover="dragOverHandler"
        @dragleave="dragLeaveHandler"
        @dragenter="dragEnterHandler"
      >
        <!-- overlay -->
        <Overlay />

        <!-- scroll area -->

        <section class="h-full overflow-auto p-8 w-full h-full flex flex-col">
          <!-- header -->
          <header
            class="border-dashed border-2 border-gray-400 py-12 flex flex-col justify-center items-center"
          >
            <p
              class="mb-3 font-semibold text-gray-900 flex flex-wrap justify-center"
            >
              <span>Перетащите ваше видео</span>&nbsp;<span
                >в любое место чтобы добавить</span
              >
            </p>
            <input
              id="hidden-input"
              type="file"
              enctype="multipart/form-data"
              multiple
              class="hidden"
              @change="inputOnChange"
            />
            <button
              id="button"
              class="mt-2 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none"
              @click="uploadOnclick"
            >
              Добавить видео
            </button>
          </header>
          <!-- gallery -->

          <!-- вкладки -->
          <div class="w-full mx-auto mt-4 rounded">
            <!-- Tabs -->
            <ul id="tabs" class="inline-flex w-full px-1 pt-2">
              <li
                class="px-4 py-2 -mb-px font-semibold text-gray-800 border-b-2 border-blue-400 rounded-t opacity-50"
              >
                <a @click="clickToggler" id="default-tab" href="#first"
                  >Загрузка видео</a
                >
              </li>
              <li
                class="px-4 py-2 font-semibold text-gray-800 rounded-t opacity-50"
              >
                <a @click="clickToggler" href="#second">Результат</a>
              </li>
            </ul>

            <!-- Tab Contents -->
            <div id="tab-contents">
              <div id="first" class="p-4">
                <ul id="gallery" class="flex flex-1 flex-wrap -m-1">
                  <Empty
                    :class="{ hidden: Object.keys(FILES).length > 0 }"
                    :text="'Не выбрано видео'"
                  />
                  <VideoTemplate
                    v-for="key in Object.keys(FILES)"
                    :key="key"
                    :name="FILES[key].name"
                    :size="fileSizeToReadable(FILES[key])"
                    :onDeleteHandler="() => onDeleteHandler(key)"
                  />
                </ul>
              </div>
              <div id="second" class="hidden p-4">
                <ul id="gallery" class="flex flex-1 flex-wrap -m-1">
                  <Empty
                    :class="{ hidden: results.length > 0 }"
                    :text="'Hет результатов'"
                  />
                  <!-- <VideoTemplate /> -->
                </ul>
              </div>
            </div>
          </div>
        </section>

        <!-- sticky footer -->
        <Footer
          :files="FILES"
          :uploadHandler="uploadHandler"
          :onCancelHandler="onCancelHandler"
        />

        <div
          id="toast-success"
          class="fixed flex items-center w-full max-w-xs p-4 space-x-4 text-gray-500 bg-white divide-x rtl:divide-x-reverse divide-gray-200 rounded-lg shadow bottom-5 left-5 dark:text-gray-400 dark:divide-gray-700 space-x dark:bg-gray-800"
          role="alert"
          v-if="showSuccessToast"
        >
          <div
            class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200"
          >
            <svg
              class="w-5 h-5"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z"
              />
            </svg>
            <span class="sr-only">Check icon</span>
          </div>
          <div class="ms-3 text-sm font-normal">Item moved successfully.</div>
          <button
            type="button"
            class="ms-auto -mx-1.5 -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700"
            data-dismiss-target="#toast-success"
            aria-label="Close"
            @click="closeToast"
          >
            <span class="sr-only">Close</span>
            <svg
              class="w-3 h-3"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 14 14"
            >
              <path
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
              />
            </svg>
          </button>
        </div>
      </article>
    </main>
  </div>
</template>
