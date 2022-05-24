<template>
  <div>
    <div class="uk-flex">
      <select v-model="updateInterval" class="uk-select uk-width-1-6">
        <option
          v-for="option in optionsInterval"
          :value="option.value"
          :key="option.value"
        >
          {{ option.text }}
        </option>
      </select>
      <select v-model="updatePeriod" class="uk-select uk-width-1-6 uk-margin-small-left">
        <option
          v-for="option in optionPeriod"
          :value="option.value"
          :key="option.value"
        >
          {{ option.text }}
        </option>
      </select>
    </div>
    <apexcharts
      ref="realtimeChart"
      width="800"
      height="700"
      type="line"
      :options="chartOptions"
      :series="series"
    ></apexcharts>
  </div>
</template>

<script>
import VueApexCharts from "vue3-apexcharts";
import axios from "axios";
import uikit from "uikit";

const chart = { type: "line" };
const xaxis = { type: "datetime" };

export default {
  name: "ChartTab",
  components: {
    apexcharts: VueApexCharts,
  },
  props: ["activeInstrument"],
  data: function () {
    return {
      updateTimerId: null,
      updateInterval: 5000,
      updatePeriod: "7d",
      lastCandleId: 0,
      getCandlesIntervalKey: "get_candles_interval",
      getCandlesPeriodKey: "get_candles_period",
      optionsInterval: [
        { text: "1 sec", value: "1000" },
        { text: "3 sec", value: "3000" },
        { text: "5 sec", value: "5000" },
        { text: "10 sec", value: "10000" },
      ],
      optionPeriod: [
        { text: "1 day", value: "1d" },
        { text: "7 days", value: "7d" },
        { text: "14 days", value: "14d" },
        { text: "1 month", value: "1m" },
        { text: "3 month", value: "3m" },
      ],
      chartOptions: {
        chart: chart,
        xaxis: xaxis,
        stroke: {
          width: 2,
        },
      },
      series: [
        {
          data: [],
        },
      ],
    };
  },
  methods: {
    // Обработка первого получения исторических данных
    getPeriodCandlesRows: function () {
      axios
        .get("/api/getPeriodCandles", {
          params: {
            period: this.updatePeriod,
            figi: this.activeInstrument.figi,
          },
        })
        .then((response) => this.handlePeriodCandlesResponse(response.data))
        .then(() => this.intervalFetchData())
        .catch((error) => console.log(error));
    },
    handlePeriodCandlesResponse: function (data) {
      if (data.Status.Status !== "success") {
        this.warningNotification(data.Status.Message);
      }
      if (data.Body === null) {
        return;
      }

      this.$refs.realtimeChart.updateSeries([{ data: data.Body }], false, true);
      this.lastCandleId = data.Body[0].Id;
    },
    // Обработка последующег получения данных
    getLastCandlesRows: function () {
      axios
        .get("/api/getLastCandles", {
          params: {
            lastId: this.lastCandleId,
            figi: this.activeInstrument.figi,
          },
        })
        .then((response) => this.handleLastCandlesResponse(response.data))
        .catch((error) => console.log(error));
    },
    handleLastCandlesResponse: function (data) {
      if (data.Status.Status !== "success") {
        this.warningNotification(data.Status.Message);
      }
      if (data.Body === null) {
        return;
      }
      this.series[0].data.push(...data.Body);

      this.$refs.realtimeChart.updateSeries(
        [{ data: this.series[0].data }],
        false,
        true
      );
      this.lastCandleId = data.Body[0].Id;
    },
    // Запуск итеративного получения данных
    intervalFetchData: function () {
      clearTimeout(this.updateTimerId);
      let getLastCandlesRows = this.getLastCandlesRows.bind(this);
      this.updateTimerId = setInterval(() => {
        getLastCandlesRows();
      }, this.updateInterval);
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
  watch: {
    updateInterval: function (value) {
      localStorage.setItem(this.getCandlesIntervalKey, value);
      this.intervalFetchData();
    },
    updatePeriod: function (value) {
      localStorage.setItem(this.getCandlesPeriodKey, value);
      this.getPeriodCandlesRows();
    },
    $props: {
      handler() {
        this.getPeriodCandlesRows();
      },
      deep: true,
    },
  },
  mounted() {
    this.updateInterval =
      localStorage.getItem(this.getCandlesIntervalKey) || 5000;
    this.updatePeriod = localStorage.getItem(this.getCandlesPeriodKey) || "7d";
    this.getPeriodCandlesRows();
  },
  unmounted() {
    clearTimeout(this.updateTimerId);
  },
};
</script>
