<template>
  <ul uk-tab>
    <DashboardCommonMenuLink activeClass="uk-active" to="/"> Statistic </DashboardCommonMenuLink>
    <DashboardCommonMenuLink activeClass="uk-active" to="/instruments">
      Instruments
    </DashboardCommonMenuLink>
    <DashboardCommonMenuLink activeClass="uk-active" to="/logs">
      Logs
    </DashboardCommonMenuLink>
  </ul>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
import DashboardCommonMenuLink from "./DashboardCommonMenuLink.vue";
</script>
