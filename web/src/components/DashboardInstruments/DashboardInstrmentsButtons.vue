<template>
  <div class="uk-flex">
    <DashboardInstrmentsModal />
    <form class="uk-search uk-search-default uk-margin-left">
        <a href="" class="uk-search-icon-flip" uk-search-icon></a>
        <input class="uk-search-input" type="search" placeholder="Search" aria-label="Search">
    </form>
  </div>
</template>

<script setup lang="ts">
import DashboardInstrmentsModal from './DashboardInstrmentsModal.vue';

</script>
