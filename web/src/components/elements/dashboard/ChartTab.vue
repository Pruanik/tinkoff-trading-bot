<template>
  <div class="example">
    <apexcharts
      width="800"
      height="700"
      type="bar"
      :options="chartOptions"
      :series="series"
    ></apexcharts>
  </div>
</template>

<script>
import VueApexCharts from 'vue3-apexcharts'

export default {
  name: "Chart",
  components: {
    apexcharts: VueApexCharts,
  },
  data: function () {
    return {
      chartOptions: {
        chart: {
          id: "basic-bar",
        },
        xaxis: {
          categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
        }
      },
      series: [
        {
          name: "series-1",
          data: [30, 40, 45, 50, 49, 60, 70, 91],
        },
      ],
    };
  },
};
</script>
