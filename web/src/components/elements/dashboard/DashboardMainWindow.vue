<template>
  <div>
    <div
      class="uk-flex"
      uk-switcher="animation: uk-animation-fade; toggle: > *"
    >
      <button class="uk-button uk-button-default" type="button">Chart</button>
      <button class="uk-button uk-button-default" type="button">History</button>
      <button class="uk-button uk-button-default" type="button">
        Conditions
      </button>
    </div>

    <ul class="uk-switcher uk-margin">
      <li><ChartTab /></li>
      <li><HistoryTab /></li>
      <li><ConditionsTab /></li>
    </ul>
  </div>
</template>

<script setup>
import ChartTab from "@/components/elements/dashboard/ChartTab.vue";
import HistoryTab from "@/components/elements/dashboard/HistoryTab.vue";
import ConditionsTab from "@/components/elements/dashboard/ConditionsTab.vue";
</script>
