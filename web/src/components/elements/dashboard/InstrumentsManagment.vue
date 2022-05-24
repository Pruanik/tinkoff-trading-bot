<template>
  <div
    class="uk-grid-small uk-child-width-1-2@s uk-flex-center uk-text-center"
    uk-grid
  >
    <div class="uk-width-1-1@m">
      <div class="uk-card uk-card-default uk-card-body uk-flex uk-flex-row">
        <v-select
          class="uk-flex-left"
          style="width: 100%"
          :options="selectOptions"
          label="instrumentName"
          index="figi"
          v-model="currentInstrumentFigiAdding"
        />
        <div
          v-if="sendRequestStatus"
          uk-spinner
          class="uk-margin-small-left"
        ></div>
        <button
          v-else
          class="uk-button uk-button-default uk-margin-small-left uk-flex-right"
          v-on:click="setCollectingInstrument()"
        >+</button>
      </div>
    </div>
    <div class="uk-width-1-5@m">
      <div class="uk-card uk-card-default uk-card-body">
        <ul class="uk-tab-left" uk-tab="animation: uk-animation-fade">
          <li
            v-for="collectingInstrument in collectingInstrumentsList"
            :key="collectingInstrument.Figi"
          >
            <a
              href="#"
              v-on:click="
                setActiveInstrument(
                  collectingInstrument.Name,
                  collectingInstrument.Figi
                )
              "
            >
              {{ collectingInstrument.Name }} ({{ collectingInstrument.Figi }})
            </a>
          </li>
        </ul>
      </div>
    </div>
    <div class="uk-width-4-5@m">
      <div class="uk-card uk-card-default uk-card-body">
        <DashboardMainWindow v-bind:activeInstrument.sync="activeInstrument" />
      </div>
    </div>
  </div>
</template>

<script>
import vSelect from "vue-select";
import axios from "axios";
import uikit from "uikit";
import DashboardMainWindow from "@/components/elements/dashboard/DashboardMainWindow.vue";

export default {
  name: "InstrumentsSidebar",
  components: {
    "v-select": vSelect,
    DashboardMainWindow,
  },
  data() {
    return {
      activeInstrument: { figi: "", name: "" },
      sendRequestStatus: false,
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
      if (Object.keys(this.currentInstrumentFigiAdding ?? {}).length == 0) {
        this.warningNotification("You should choose instrument!");
        return;
      }
      this.sendRequestStatus = true;
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
          this.sendRequestStatus = false;
        })
        .catch((error) => console.log(error));
    },
    setActiveInstrument: function (name, figi) {
      this.activeInstrument.name = name;
      this.activeInstrument.figi = figi;
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
