<template>
  <div class="example">
    <apexcharts
      width="800"
      height="700"
      type="line"
      :options="chartOptions"
      :series="series"
    ></apexcharts>
  </div>
</template>

<script>
import VueApexCharts from 'vue3-apexcharts'

const chart = { type: "line" };
const xaxis = { type: "datetime" };
const title = { text: "Line with Annotations", align: "left" };
const point = {
  x: new Date("25 Dec 2011").getTime(),
  y: 45,
  marker: {
    size: 5,
    fillColor: "#fff",
    strokeColor: "#333",
    strokeWidth: 3,
    shape: "circle",
    radius: 2,
  },
};

export default {
  name: "ChartTab",
  components: {
    apexcharts: VueApexCharts,
  },
  data: function () {
    return {
      chartOptions: {
        chart: chart,
        xaxis: xaxis,
        title: title,
        annotations: {
          points: [point],
        },
        stroke: {
          width: 2,
        },
      },
      series: [
        {
          data: [
            [1324508400000, 34],
            [1324594800000, 54],
            [1326236400000, 43],
          ],
        },
      ],
    };
  },
  mounted() {
    console.log("mounted tab");
  },
  unmounted() {
    console.log("unmounted tab");
  },
};
</script>
