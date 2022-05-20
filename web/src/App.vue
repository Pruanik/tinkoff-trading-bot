<script setup>
import { RouterView } from "vue-router";
import HeaderPage from "@/components/layouts/HeaderPage.vue";
</script>

<template>
  <HeaderPage />
  <main>
    <RouterView />
  </main>
</template>
