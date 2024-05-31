<script setup>
import Empty from "./components/Empty.vue";
import Gallery from "./components/Gallery.vue";
import VideoTemplate from "./components/VideoTemplate.vue";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import Overlay from "./components/Overlay.vue";
</script>

<template>
  <!-- component -->
  <div class="bg-gray-500 h-screen w-screen sm:px-8 md:px-16 sm:py-8">
    <main class="container mx-auto max-w-screen-lg h-full">
      <!-- file upload modal -->
      <article
        aria-label="File Upload Modal"
        class="relative h-full flex flex-col bg-white shadow-xl rounded-md"
        ondrop="dropHandler(event);"
        ondragover="dragOverHandler(event);"
        ondragleave="dragLeaveHandler(event);"
        ondragenter="dragEnterHandler(event);"
      >
        <!-- overlay -->
        <Overlay/>

        <!-- scroll area -->
        <section class="h-full overflow-auto p-8 w-full h-full flex flex-col">
          <!-- header -->
          <Header />
          <!-- gallery -->
          <Gallery />
        </section>

        <!-- sticky footer -->
        <Footer />
      </article>
    </main>
  </div>
</template>