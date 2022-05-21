<template>
  <div id="dashboard-main-window">
    <div
      class="uk-flex uk-flex-wrap"
      uk-switcher="animation: uk-animation-fade; toggle: > *"
    >
      <button class="uk-button uk-button-default" type="button">Chart</button>
      <button class="uk-button uk-button-default" type="button">History</button>
      <button class="uk-button uk-button-default" type="button">
        Conditions
      </button>
      <button class="uk-button uk-button-default" type="button">
        Settings
      </button>
    </div>

    <ul class="uk-switcher uk-margin">
      <li><ChartTab /></li>
      <li><HistoryTab /></li>
      <li><ConditionsTab /></li>
      <li><SettingsTab /></li>
    </ul>
  </div>
</template>

<script setup>
import ChartTab from "@/components/elements/dashboard/tabs/ChartTab.vue";
import HistoryTab from "@/components/elements/dashboard/tabs/HistoryTab.vue";
import ConditionsTab from "@/components/elements/dashboard/tabs/ConditionsTab.vue";
import SettingsTab from "@/components/elements/dashboard/tabs/SettingsTab.vue";
</script>
