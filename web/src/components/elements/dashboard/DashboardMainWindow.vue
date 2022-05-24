<template>
  <div id="dashboard-main-window">
    <div v-if="activeInstrument.name == ''">
      <hr class="uk-divider-icon" />
      <h1 class="uk-heading-small">
        Let's choose one instrument on the sidebar.
      </h1>
      <hr class="uk-divider-icon" />
    </div>
    <div v-else>
      <h2 class="uk-heading-small uk-text-left">
        {{ activeInstrument.name }}
      </h2>
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

      <component :is="activeWindow" v-bind:activeInstrument="activeInstrument" class="uk-margin"></component>
    </div>
  </div>
</template>

<script>
import ChartTab from "@/components/elements/dashboard/tabs/ChartTab.vue";
import HistoryTab from "@/components/elements/dashboard/tabs/HistoryTab.vue";
import ConditionsTab from "@/components/elements/dashboard/tabs/ConditionsTab.vue";
import SettingsTab from "@/components/elements/dashboard/tabs/SettingsTab.vue";

export default {
  props: ['activeInstrument'],
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
