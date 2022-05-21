<script setup>
import LogTable from "@/components/elements/logs/LogTable.vue";
</script>

<template>
  <div class="uk-container uk-margin-medium-top">
    <h1 class="uk-heading-line"><span>Logs</span></h1>
    <div class="uk-text-center" uk-grid>
      <div class="uk-width-1-1@m">
        <div class="uk-card uk-card-default uk-card-body">
          <LogTable />
        </div>
      </div>
    </div>
  </div>
</template>
