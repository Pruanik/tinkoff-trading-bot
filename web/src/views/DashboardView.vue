<script setup>
import InstrumentsSidebar from "@/components/elements/dashboard/InstrumentsSidebar.vue";
import DashboardMainWindow from "@/components/elements/dashboard/DashboardMainWindow.vue";
</script>

<template>
  <div class="uk-container uk-margin-medium-top">
    <h1 class="uk-heading-line"><span>Dashboard</span></h1>
    <div class="uk-text-center" uk-grid>
      <div class="uk-width-1-4@m">
        <div class="uk-card uk-card-default uk-card-body">
          <InstrumentsSidebar />
        </div>
      </div>
      <div class="uk-width-expand@m">
        <div class="uk-card uk-card-default uk-card-body">
          <DashboardMainWindow />
        </div>
      </div>
    </div>
  </div>
</template>
