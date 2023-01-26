<template>
  <AppLayoutHeader />
  <main class="uk-container uk-container-expand">
    <RouterView />
  </main>
  <AppLayoutSidebarMenu />
</template>

<script lang="ts" setup>
import { RouterView } from "vue-router";
import AppLayoutHeader from "./AppLayoutHeader.vue";
import AppLayoutSidebarMenu from "./AppLayoutSidebarMenu.vue";
</script>
