<template>
  <div>
    <div class="uk-flex uk-flex-column">
      <v-select
        :options="selectOptions"
        label="instrumentName"
        index="figi"
        v-model="currentInstrumentFigiAdding"
      />
      <button
        class="uk-button uk-button-default uk-margin-small-top"
        v-on:click="setCollectingInstrument()"
      >
        +
      </button>
    </div>

    <ul class="uk-tab-left" uk-tab="animation: uk-animation-fade">
      <li
        v-for="collectingInstrument in collectingInstrumentsList"
        :key="collectingInstrument.Figi"
      >
        <a href="#">
          {{ collectingInstrument.Name }} ({{ collectingInstrument.Figi }})
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import vSelect from "vue-select";
import axios from "axios";
import uikit from "uikit";

export default {
  name: "InstrumentsSidebar",
  components: {
    "v-select": vSelect,
  },
  data() {
    return {
      currentInstrumentFigiAdding: {},
      selectOptions: [],
      collectingInstrumentsList: [],
    };
  },
  methods: {
    getInstrumentsList() {
      axios
        .get("/api/getInstruments")
        .then((response) => this.handleInstrumentsResponse(response.data))
        .catch((error) => console.log(error));
    },
    handleInstrumentsResponse: function (data) {
      if (data.Status.Status !== "success") {
        this.warningNotification(data.Status.Message);
      }
      if (data.Body === null) {
        return;
      }
      let options = [];
      data.Body.forEach(function (item) {
        let option = {};
        option.instrumentName = item.Name + " (" + item.Figi + ")";
        option.figi = item.Figi;
        options.push(option);
      });
      this.selectOptions = options;
    },
    getCollectingInstrumentsList() {
      axios
        .get("/api/getCollectingInstruments")
        .then((response) =>
          this.handleCollectingInstrumentsResponse(response.data)
        )
        .catch((error) => console.log(error));
    },
    handleCollectingInstrumentsResponse: function (data) {
      if (data.Status.Status !== "success") {
        this.warningNotification(data.Status.Message);
      }
      if (data.Body === null) {
        return;
      }
      this.collectingInstrumentsList = data.Body;
    },
    setCollectingInstrument() {
      if (!this.currentInstrumentFigiAdding.hasOwnProperty("figi")) {
        this.warningNotification("You should choose instrument!");
        return;
      }

      axios
        .get("/api/setCollectingInstrument", {
          params: { figi: this.currentInstrumentFigiAdding.figi, status: true },
        })
        .then((response) => {
          if (response.data.Status.Status !== "success") {
            this.warningNotification(response.data.Status.Message);
          } else {
            this.getCollectingInstrumentsList();
          }
        })
        .catch((error) => console.log(error));
    },
    warningNotification: function (message) {
      uikit.notification({
        message: message,
        status: "warning",
        pos: "top-right",
        timeout: 5000,
      });
    },
  },
  mounted() {
    this.getInstrumentsList();
    this.getCollectingInstrumentsList();
  },
};
</script>
