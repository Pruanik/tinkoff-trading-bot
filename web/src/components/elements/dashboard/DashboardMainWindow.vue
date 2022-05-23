<template>
  <div id="dashboard-main-window">
    <div class="uk-flex uk-flex-wrap">
      <button
        @click="switchComponent('ChartTab')"
        class="uk-button uk-button-default"
        type="button"
      >
        Chart
      </button>
      <button
        @click="switchComponent('HistoryTab')"
        class="uk-button uk-button-default"
        type="button"
      >
        History
      </button>
      <button
        @click="switchComponent('ConditionsTab')"
        class="uk-button uk-button-default"
        type="button"
      >
        Conditions
      </button>
      <button
        @click="switchComponent('SettingsTab')"
        class="uk-button uk-button-default"
        type="button"
      >
        Settings
      </button>
    </div>

    <component :is="activeWindow" class="uk-margin"></component>
  </div>
</template>

<script>
import ChartTab from "@/components/elements/dashboard/tabs/ChartTab.vue";
import HistoryTab from "@/components/elements/dashboard/tabs/HistoryTab.vue";
import ConditionsTab from "@/components/elements/dashboard/tabs/ConditionsTab.vue";
import SettingsTab from "@/components/elements/dashboard/tabs/SettingsTab.vue";

export default {
  props: ['figi'],
  data() {
    return {
      activeWindow: "chartTab",
    };
  },
  components: {
    ChartTab,
    HistoryTab,
    ConditionsTab,
    SettingsTab,
  },
  methods: {
    switchComponent: function (component) {
      this.activeWindow = component;
    },
  },
};
</script>
