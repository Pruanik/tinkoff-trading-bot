<template>
  <ul uk-tab>
    <li>
      <RouterLink activeClass="uk-active" to="/"> Statistic </RouterLink>
    </li>
    <li>
      <RouterLink activeClass="uk-active" to="/instruments">
        Instruments
      </RouterLink>
    </li>
    <li>
      <RouterLink activeClass="uk-active" to="/logs"> Logs </RouterLink>
    </li>
  </ul>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>
