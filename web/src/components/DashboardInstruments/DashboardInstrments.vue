<template>
  <DashboardCommonHeader />
  <div class="uk-card uk-card-default uk-card-body">
    <DashboardInstrmentsButtons />
    <DashboardInstrmentsTable />
  </div>
</template>

<script setup lang="ts">
import DashboardCommonHeader from "@/components/DashboardCommon/DashboardCommonHeader.vue";
import DashboardInstrmentsButtons from "./DashboardInstrmentsButtons.vue";
import DashboardInstrmentsTable from "./DashboardInstrmentsTable.vue";
</script>
